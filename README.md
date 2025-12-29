# DapperGo

A lightweight distributed tracing library for Go, inspired by Google's Dapper.

## Features

- Low overhead and low latency tracing.
- Support for spans, tracers, and recorders.
- gRPC based span collection.
- Context-based trace propagation.
- Probabilistic sampling.

## Installation

```bash
go get github.com/sagnik3788/dappergo
```

## Usage

Basic usage example:

```go
recorder := tracer.NewRecorder(1024, grpcClient)
t := tracer.NewTracer(recorder)

span := t.StartSpan("my-operation")
defer span.Finish()

// ... do work ...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
