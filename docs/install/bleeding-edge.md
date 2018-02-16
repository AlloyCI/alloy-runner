# AlloyCI Runner bleeding edge releases

CAUTION: **Warning:**
These are the latest, probably untested releases of AlloyCI Runner built straight
from `master` branch. Use at your own risk.

## Download the standalone binaries

* https://alloy-runner-downloads.s3.amazonaws.com/master/binaries/alloy-runner-linux-386
* https://alloy-runner-downloads.s3.amazonaws.com/master/binaries/alloy-runner-linux-amd64
* https://alloy-runner-downloads.s3.amazonaws.com/master/binaries/alloy-runner-linux-arm
* https://alloy-runner-downloads.s3.amazonaws.com/master/binaries/alloy-runner-darwin-386
* https://alloy-runner-downloads.s3.amazonaws.com/master/binaries/alloy-runner-darwin-amd64
* https://alloy-runner-downloads.s3.amazonaws.com/master/binaries/alloy-runner-windows-386.exe
* https://alloy-runner-downloads.s3.amazonaws.com/master/binaries/alloy-runner-windows-amd64.exe
* https://alloy-runner-downloads.s3.amazonaws.com/master/binaries/alloy-runner-freebsd-386
* https://alloy-runner-downloads.s3.amazonaws.com/master/binaries/alloy-runner-freebsd-amd64
* https://alloy-runner-downloads.s3.amazonaws.com/master/binaries/alloy-runner-freebsd-arm

You can then run the Runner with:
```bash
chmod +x alloy-runner-linux-amd64
./alloy-runner-linux-amd64 run
```

## Download one of the packages for Debian or Ubuntu

* https://alloy-runner-downloads.s3.amazonaws.com/master/deb/alloy-runner_i386.deb
* https://alloy-runner-downloads.s3.amazonaws.com/master/deb/alloy-runner_amd64.deb
* https://alloy-runner-downloads.s3.amazonaws.com/master/deb/alloy-runner_arm.deb
* https://alloy-runner-downloads.s3.amazonaws.com/master/deb/alloy-runner_armhf.deb

You can then install it with:
```bash
dpkg -i alloy-runner_386.deb
```

## Download one of the packages for RedHat or CentOS

* https://alloy-runner-downloads.s3.amazonaws.com/master/rpm/alloy-runner_i686.rpm
* https://alloy-runner-downloads.s3.amazonaws.com/master/rpm/alloy-runner_amd64.rpm
* https://alloy-runner-downloads.s3.amazonaws.com/master/rpm/alloy-runner_arm.rpm
* https://alloy-runner-downloads.s3.amazonaws.com/master/rpm/alloy-runner_armhf.rpm

You can then install it with:
```bash
rpm -i alloy-runner_386.rpm
```

## Download any other tagged release

Simply replace `master` with either `tag` (v0.2.0 or 0.4.2) or `latest` (the latest
stable). For a list of tags see <https://gitlab.com/AlloyCI/alloy-runner/tags>.
For example:

* https://alloy-runner-downloads.s3.amazonaws.com/master/binaries/alloy-runner-linux-386
* https://alloy-runner-downloads.s3.amazonaws.com/latest/binaries/alloy-runner-linux-386
* https://alloy-runner-downloads.s3.amazonaws.com/v0.2.0/binaries/alloy-runner-linux-386

If you have problem downloading through `https`, fallback to plain `http`:

* http://alloy-runner-downloads.s3.amazonaws.com/master/binaries/alloy-runner-linux-386
* http://alloy-runner-downloads.s3.amazonaws.com/latest/binaries/alloy-runner-linux-386
* http://alloy-runner-downloads.s3.amazonaws.com/v0.2.0/binaries/alloy-runner-linux-386
