# Alloy Runner Commands

Alloy Runner contains a set of commands with which you register, manage and
run your builds.

You can check a recent list of commands by executing:

```bash
alloy-runner --help
```

Append `--help` after a command to see its specific help page:

```bash
alloy-runner <command> --help
```

## Using environment variables

Most of the commands support environment variables as a method to pass the
configuration to the command.

You can see the name of the environment variable when invoking `--help` for a
specific command. For example, you can see below the help message for the `run`
command:

```bash
alloy-runner run --help
```

The output would be similar to:

```bash
NAME:
   alloy-runner run - run multi runner service

USAGE:
   alloy-runner run [command options] [arguments...]

OPTIONS:
   -c, --config "/Users/username/.alloy-runner/config.toml"	Config file [$CONFIG_FILE]
```

## Running in debug mode

Debug mode is especially useful when looking for the cause of some undefined
behavior or error.

To run a command in debug mode, prepend the command with `--debug`:

```bash
alloy-runner --debug <command>
```

## Super-user permission

Commands that access the configuration of AlloyCI Runner behave differently when
executed as super-user (`root`). The file location depends on the user executing
the command.

Be aware of the notice that is written when executing the commands that are
used for running builds, registering services or managing registered runners:

```bash
alloy-runner run

INFO[0000] Starting multi-runner from /Users/username/.alloy-runner/config.toml ...  builds=0
WARN[0000] Running in user-mode.
WARN[0000] Use sudo for system-mode:
WARN[0000] $ sudo alloy-runner...
```

You should use `user-mode` if you are really sure that this is a mode that you
want to work with. Otherwise, prefix your command with `sudo`:

```
sudo alloy-runner run

INFO[0000] Starting multi-runner from /etc/alloy-runner/config.toml ...  builds=0
INFO[0000] Running in system-mode.
```

In the case of **Windows** you may need to run the **Command Prompt** in
**Administrative Mode**.

## Configuration file

Alloy Runner configuration uses the [TOML] format.

The file to be edited can be found in:

1. `/etc/alloy-runner/config.toml` on \*nix systems when alloy-runner is
   executed as super-user (`root`)
1. `~/.alloy-runner/config.toml` on \*nix systems when alloy-runner is
   executed as non-root
1. `./config.toml` on other systems

Most of the commands accept an argument to specify a custom configuration file,
allowing you to have a multiple different configurations on a single machine.
To specify a custom configuration file use the `-c` or `--config` flag, or use
the `CONFIG_FILE` environment variable.

[TOML]: https://github.com/toml-lang/toml

## Signals

It is possible to use system signals to interact with AlloyCI Runner. The
following commands support the following signals:

| Command | Signal | Action |
|---------|--------|--------|
| `register` | **SIGINT** | Cancel runner registration and delete if it was already registered |
| `run`, `exec`, `run-single` | **SIGINT**, **SIGTERM** | Abort all running builds and exit as soon as possible. Use twice to exit now (**forceful shutdown**). |
| `run`, `exec`, `run-single` | **SIGQUIT** | Stop accepting a new builds. Exit as soon as currently running builds do finish (**graceful shutdown**). |
| `run` | **SIGHUP** | Force to reload configuration file |

If your operating system is configured to automatically restart the service if it fails (which is the default on some platforms) it may automatically restart the runner if it's shut down by the signals above.

## Commands overview

This is what you see if you run `alloy-runner` without any arguments:

```bash
NAME:
   alloy-runner - a AlloyCI Runner

USAGE:
   alloy-runner [global options] command [command options] [arguments...]

VERSION:
   1.0.0~beta.142.ga8d37f3 (a8d37f3)

AUTHOR(S):
   Patricio Cano <support@alloy-ci.com>

COMMANDS:
   exec		execute a build locally
   run		run multi runner service
   register	register a new runner
   install	install service
   uninstall	uninstall service
   start	start service
   stop		stop service
   restart	restart service
   status	get status of a service
   run-single	start single runner
   unregister	unregister specific runner
   verify	verify all registered runners
   archive	find and archive files (internal)
   artifacts	upload build artifacts (internal)
   extract	extract files from an archive (internal)
   help, h	Shows a list of commands or help for one command
```

Below we will explain what each command does in detail.

