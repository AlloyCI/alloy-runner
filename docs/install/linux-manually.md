# Install AlloyCI Runner manually on GNU/Linux

If you don't want to use a [deb/rpm repository](linux-repository.md) to install
AlloyCI Runner, or your GNU/Linux OS is not among the supported ones, you can
install it manually.

Make sure that you read the [FAQ](../faq/README.md) section which describes
some of the most common problems with AlloyCI Runner.

## Install

CAUTION: **Important:**
With AlloyCI Runner 10, the executable was renamed to `alloy-runner`. If you
want to install a version prior to AlloyCI Runner 10, [visit the old docs](old.md).

1. Simply download one of the binaries for your system:

    ```sh
    # Linux x86-64
    sudo wget -O /usr/local/bin/alloy-runner https://alloy-runner-downloads.s3.amazonaws.com/latest/binaries/alloy-runner-linux-amd64

    # Linux x86
    sudo wget -O /usr/local/bin/alloy-runner https://alloy-runner-downloads.s3.amazonaws.com/latest/binaries/alloy-runner-linux-386

    # Linux arm
    sudo wget -O /usr/local/bin/alloy-runner https://alloy-runner-downloads.s3.amazonaws.com/latest/binaries/alloy-runner-linux-arm
    ```

    You can download a binary for every available version as described in
    [Bleeding Edge - download any other tagged release](bleeding-edge.md#download-any-other-tagged-release).

1. Give it permissions to execute:

    ```sh
    sudo chmod +x /usr/local/bin/alloy-runner
    ```

1. Optionally, if you want to use Docker, install Docker with:

    ```sh
    curl -sSL https://get.docker.com/ | sh
    ```

1. Create a AlloyCI CI user:

    ```sh
    sudo useradd --comment 'AlloyCI Runner' --create-home alloy-runner --shell /bin/bash
    ```

1. Install and run as service:

    ```sh
    sudo alloy-runner install --user=alloy-runner --working-directory=/home/alloy-runner
    sudo alloy-runner start
    ```

1. [Register the Runner](../register/index.md)

NOTE: **Note**
If `alloy-runner` is installed and run as service (what is described
in this page), it will run as root, but will execute jobs as user specified by
the `install` command. This means that some of the job functions like cache and
artifacts will need to execute `/usr/local/bin/alloy-runner` command,
therefore the user under which jobs are run, needs to have access to the executable.

## Update

1. Stop the service (you need elevated command prompt as before):

    ```sh
    sudo alloy-runner stop
    ```

1. Download the binary to replace Runner's executable:

    ```sh
    sudo wget -O /usr/local/bin/alloy-runner https://alloy-runner-downloads.s3.amazonaws.com/latest/binaries/alloy-runner-linux-386
    sudo wget -O /usr/local/bin/alloy-runner https://alloy-runner-downloads.s3.amazonaws.com/latest/binaries/alloy-runner-linux-amd64
    ```

    You can download a binary for every available version as described in
    [Bleeding Edge - download any other tagged release](bleeding-edge.md#download-any-other-tagged-release).

1. Give it permissions to execute:

    ```sh
    sudo chmod +x /usr/local/bin/alloy-runner
    ```

1. Start the service:

    ```sh
    sudo alloy-runner start
    ```
