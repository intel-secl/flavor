# Flavor Library

The Flavor library is responsible for retrieving the hardware/software measurements of the host in a pre-defined format.
When integrated with the Host Verification service, Flavor would be equivalent to what is known today as the Whitelist.

cmd - contains the command line interface utility to use the library

## Key features
- Create a flavors for VM and container


## System Requirements
- RHEL 8.1
- Epel 8 Repo
- Proxy settings if applicable

## Software requirements
- git
- `go` version `1.14.4`

# Step By Step Build Instructions

## Install required shell commands

### Install `go` version `go1.14.4`
The `flavor-library` requires Go version 1.14.4 that has support for `go modules`. The build was validated with the latest version go1.14.4 of `go`. It is recommended that you use go1.14.4 version of `go`. You can use the following to install `go`.

```shell
wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
tar -xzf go1.14.4.linux-amd64.tar.gz
sudo mv go /usr/local
export GOROOT=/usr/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

## Build Flavor

- Git clone the flavor
- Run scripts to build the flavor

```shell
git clone https://github.com/intel-secl/flavor.git
cd flavor
GOOS=linux GOSUMDB=off GOPROXY=direct go mod tidy
go build ./...
```

## Third Party Dependencies

## Flavor library

### Direct dependencies

| Name                  | Repo URL                        | Minimum Version Required              |
| ----------------------| --------------------------------| :------------------------------------:|
| logrus                | github.com/sirupsen/logrus      | v1.4.0                                |
| testify               | github.com/stretchr/testify     | v1.3.0                                |
| uuid                  | github.com/satori/go.uuid       | v1.2.1-0.20181028125025-b2ce2384e17b  |

*Note: All dependencies are listed in go.mod*

# Links
https://01.org/intel-secl/
