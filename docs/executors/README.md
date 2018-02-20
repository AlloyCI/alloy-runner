# Executors

AlloyCI Runner implements a number of executors that can be used to run your
builds in different scenarios. If you are not sure what to select, read the
[I am not sure](#i-am-not-sure) section.
Visit the [compatibility chart](#compatibility-chart) to find
out what features each executor supports and what not.

To jump into the specific documentation of each executor, visit:

- [Shell](shell.md)
- [Docker](docker.md)
- [Docker Machine (auto-scaling)](../install/autoscaling.md)
- [Parallels](parallels.md)
- [VirtualBox](virtualbox.md)
- [SSH](ssh.md)
- [Kubernetes](kubernetes.md)

## Selecting the executor

The executors support different platforms and methodologies for building a
project. The table below shows the key facts for each executor which will help
you decide.

| Executor                                          | Shell   | Docker | VirtualBox | Parallels | SSH  | Kubernetes |
|---------------------------------------------------|---------|--------|------------|-----------|------|------------|
| Clean build environment for every build           | no      | ✓      | ✓          | ✓         | no   | ✓          |
| Migrate runner machine                            | no      | ✓      | partial    | partial   | no   | ✓          |
| Zero-configuration support for concurrent builds  | no (1)  | ✓      | ✓          | ✓         | no   | ✓          |
| Complicated build environments                    | no (2)  | ✓      | ✓ (3)      | ✓ (3)     | no   | ✓          |
| Debugging build problems                          | easy    | medium | hard       | hard      | easy | medium     |

1. it's possible, but in most cases it is problematic if the build uses services
   installed on the build machine
2. it requires to install all dependencies by hand
3. for example using Vagrant

### I am not sure

**Shell** is the simplest executor to configure. All required dependencies for
your builds need to be installed manually on the machine that the Runner is
installed.

---

A better way is to use **Docker** as it allows to have a clean build environment,
with easy dependency management (all dependencies for building the project could
be put in the Docker image). The Docker executor allows you to easily create
a build environment with dependent [services], like MySQL.

---

The **Docker Machine** is a special version of the **Docker** executor
with support for auto-scaling. It works like the normal **Docker** executor
but with build hosts created on demand by _Docker Machine_.

---

The **Kubernetes**  executor allows you to use an existing Kubernetes cluster
for your builds. The executor will call the Kubernetes cluster API
and create a new Pod (with build container and services containers) for
each AlloyCI job.

---

We also offer two full system virtualization options: **VirtualBox** and
**Parallels**. This type of executor allows you to use an already created
virtual machine, which will be cloned and used to run your build. It can prove
useful if you want to run your builds on different Operating Systems since it
allows to create virtual machines with Windows, Linux, OSX or FreeBSD and make
AlloyCI Runner to connect to the virtual machine and run the build on it. Its
usage can also be useful to reduce the cost of infrastructure.

---

The **SSH** executor is added for completeness. It's the least supported
executor from all of the already mentioned ones. It makes AlloyCI Runner to
connect to some external server and run the builds there. We have some success
stories from organizations using that executor, but generally we advise to use
any of the above.

## Compatibility chart

Supported features by different executors:

| Executor                              | Shell   | Docker | VirtualBox | Parallels | SSH  | Kubernetes |
|---------------------------------------|---------|--------|------------|-----------|------|------------|
| Secure Variables                      | ✓       | ✓      | ✓          | ✓         | ✓    | ✓          |
| AlloyCI Runner Exec command            | ✓       | ✓      | no         | no        | no   | ✓          |
| alloy-ci.json: image                  | no      | ✓      | no         | no        | no   | ✓          |
| alloy-ci.json: services               | no      | ✓      | no         | no        | no   | ✓          |
| alloy-ci.json: cache                  | ✓       | ✓      | ✓          | ✓         | ✓    | ✓          |
| alloy-ci.json: artifacts              | ✓       | ✓      | ✓          | ✓         | ✓    | ✓          |
| Absolute paths: caching, artifacts    | no      | no     | no         | no        | no   | ✓          |
| Passing artifacts between stages      | ✓       | ✓      | ✓          | ✓         | ✓    | ✓          |
| Use AlloyCI Container Registry private images | n/a | ✓   | n/a        | n/a       | n/a  | ✓          |

Supported systems by different shells:

| Shells                                | Bash        | Windows Batch  | PowerShell |
|---------------------------------------|-------------|----------------|------------|
| Windows                               | ✓           | ✓ (default)    | ✓          |
| Linux                                 | ✓ (default) | no             | no         |
| OSX                                   | ✓ (default) | no             | no         |
| FreeBSD                               | ✓ (default) | no             | no         |

[services]: https://github.com/AlloyCI/alloy_ci/tree/master/doc/services/README.html