## Registration-related commands

The following commands allow you to register a new runner, or list and verify
them if they are still registered.

- [alloy-runner register](#alloy-runner-register)
    - [Interactive registration](#interactive-registration)
    - [Non-interactive registration](#non-interactive-registration)
- [alloy-runner list](#alloy-runner-list)
- [alloy-runner verify](#alloy-runner-verify)
- [alloy-runner unregister](#alloy-runner-unregister)

The above commands support the following arguments:

| Parameter   | Default | Description |
|-------------|---------|-------------|
| `--config`  | See the [configuration file section](#configuration-file) | Specify a custom configuration file to be used |

### alloy-runner register

This command registers your Alloy Runner in AlloyCI. The registered runner is
added to the [configuration file](#configuration-file).
You can use multiple configurations in a single AlloyCI Runner. Executing
`alloy-runner register` adds a new configuration entry, it doesn't remove the
previous ones.

There are two options to register a Runner, interactive and non-interactive.

#### Interactive registration

This command is usually used in interactive mode (**default**). You will be
asked multiple questions during a Runner's registration.

This question can be pre-filled by adding arguments when invoking the registration command:

    alloy-runner register --name my-runner --url http://alloyci.example.com --registration-token my-registration-token

Or by configuring the environment variable before the `register` command:

    export CI_SERVER_URL=http://alloyci.example.com
    export RUNNER_NAME=my-runner
    export REGISTRATION_TOKEN=my-registration-token
    export REGISTER_NON_INTERACTIVE=true
    alloy-runner register

To check all possible arguments and environments execute:

    alloy-runner register --help

#### Non-interactive registration

It's possible to use registration in non-interactive / unattended mode.

You can specify the arguments when invoking the registration command:

    alloy-runner register --non-interactive <other-arguments>

Or by configuring the environment variable before the `register` command:

    <other-environment-variables>
    export REGISTER_NON_INTERACTIVE=true
    alloy-runner register

> **Note:** Boolean parameters must be passed in the command line with `--key={true|false}`.

### alloy-runner list

This command lists all runners saved in the
[configuration file](#configuration-file).

### alloy-runner verify

This command checks if the registered runners can connect to AlloyCI, but it
doesn't verify if the runners are being used by the AlloyCI Runner service. An
example output is:

```bash
Verifying runner... is alive                        runner=fee9938e
Verifying runner... is alive                        runner=0db52b31
Verifying runner... is alive                        runner=826f687f
Verifying runner... is alive                        runner=32773c0f
```

To delete the old and removed from AlloyCI runners, execute the following
command.

>**Warning:**
This operation cannot be undone, it will update the configuration file, so
make sure to have a backup of `config.toml` before executing it.

```bash
alloy-runner verify --delete
```

### alloy-runner unregister

This command allows to unregister one of the registered runners. It expects
either a full URL and the runner's token, or the runner's name. With the
`--all-runners` option it will unregister all the attached runners.

To unregister a specific runner, first get the runner's details by executing
`alloy-runner list`:

```bash
test-runner     Executor=shell Token=t0k3n URL=http://alloyci.example.com
```

Then use this information to unregister it, using one of the following commands.

>**Warning:**
This operation cannot be undone, it will update the configuration file, so
make sure to have a backup of `config.toml` before executing it.

#### By URL and token:

```bash
alloy-runner unregister --url http://alloyci.example.com/ --token t0k3n
```

#### By name:

> **Note:** If there is more than one runner with the given name, only the first one will be removed

```bash
alloy-runner unregister --name test-runner
```

#### All Runners:

```bash
alloy-runner unregister --all-runners
```

## Service-related commands

> **Note:** Starting with AlloyCI Runner 1.0.0, service-related commands are **deprecated**
and will be removed in one of the upcoming releases.

The following commands allow you to manage the runner as a system or user
service. Use them to install, uninstall, start and stop the runner service.

- [alloy-runner install](#alloy-runner-install)
- [alloy-runner uninstall](#alloy-runner-uninstall)
- [alloy-runner start](#alloy-runner-start)
- [alloy-runner stop](#alloy-runner-stop)
- [alloy-runner restart](#alloy-runner-restart)
- [alloy-runner status](#alloy-runner-status)
- [Multiple services](#multiple-services)
- [**Access Denied** when running the service-related commands](#access-denied-when-running-the-service-related-commands)

All service related commands accept these arguments:

| Parameter | Default | Description |
|-----------|---------|-------------|
| `--service` | `alloy-runner` | Specify custom service name |
| `--config` | See the [configuration file](#configuration-file) | Specify a custom configuration file to use |

### alloy-runner install

This command installs AlloyCI Runner as a service. It accepts different sets of
arguments depending on which system it's run on.

When run on **Windows** or as super-user, it accepts the `--user` flag which
allows you to drop privileges of builds run with the **shell** executor.

| Parameter | Default | Description |
|-----------|---------|-------------|
| `--service`      | `alloy-runner` | Specify a custom name for the Runner |
| `--working-directory` | the current directory | Specify the root directory where all data will be stored when builds will be run with the **shell** executor |
| `--user`              | `root` | Specify the user which will be used to execute builds |
| `--password`          | none   | Specify the password for the user that will be used to execute the builds |

### alloy-runner uninstall

This command stops and uninstalls the AlloyCI Runner from being run as an
service.

### alloy-runner start

This command starts the AlloyCI Runner service.

### alloy-runner stop

This command stops the AlloyCI Runner service.

### alloy-runner restart

This command stops and then starts the AlloyCI Runner service.

### alloy-runner status

This command prints the status of the AlloyCI Runner service. The exit code is zero when the service is running and non-zero when the service is not running.

### Multiple services

By specifying the `--service` flag, it is possible to have multiple AlloyCI
Runner services installed, with multiple separate configurations.

## Run-related commands

This command allows to fetch and process builds from AlloyCI.

### alloy-runner run

This is main command that is executed when AlloyCI Runner is started as a
service. It reads all defined Runners from `config.toml` and tries to run all
of them.

The command is executed and works until it [receives a signal](#signals).

It accepts the following parameters.

| Parameter | Default | Description |
|-----------|---------|-------------|
| `--config`  | See [#configuration-file](#configuration-file) | Specify a custom configuration file to be used |
| `--working-directory` | the current directory | Specify the root directory where all data will be stored when builds will be run with the **shell** executor |
| `--user`    | the current user | Specify the user that will be used to execute builds |
| `--syslog`  | `false` | Send all logs to SysLog (Unix) or EventLog (Windows) |
| `--metrics-server` | empty | Address (`<host>:<port>`) on which the Prometheus metrics HTTP server should be listening |

### alloy-runner run-single

This is a supplementary command that can be used to run only a single build
from a single AlloyCI instance. It doesn't use any configuration file and
requires to pass all options either as parameters or environment variables.
The AlloyCI URL and Runner token need to be specified too.

For example:

```bash
alloy-runner run-single -u http://alloyci.example.com -t my-runner-token --executor docker --docker-image ruby:2.1
```

You can see all possible configuration options by using the `--help` flag:

```bash
alloy-runner run-single --help
```

You can use the `--max-builds` option to control how many builds the runner will execute before exiting.  The
default of `0` means that the runner has no build limit and will run jobs forever.

You can also use the `--wait-timeout` option to control how long the runner will wait for a job before
exiting.  The default of `0` means that the runner has no timeout and will wait forever between jobs.

## Internal commands

AlloyCI Runner is distributed as a single binary and contains a few internal
commands that are used during builds.

### alloy-runner artifacts-downloader

Download the artifacts archive from AlloyCI.

### alloy-runner artifacts-uploader

Upload the artifacts archive to AlloyCI.

### alloy-runner cache-archiver

Create a cache archive, store it locally or upload it to an external server.

### alloy-runner cache-extractor

Restore the cache archive from a locally or externally stored file.

## Troubleshooting

Below are some common pitfalls.

### **Access Denied** when running the service-related commands

Usually the [service related commands](#service-related-commands) require
administrator privileges:

- On Unix (Linux, OSX, FreeBSD) systems, prefix `alloy-runner` with `sudo`
- On Windows systems use the elevated command prompt.
  Run an `Administrator` command prompt ([How to][prompt]).
  The simplest way is to write `Command Prompt` in the Windows search field,
  right click and select `Run as administrator`. You will be asked to confirm
  that you want to execute the elevated command prompt.
