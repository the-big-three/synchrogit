# synchrogit

Parallelized synchronization of git repositories, written in go

This is a fun project to learn [golang](https://golang.org/)


## How does this work?

TODO

## Dependencies

go >= 1.6

## Development set-up

Follow the [official documentation](https://golang.org/doc/install) on setting up your `go workspace

## Installation

TODO

## Usage

- Change target and source repos in synchroGitSync.json
- Execute this programm: go run main.go OR use the executable generated in installation

## Contributing`

TODO

## Contributors (go newbies) and Kudos

Aliaksandr Bulyha

Erik Auer

[supergicko](https://github.com/supergicko)

## Changelog

TODO

## License

TODO

## Other

### File naming conventions
There's a few guidelines to follow.
- File names that begin with "." or "_" are ignored by the go tool
- Files with the suffix _test.go are only compiled and run by the go test tool.
- Files with os and architecture specific suffixes automatically follow those same constraints, e.g. name_linux.go will only build on linux, name_amd64.go will only build on amd64. This is the same as having a //+build amd64 line at the top of the file

See the docs for the go build tool for more details: https://golang.org/pkg/go/build/