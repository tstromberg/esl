# esl

Consume EndpointSecurity events on macOS via eslogger(1)

## Usage

This is mostly an experimental library for consuming EndpointSecurity events from Go, but we also provide an example command-line.

Will show you information on file open calls:

```shell
go install github.com/tstromberg/esl/cmd/esl@latest
esl open
```

## Note

The `eslogger(1)` commannd provided by Apple is not designed to be a stable API. YMMV.

