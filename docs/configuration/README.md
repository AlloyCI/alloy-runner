# Configuring AlloyCI Runner

Below you can find some specific documentation on configuring AlloyCI Runner, the
shells supported, the security implications using the various executors, as
well as information how to set up Prometheus metrics:

- [Advanced configuration options](advanced-configuration.md) Learn how to use the [TOML][] configuration file that AlloyCI Runner uses.
- [Use self-signed certificates](tls-self-signed.md) Configure certificates that are used to verify TLS peer when connecting to the AlloyCI server.
- [Auto-scaling using Docker machine](autoscale.md) Execute jobs on machines that are created on demand using Docker machine.
- [Autoscaling AlloyCI Runner on AWS](runner_autoscale_aws/index.md)
- [Supported shells](../shells/README.md) Learn what shell script generators are supported that allow to execute builds on different systems.
- [Security considerations](../security/README.md) Be aware of potential security implications when running your jobs with AlloyCI Runner.
- [Prometheus monitoring](../monitoring/README.md) Learn how to use the Prometheus metrics HTTP server.
- [Cleanup the Docker images automatically](https://gitlab.com/gitlab-org/gitlab-runner-docker-cleanup) A simple Docker application that automatically garbage collects the AlloyCI Runner caches and images when running low on disk space.

[TOML]: https://github.com/toml-lang/toml
