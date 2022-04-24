package main

import (
	_ "k8s.io/api/admission/v1"
	_ "k8s.io/api/admission/v1beta1"
	_ "k8s.io/api/admissionregistration/v1"
	_ "k8s.io/api/admissionregistration/v1beta1"
	_ "k8s.io/api/apiserverinternal/v1alpha1"
	_ "k8s.io/api/apps/v1"
	_ "k8s.io/api/apps/v1beta1"
	_ "k8s.io/api/apps/v1beta2"
	_ "k8s.io/api/authentication/v1"
	_ "k8s.io/api/authentication/v1beta1"
	_ "k8s.io/api/authorization/v1"
	_ "k8s.io/api/authorization/v1beta1"
	_ "k8s.io/api/autoscaling/v1"
	_ "k8s.io/api/autoscaling/v2"
	_ "k8s.io/api/autoscaling/v2beta1"
	_ "k8s.io/api/autoscaling/v2beta2"
	_ "k8s.io/api/batch/v1"
	_ "k8s.io/api/batch/v1beta1"
	_ "k8s.io/api/certificates/v1"
	_ "k8s.io/api/certificates/v1beta1"
	_ "k8s.io/api/coordination/v1"
	_ "k8s.io/api/coordination/v1beta1"
	_ "k8s.io/api/core/v1"
	_ "k8s.io/api/discovery/v1"
	_ "k8s.io/api/discovery/v1beta1"
	_ "k8s.io/api/events/v1"
	_ "k8s.io/api/events/v1beta1"
	_ "k8s.io/api/extensions/v1beta1"
	_ "k8s.io/api/flowcontrol/v1alpha1"
	_ "k8s.io/api/flowcontrol/v1beta1"
	_ "k8s.io/api/flowcontrol/v1beta2"
	_ "k8s.io/api/imagepolicy/v1alpha1"
	_ "k8s.io/api/networking/v1"
	_ "k8s.io/api/networking/v1beta1"
	_ "k8s.io/api/node/v1"
	_ "k8s.io/api/node/v1alpha1"
	_ "k8s.io/api/node/v1beta1"
	_ "k8s.io/api/policy/v1"
	_ "k8s.io/api/policy/v1beta1"
	_ "k8s.io/api/rbac/v1"
	_ "k8s.io/api/rbac/v1alpha1"
	_ "k8s.io/api/rbac/v1beta1"
	_ "k8s.io/api/scheduling/v1"
	_ "k8s.io/api/scheduling/v1alpha1"
	_ "k8s.io/api/scheduling/v1beta1"
	_ "k8s.io/api/storage/v1"
	_ "k8s.io/api/storage/v1alpha1"
	_ "k8s.io/api/storage/v1beta1"
	_ "k8s.io/apimachinery/pkg/api/resource"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	_ "k8s.io/apimachinery/pkg/apis/testapigroup/v1"
	_ "k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/apimachinery/pkg/runtime/schema"
	_ "k8s.io/apimachinery/pkg/util/intstr"
)
