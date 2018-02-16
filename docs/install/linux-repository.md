# Install AlloyCI Runner using the official AlloyCI repositories

We provide packages for the currently supported versions of Debian, Ubuntu, Mint, RHEL, Fedora, and CentOS.

| Distribution | Version | End of Life date   |
| ------------ | ------- | ------------------ |
| Debian       | buster  |                    |
| Debian       | stretch | approx. 2022       |
| Debian       | jessie  | June 2020          |
| Debian       | wheezy  | May 2018           |
| Ubuntu       | artful  |                    |
| Ubuntu       | zesty   | January 2018       |
| Ubuntu       | xenial  | April 2021         |
| Ubuntu       | trusty  | April 2019         |
| Mint         | sonya   | approx. 2021       |
| Mint         | serena  | approx. 2021       |
| Mint         | sarah   | approx. 2021       |
| Mint         | rosa    | April 2019         |
| Mint         | rafaela | April 2019         |
| Mint         | rebecca | April 2019         |
| Mint         | qiana   | April 2019         |
| RHEL/CentOS  | 7       | June 2024          |
| RHEL/CentOS  | 6       | November 2020      |
| Fedora       | 25      |                    |
| Fedora       | 26      |                    |

## Prerequisites

If you want to use the [Docker executor], make sure to install Docker before
using the Runner. [Read how to install Docker for your distribution](https://docs.docker.com/engine/installation/).

## Installing the Runner

To install the Runner:

1. Add AlloyCI's official repository:

    ```bash
    # For Debian/Ubuntu/Mint
    curl -L https://packagecloud.io/install/repositories/alloyci/alloy-runner/script.deb.sh | sudo bash

    # For RHEL/CentOS/Fedora
    curl -L https://packagecloud.io/install/repositories/alloyci/alloy-runner/script.rpm.sh | sudo bash
    ```

    >**Note:**
    _Debian users should use APT pinning_
    >
    Since Debian Stretch, Debian maintainers added their native package
    with the same name as is used by our package, and by default the official
    repositories will have a higher priority.
    >
    If you want to use our package you should manually set the source of
    the package. The best would be to add the pinning configuration file.
    Thanks to this every next update of the Runner's package - whether it will
    be done manually or automatically - will be done using the same source:
    >
    ```bash
    cat > /etc/apt/preferences.d/pin-alloy-runner.pref <<EOF
    Explanation: Prefer AlloyCI provided packages over the Debian native ones
    Package: alloy-runner
    Pin: origin packagecloud.io
    Pin-Priority: 1001
    EOF
    ```

1. Install the latest version of AlloyCI Runner, or skip to the next step to
   install a specific version:

    ```bash
    # For Debian/Ubuntu/Mint
    sudo apt-get install alloy-runner

    # For RHEL/CentOS/Fedora
    sudo yum install alloy-runner
    ```

1. To install a specific version of AlloyCI Runner:

    ```bash
    # for DEB based systems
    apt-cache madison alloy-runner
    sudo apt-get install alloy-runner=1.0.0

    # for RPM based systems
    yum list alloy-runner --showduplicates | sort -r
    sudo yum install alloy-runner-1.0.0-1
    ```

1. [Register the Runner](../register/README.md)

After completing the step above, the Runner should be started already being
ready to be used by your projects!

Make sure that you read the [FAQ](../faq/README.md) section which describes
some of the most common problems with AlloyCI Runner.

## Updating the Runner

Simply execute to install latest version:

```bash
# For Debian/Ubuntu/Mint
sudo apt-get update
sudo apt-get install alloy-runner

# For RHEL/CentOS/Fedora
sudo yum update
sudo yum install alloy-runner
```
## Manually download packages

You can manually download the packages from the following URL:
<https://packagecloud.io/alloyci/alloy-runner/>

[docker executor]: ../executors/docker.md
