# pakku

[![Go Reference](https://pkg.go.dev/badge/github.com/mycreepy/pakku.svg)](https://pkg.go.dev/github.com/mycreepy/pakku)
[![Go Report Card](https://goreportcard.com/badge/github.com/mycreepy/pakku?style=flat-square)](https://goreportcard.com/report/github.com/mycreepy/pakku)
[![Go Build & Test](https://github.com/mycrEEpy/pakku/actions/workflows/build.yml/badge.svg)](https://github.com/mycrEEpy/pakku/actions/workflows/build.yml)

`pakku` is a declarative approach to your system package managers.

## Installation

You can download the pre-built binaries from the GitHub Releases or use various other installation methods.

Using Go:

```shell
go install github.com/mycreepy/pakku@latest
```

Using Pkgx:

```shell
pkgx install github.com/mycreepy/pakku
```

## Usage

* Initialize a new config:
  * `pakku init`
* Add some packages to your config:
  * `pakku add apt curl`
  * `pakku add brew awscli`
* Install the packages on your system:
  * `pakku apply`
