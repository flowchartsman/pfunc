![logo](.github/pfunc.svg)

[![last commit](https://img.shields.io/github/last-commit/flowchartsman/pfunc?style=plastic)](https://github.com/flowchartsman/pfunc/commits/master)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/flowchartsman/pfunc?label=Latest%20Release&style=plastic)
![downloads](https://img.shields.io/github/downloads/flowchartsman/pfunc/total?style=plastic)

**PRE-RELEASE**

PFunc is a [Pulsar Function](https://pulsar.apache.org/docs/next/functions-overview/) SDK for Go.

[Pulsar](https://pulsar.apache.org/) is a distributed pub-sub messaging platform with a very flexible messaging model and an intuitive client API.

PFunc was created to add some features and to experiment with a more ergonomic API, while also serving as a proof-of-concept for the official SDK. The goal is not to compete with the official SDK, but to help improve it and potentially eventually serve as a replacement.

## Status

**PFunc is in early pre-release and is not considered production-ready**. The API is still in active development, and will likely have at several breaking changes before beta, and may include some before v1.0.0. If you want to experiment with it, version pinning is highly recommended pre 1.0.0.

Please check the [milestones](https://github.com/flowchartsman/pfunc/milestones?with_issues=yes) for more info on progress to release.

## Planned Features

- [ ] Streamlined Messaging API
  - [ ] Return one or more messages for multiple topics
  - [ ] Support for message properties
  - [ ] Support for message key (with automatic defaults)
- [ ] Improved access to context, and user config
- [ ] Support for [`slog`](https://pkg.go.dev/golang.org/x/exp/slog) as a logging interface
- [ ] Better crash reporting (initialize logging as soon as possible and report panics when possible)
- [ ] Java SDK Parity
  - [ ] [State Storage](https://pulsar.apache.org/docs/next/functions-develop-state/)
  - [ ] [Secrets](https://pulsar.apache.org/docs/next/functions-develop-security/)
  - [ ] [Admin API](https://pulsar.apache.org/docs/next/functions-develop-admin-api/)
- [ ] New Features
  - [ ] Support for go 1.19 [soft memory limit](https://github.com/golang/go/issues/48409)
  - [ ] Expanded Standard Metrics (messages output, output percentage)
  - [ ] Extended User Metrics (more than just an overloaded histogram)
  - [ ] Continuous Profiling Reporting (w/ optional reporting on a topic)

### Versioning

TBD, but will follow semver conventions first, pulsar major versioning second, if possible.

### Usage

TODO

## License

Licensed under the Apache License, Version 2.0: http://www.apache.org/licenses/LICENSE-2.0
