---
last_updated: 2017-10-09
---

# Install AlloyCI Runner on FreeBSD

NOTE: **Note:**
The FreeBSD version is also available as a [bleeding edge](bleeding-edge.md)
release. Make sure that you read the [FAQ](../faq/README.md) section which
describes some of the most common problems with AlloyCI Runner.

## Installing AlloyCI Runner

Here are the steps to install and configure AlloyCI Runner under FreeBSD:

1. Create the `alloy-runner` user and group:

    ```sh
    sudo pw group add -n alloy-runner
    sudo pw user add -n alloy-runner -g alloy-runner -s /usr/local/bin/bash
    sudo mkdir /home/alloy-runner
    sudo chown alloy-runner:alloy-runner /home/alloy-runner
    ```

1. Download the binary for your system:

    ```sh
    # For amd64
    sudo fetch -o /usr/local/bin/alloy-runner https://alloy-runner-downloads.s3.amazonaws.com/latest/binaries/alloy-runner-freebsd-amd64

    # For i386
    sudo fetch -o /usr/local/bin/alloy-runner https://alloy-runner-downloads.s3.amazonaws.com/latest/binaries/alloy-runner-freebsd-386
    ```

    You can download a binary for every available version as described in
    [Bleeding Edge - download any other tagged release](bleeding-edge.md#download-any-other-tagged-release).

1. Give it permissions to execute:

    ```sh
    sudo chmod +x /usr/local/bin/alloy-runner
    ```

1. Create an empty log file with correct permissions:

    ```sh
    sudo touch /var/log/alloy_runner.log && sudo chown alloy-runner:alloy-runner /var/log/alloy_runner.log
    ```

1. Create the `rc.d` directory in case it doesn't exist:

    ```sh
    mkdir -p /usr/local/etc/rc.d
    ```

1. Create the `rc.d` script:

    ```sh
    sudo bash -c 'cat > /usr/local/etc/rc.d/alloy_runner' << "EOF"
    #!/bin/sh
    # PROVIDE: alloy_runner
    # REQUIRE: DAEMON NETWORKING
    # BEFORE:
    # KEYWORD:

    . /etc/rc.subr

    name="alloy_runner"
    rcvar="alloy_runner_enable"

    load_rc_config $name

    user="alloy-runner"
    user_home="/home/alloy-runner"
    command="/usr/local/bin/alloy-runner run"
    pidfile="/var/run/${name}.pid"

    start_cmd="alloy_runner_start"
    stop_cmd="alloy_runner_stop"
    status_cmd="alloy_runner_status"

    alloy_runner_start()
    {
        export USER=${user}
        export HOME=${user_home}
        if checkyesno ${rcvar}; then
            cd ${user_home}
            /usr/sbin/daemon -u ${user} -p ${pidfile} ${command} > /var/log/alloy_runner.log 2>&1
        fi
    }

    alloy_runner_stop()
    {
        if [ -f ${pidfile} ]; then
            kill `cat ${pidfile}`
        fi
    }

    alloy_runner_status()
    {
        if [ ! -f ${pidfile} ] || kill -0 `cat ${pidfile}`; then
            echo "Service ${name} is not running."
        else
            echo "${name} appears to be running."
        fi
    }

    run_rc_command $1
    EOF
    ```

1. Make it executable:

    ```sh
    sudo chmod +x /usr/local/etc/rc.d/alloy_runner
    ```

1. [Register the Runner](../register/README.md)
1. Enable the `alloy-runner` service and start it:

    ```sh
    sudo sysrc -f /etc/rc.conf "alloy_runner_enable=YES"
    sudo service alloy_runner start
    ```

    If you don't want to enable the `alloy-runner` service to start after a
    reboot, use:

    ```sh
    sudo service alloy_runner onestart
    ```

## Upgrading to AlloyCI Runner 1.0

To upgrade AlloyCI Runner from a version of GitLab Runner prior to 10.0:

1. Stop the Runner:

    ```sh
    sudo service gitlab_runner stop
    ```

1. Optionally, preserve the previous version of the Runner just in case:

    ```sh
    sudo mv /usr/local/bin/alloy-ci-multi-runner{,.$(/usr/local/bin/alloy-ci-multi-runner --version| grep Version | cut -d ':' -f 2 | sed 's/ //g')}
    ```

1. Download the new Runner and make it executable:

    ```sh
    # For amd64
    sudo fetch -o /usr/local/bin/alloy-runner https://alloy-runner-downloads.s3.amazonaws.com/latest/binaries/alloy-runner-freebsd-amd64

    # For i386
    sudo fetch -o /usr/local/bin/alloy-runner https://alloy-runner-downloads.s3.amazonaws.com/latest/binaries/alloy-runner-freebsd-386

    sudo chmod +x /usr/local/bin/alloy-runner
    ```

1. Edit `/usr/local/etc/rc.d/alloy_runner` and change:

    ```
    command="/usr/local/bin/gitlab-ci-multi-runner run"
    ```

    to:

    ```
    command="/usr/local/bin/alloy-runner run"
    ```

1. Start the Runner:

    ```sh
    sudo service alloy_runner start
    ```

1. After you confirm all is working correctly, you can remove the old binary:

    ```sh
    sudo rm /usr/local/bin/gitlab-ci-multi-runner.*
    ```
