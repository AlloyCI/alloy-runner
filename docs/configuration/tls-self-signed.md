# The self-signed certificates or custom Certification Authorities

AlloyCI Runner allows you to configure certificates that
are used to verify TLS peer when connecting to the AlloyCI server.

**This allows to solve the `x509: certificate signed by unknown authority` problem when registering runner.**

## Supported options for self-signed certificates

AlloyCI Runner provides these options:

1. **Default**: AlloyCI Runner reads system certificate store and verifies the AlloyCI server against the CA's stored in system.

2. AlloyCI Runner reads the PEM (**DER format is not supported**) certificate from predefined file:

        - `/etc/alloy-runner/certs/hostname.crt` on *nix systems when alloy-runner is executed as root.
        - `~/.alloy-runner/certs/hostname.crt` on *nix systems when alloy-runner is executed as non-root,
        - `./certs/hostname.crt` on other systems.

        If the address of your server is: `https://my.alloy.server.com:8443/`.
        Create the certificate file at: `/etc/alloy-runner/certs/my.alloy.server.com.crt`.

    > **Note:** You may need to concatenate the intermediate and server certificate
      for the chain to be properly identified.
3. AlloyCI Runner exposes `tls-ca-file` option during registration and in [`config.toml`](advanced-configuration.md)
which allows you to specify custom file with certificates. This file will be read every time when runner tries to
access the AlloyCI server.

## Git cloning

The runner injects missing certificates to build CA chain to build containers.
This allows the `git clone` and `artifacts` to work with servers that do not use publicly trusted certificates.

This approach is secure, but makes the runner a single point of trust.
