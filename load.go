// Copyright 2023 Edward McFarlane. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"

	"go.starlark.net/starlark"
)

// call is an in-flight or completed load call
type call struct {
	wg  sync.WaitGroup
	val starlark.StringDict
	err error

	caller *starlark.Thread // cycle detection
}

// Loader is a starlark.Thread-compatible module loader.
type Loader struct {
	fs      fs.FS
	globals starlark.StringDict // Predeclared globals

	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}

func NewLoader(fs fs.FS, globals starlark.StringDict) *Loader {
	return &Loader{
		fs:      fs,
		globals: globals,
	}
}

// errCycle indicates the load caused a cycle.
var errCycle = errors.New("cycle in loading module")

// A panicError is an arbitrary value recovered from a panic
// with the stack trace during the execution of given function.
type panicError struct {
	value interface{}
	stack []byte
}

// Error implements error interface.
func (p *panicError) Error() string {
	return fmt.Sprintf("%v\n\n%s", p.value, p.stack)
}
func newPanicError(v interface{}) error {
	stack := debug.Stack()

	// The first line of the stack trace is of the form "goroutine N [status]:"
	// but by the time the panic reaches Do the goroutine may no longer exist
	// and its status will have changed. Trim out the misleading line.
	if line := bytes.IndexByte(stack[:], '\n'); line >= 0 {
		stack = stack[line+1:]
	}
	return &panicError{value: v, stack: stack}
}

// Load implements starlark.Thread-compatible module loader.
func (l *Loader) Load(thread *starlark.Thread, module string) (starlark.StringDict, error) {
	// handle relative paths using the current thread's filename
	if strings.HasPrefix(module, ".") {
		module = filepath.Join(filepath.Dir(thread.Name), module)
	}

	l.mu.Lock()
	if c, ok := l.m[module]; ok {
		if c.caller == thread {
			l.mu.Unlock()
			return nil, fmt.Errorf("%s: %w", module, errCycle)
		}
		l.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)
	c.caller = thread
	if l.m == nil {
		l.m = make(map[string]*call)
	}
	l.m[module] = c
	l.mu.Unlock()

	c.val, c.err = l.load(thread, module)
	c.wg.Done()

	l.mu.Lock()
	c.caller = nil
	l.mu.Unlock()

	return c.val, c.err
}

func (l *Loader) load(thread *starlark.Thread, module string) (v starlark.StringDict, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = newPanicError(r)
		}
	}()

	src, err := fs.ReadFile(l.fs, module)
	if err != nil {
		return nil, err
	}

	oldName, newName := thread.Name, module
	thread.Name = newName
	defer func() { thread.Name = oldName }()

	return starlark.ExecFile(thread, module, src, l.globals)
}
