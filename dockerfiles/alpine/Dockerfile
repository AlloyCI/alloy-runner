FROM alpine

RUN adduser -D -S -h /home/alloy-runner alloy-runner

RUN apk add --update \
    bash \
    ca-certificates \
    git \
    openssl \
    tzdata \
    wget

ARG DOCKER_MACHINE_VERSION=0.13.0
ARG DUMB_INIT_VERSION=1.0.2

COPY alloy-runner-linux-amd64 /usr/bin/alloy-runner
COPY checksums /tmp/
RUN chmod +x /usr/bin/alloy-runner && \
    ln -s /usr/bin/alloy-runner /usr/bin/alloy-ci-multi-runner && \
    alloy-runner --version && \
    mkdir -p /etc/alloy-runner/certs && \
    chmod -R 700 /etc/alloy-runner && \
    wget -q https://github.com/docker/machine/releases/download/v${DOCKER_MACHINE_VERSION}/docker-machine-Linux-x86_64 -O /usr/bin/docker-machine && \
    chmod +x /usr/bin/docker-machine && \
    docker-machine --version && \
    wget -q https://github.com/Yelp/dumb-init/releases/download/v${DUMB_INIT_VERSION}/dumb-init_${DUMB_INIT_VERSION}_amd64 -O /usr/bin/dumb-init && \
    chmod +x /usr/bin/dumb-init && \
    dumb-init --version && \
    sha256sum -c -w /tmp/checksums

COPY entrypoint /
RUN chmod +x /entrypoint

VOLUME ["/etc/alloy-runner", "/home/alloy-runner"]
ENTRYPOINT ["/usr/bin/dumb-init", "/entrypoint"]
CMD ["run", "--user=alloy-runner", "--working-directory=/home/alloy-runner"]
