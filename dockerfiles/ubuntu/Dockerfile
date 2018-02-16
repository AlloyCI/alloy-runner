FROM ubuntu:16.04

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y ca-certificates wget apt-transport-https vim nano git curl && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ARG DOCKER_MACHINE_VERSION=0.13.0
ARG DUMB_INIT_VERSION=1.0.2

COPY alloy-runner_amd64.deb /tmp/
COPY checksums /tmp/
RUN dpkg -i /tmp/alloy-runner_amd64.deb; \
    apt-get update &&  \
    apt-get -f install -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    rm /tmp/alloy-runner_amd64.deb && \
    alloy-runner --version && \
    mkdir -p /etc/alloy-runner/certs && \
    chmod -R 700 /etc/alloy-runner && \
    wget -q https://github.com/docker/machine/releases/download/v${DOCKER_MACHINE_VERSION}/docker-machine-Linux-x86_64 -O /usr/bin/docker-machine && \
    chmod +x /usr/bin/docker-machine && \
    docker-machine --version && \
    wget -q https://github.com/Yelp/dumb-init/releases/download/v${DUMB_INIT_VERSION}/dumb-init_${DUMB_INIT_VERSION}_amd64 -O /usr/bin/dumb-init && \
    chmod +x /usr/bin/dumb-init && \
    dumb-init --version && \
    sha256sum --check --strict /tmp/checksums

COPY entrypoint /
RUN chmod +x /entrypoint

VOLUME ["/etc/alloy-runner", "/home/alloy-runner"]
ENTRYPOINT ["/usr/bin/dumb-init", "/entrypoint"]
CMD ["run", "--user=alloy-runner", "--working-directory=/home/alloy-runner"]
