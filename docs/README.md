# AlloyCI Runner

[![Build Status](https://gitlab.com/AlloyCI/alloy-runner/badges/master/build.svg)](https://github.com/AlloyCI/alloy-runner)

AlloyCI Runner is the open source project that is used to run your jobs and
send the results back to AlloyCI.
## Requirements

AlloyCI Runner is written in [Go][golang] and can be run as a single binary, no
language specific requirements are needed.

It is designed to run on the GNU/Linux, macOS, and Windows operating systems.
Other operating systems will probably work as long as you can compile a Go
binary on them.

## Features

- Allows to run:
 - multiple jobs concurrently
 - use multiple tokens with multiple server (even per-project)
 - limit number of concurrent jobs per-token
- Jobs can be run:
 - locally
 - using Docker containers
 - using Docker containers and executing job over SSH
 - using Docker containers with autoscaling on different clouds and virtualization hypervisors
 - connecting to remote SSH server
- Is written in Go and distributed as single binary without any other requirements
- Supports Bash, Windows Batch and Windows PowerShell
- Works on GNU/Linux, OS X and Windows (pretty much anywhere you can run Docker)
- Allows to customize the job running environment
- Automatic configuration reload without restart
- Easy to use setup with support for Docker, Docker-SSH, Parallels or SSH running environments
- Enables caching of Docker containers
- Easy installation as a service for GNU/Linux, OSX and Windows
- Embedded Prometheus metrics HTTP server

## [Install AlloyCI Runner](install/README.md)

AlloyCI Runner can be installed and used on GNU/Linux, macOS, FreeBSD and Windows.
You can install it using Docker, download the binary manually or use the
repository for rpm/deb packages that AlloyCI offers. Below you can find
information on the different installation methods:

- [Install using AlloyCI's repository for Debian/Ubuntu/CentOS/RedHat (preferred)](install/linux-repository.md)
- [Install on GNU/Linux manually (advanced)](install/linux-manually.md)
- [Install on macOS (preferred)](install/osx.md)
- [Install on Windows (preferred)](install/windows.md)
- [Install as a Docker Service](install/docker.md)
- [Install in Auto-scaling mode using Docker machine](install/autoscaling.md)
- [Install on FreeBSD](install/freebsd.md)
- [Install on Kubernetes](install/kubernetes.md)
- [Install the nightly binary manually (development)](install/bleeding-edge.md)

## [Register AlloyCI Runner](register/README.md)

Once AlloyCI Runner is installed, you need to register it with AlloyCI.

Learn how to [register a AlloyCI Runner](register/README.md).

## Using AlloyCI Runner

- [See the commands documentation](commands/README.md)

## [Selecting the executor](executors/README.md)

AlloyCI Runner implements a number of executors that can be used to run your
builds in different scenarios. If you are not sure what to select, read the
[I am not sure](executors/README.md#i-am-not-sure) section.
Visit the [compatibility chart](executors/README.md#compatibility-chart) to find
out what features each executor supports and what not.

To jump into the specific documentation of each executor, visit:

- [Shell](executors/shell.md)
- [Docker](executors/docker.md)
- [Docker Machine and Docker Machine SSH (auto-scaling)](install/autoscaling.md)
- [Parallels](executors/parallels.md)
- [VirtualBox](executors/virtualbox.md)
- [SSH](executors/ssh.md)
- [Kubernetes](executors/kubernetes.md)

## [Advanced Configuration](configuration/index.md)

- [Advanced configuration options](configuration/advanced-configuration.md) Learn how to use the [TOML][] configuration file that AlloyCI Runner uses.
- [Use self-signed certificates](configuration/tls-self-signed.md) Configure certificates that are used to verify TLS peer when connecting to the AlloyCI server.
- [Auto-scaling using Docker machine](configuration/autoscale.md) Execute jobs on machines that are created on demand using Docker machine.
- [Autoscaling AlloyCI Runner on AWS](configuration/runner_autoscale_aws/index.md)
- [Supported shells](shells/README.md) Learn what shell script generators are supported that allow to execute builds on different systems.
- [Security considerations](security/index.md) Be aware of potential security implications when running your jobs with AlloyCI Runner.
- [Runner monitoring](monitoring/README.md) Learn how to monitor Runner's behavior.
- [Cleanup the Docker images automatically](https://alloy.com/alloy-org/alloy-runner-docker-cleanup) A simple Docker application that automatically garbage collects the AlloyCI Runner caches and images when running low on disk space.

## Troubleshooting

Read the [FAQ](faq/README.md) for troubleshooting common issues.

## Changelog

Visit [Changelog] to view recent changes.

## License

This code is distributed under the MIT license, see the [LICENSE][] file.

[Changelog]: https://github.com/AlloyCI/alloy-runner/blob/master/CHANGELOG.md
[golang]: https://golang.org/
[LICENSE]: https://github.com/AlloyCI/alloy-runner/tree/master/LICENSE
[TOML]: https://github.com/toml-lang/toml
