# Run AlloyCI Runner in a container

This is how you can run AlloyCI Runner inside a Docker container.

## Docker image installation and configuration

1. Install Docker first:

    ```bash
    curl -sSL https://get.docker.com/ | sh
    ```

1. You need to mount a config volume into the `alloy-runner` container to
   be used for configs and other resources:

    ```bash
    docker run -d --name alloy-runner --restart always \
      -v /srv/alloy-runner/config:/etc/alloy-runner \
      -v /var/run/docker.sock:/var/run/docker.sock \
      alloy/alloy-runner:latest
    ```

    *On OSX, substitute the path "/Users/Shared" for "/srv".*

    Or, you can use a config container to mount your custom data volume:

    ```bash
    docker run -d --name alloy-runner-config \
        -v /etc/alloy-runner \
        busybox:latest \
        /bin/true

    docker run -d --name alloy-runner --restart always \
        --volumes-from alloy-runner-config \
        alloy/alloy-runner:latest
    ```

    If you plan on using Docker as the method of spawning Runners, you will need to
    mount your docker socket like this:

    ```bash
    docker run -d --name alloy-runner --restart always \
      -v /var/run/docker.sock:/var/run/docker.sock \
      -v /srv/alloy-runner/config:/etc/alloy-runner \
      alloy/alloy-runner:latest
    ```

1. [Register the Runner](../register/README.md)

Make sure that you read the [FAQ](../faq/README.md) section which describes
some of the most common problems with AlloyCI Runner.

## Update

Pull the latest version:

```bash
docker pull alloyci/alloy-runner:latest
```

Stop and remove the existing container:

```bash
docker stop alloy-runner && docker rm alloy-runner
```

Start the container as you did originally:

```bash
docker run -d --name alloy-runner --restart always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /srv/alloy-runner/config:/etc/alloy-runner \
  alloy/alloy-runner:latest
```

>**Note**:
you need to use the same method for mounting you data volume as you
did originally (`-v /srv/alloy-runner/config:/etc/alloy-runner` or
`--volumes-from alloy-runner`).

## Installing trusted SSL server certificates

If your AlloyCI server is using self-signed SSL certificates then you should
make sure the AlloyCI server certificate is trusted by the alloy-runner
container for them to be able to talk to each other.

The `alloyci/alloy-runner` image is configured to look for the trusted SSL
certificates at `/etc/alloy-runner/certs/ca.crt`, this can however be changed using the
`-e "CA_CERTIFICATES_PATH=/DIR/CERT"` configuration option.

Copy the `ca.crt` file into the `certs` directory on the data volume (or container).
The `ca.crt` file should contain the root certificates of all the servers you
want alloy-runner to trust. The alloy-runner container will
import the `ca.crt` file on startup so if your container is already running you
may need to restart it for the changes to take effect.

## Alpine Linux

You can also use alternative [Alpine Linux](https://www.alpinelinux.org/) based image with much smaller footprint:
```
alloyci/alloy-runner    latest              3e8077e209f5        13 hours ago        304.3 MB
alloyci/alloy-runner    alpine              7c431ac8f30f        13 hours ago        25.98 MB
```

**Alpine Linux image is designed to use only Docker as the method of spawning runners.**

The original `alloyci/alloy-runner:latest` is based on Ubuntu 16.04 LTS.

## SELinux

Some distributions (CentOS, RedHat, Fedora) use SELinux by default to enhance the security of the underlying system.

The special care must be taken when dealing with such configuration.

1. If you want to use Docker executor to run builds in containers you need to access the `/var/run/docker.sock`.
However, if you have a SELinux in enforcing mode, you will see the `Permission denied` when accessing the `/var/run/docker.sock`.
Install the `selinux-dockersock` and to resolve the issue: https://github.com/dpw/selinux-dockersock.

1. Make sure that persistent directory is created on host: `mkdir -p /srv/alloy-runner/config`.

1. Run docker with `:Z` on volumes:

```bash
docker run -d --name alloy-runner --restart always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /srv/alloy-runner/config:/etc/alloy-runner:Z \
  alloyci/alloy-runner:latest
```

More information about the cause and resolution can be found here:
http://www.projectatomic.io/blog/2015/06/using-volumes-with-docker-can-cause-problems-with-selinux/
