# Build the deployment spec documented here:
# https://kubernetes.io/docs/concepts/workloads/controllers/deployment/


def nginx_container(port=80):
    return api_core_v1.Container(
        name="nginx",
        image="nginx:1.14.2",
        ports=[
            api_core_v1.ContainerPort(containerPort=port),
        ],
    )


def deployment():
    # Returns the top level protobuf object.
    return api_apps_v1.Deployment(
        metadata={
            "name": "nginx-deployment",
            "labels": {
                "app": "nginx",
            },
        },
        spec=api_apps_v1.DeploymentSpec(
            replicas=3,
            selector=struct(matchLabels={
                "app": "nginx",
            }),
            template=api_core_v1.PodTemplateSpec(
                metadata=apimachinery_pkg_apis_meta_v1.ObjectMeta(
                    labels={
                        "app": "nginx",
                    }),
                spec=api_core_v1.PodSpec(containers=[
                    nginx_container(),
                ]),
            ),
        ),
    )


def main():
    return [
        deployment(),
    ]
