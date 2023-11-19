# github.com/m-sign/tools
[![License][license-img]][license] [![Actions Status][action-img]][action] [![PkgGoDev][pkggodev-img]][pkggodev] [![Go Report Card][goreport-img]][goreport] [![Coverage Status][codecov-img]][codecov]

## Overview
This repository contains a command-line tool for managing the M-Sign

## Installation
Download the latest release from [GitHub Releases]() or use curl to download and install in one step:
```bash
curl -sfL -o msign https://github.com/m-sign/tools/releases/download/v0.0.1/msign_0.0.1_linux_amd64
```
## Usage
```bash
Usage:
  msign [flags]
  msign [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  key         Key operations
  self-update Self Update operation
  sign        Sign file(s)
  verify      Verify files' signatures
  version     Version info

Flags:
  -h, --help   help for msign

Use "msign [command] --help" for more information about a command.
```

## Contributing
We encourage and support an active, healthy community of contributors &mdash;
including you! Details are in the [contribution guide](CONTRIBUTING.md) and
the [code of conduct](CODE_OF_CONDUCT.md).

[![Contributor Covenant][covenant-img]](CODE_OF_CONDUCT.md)

[covenant-img]: https://img.shields.io/badge/contributor%20covenant-v2.1%20adopted-ff69b4.svg
[license-img]: https://img.shields.io/badge/license-MIT-blue.svg
[license]: LICENSE
[action-img]: ../../workflows/Test/badge.svg
[action]: ../../actions
[goreport-img]: https://goreportcard.com/badge/github.com/m-sign/tools
[goreport]: https://goreportcard.com/report/github.com/m-sign/tools
[codecov-img]: https://codecov.io/gh/github.com/m-sign/tools/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/github.com/m-sign/tools
[pkggodev-img]: https://pkg.go.dev/badge/github.com/m-sign/tools
[pkggodev]: https://pkg.go.dev/github.com/m-sign/tools

