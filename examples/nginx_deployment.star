# Build the deployment spec documented here:
# https://kubernetes.io/docs/concepts/workloads/controllers/deployment/


def nginx_container(port=80):
    return core_v1.Container(
        name="nginx",
        image="nginx:1.14.2",
        ports=[
            core_v1.ContainerPort(containerPort=port),
        ],
    )


def deployment():
    # Returns the top level protobuf object.
    return apps_v1.Deployment(
        metadata={
            "name": "nginx-deployment",
            "labels": {
                "app": "nginx",
            },
        },
        spec=apps_v1.DeploymentSpec(
            replicas=3,
            selector=struct(matchLabels={
                "app": "nginx",
            }),
            template=core_v1.PodTemplateSpec(
                metadata=apis_meta_v1.ObjectMeta(labels={
                    "app": "nginx",
                }),
                spec=core_v1.PodSpec(containers=[
                    nginx_container(),
                ]),
            ),
        ),
    )


def main():
    return [
        deployment(),
    ]
