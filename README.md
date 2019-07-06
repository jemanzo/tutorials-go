# Go language - Tutorials

This repository was initialized with the command below.

```shell
$ go mod init github.com/jemanzo/tutorials-go
```

## Project Folders

Documentation:

- [`/datatypes`](./datatypes/README.md)
- [`/parsing`](./parsing/README.md)
  - `json` - Read/Write/Verify
  - `yaml` - Read/Write/Verify
  - `base64` - Read/Write/Verify
  - `strconv` - Converting Datatypes
  - `path`
    - Internet paths <URL/URI>
    - Windows & Linux folder paths
- [`/messaging`](./messaging/README.md)
  - Kafka
  - RabbitMQ
- [`/databases`](./databases/README.md)
  - Elasticsearch
  - Firebase

## Testing

```shell
$ go test .
```

## Debug

### PPROF

MacOS Install:

```shell
$ brew install graphviz
```

Download pprof file and run the following commands:

- http://localhost:8080/debug/pprof/heap
- http://localhost:8080/debug/pprof/goroutine

```shell
$ go tool pprof ~/Downloads/heap
$ go tool pprof ~/Downloads/goroutine

// Type "web" to create a .SVG file
// Type "pdf" to create a .PDF file

$ (pprof) web
$ (pprof) pdf
```

## Links

- https://github.com/hashrocket/ws
- https://github.com/hlandau/acmetool
- https://github.com/GeertJohan/go.rice
