# synchrogit

Parallelized synchronization of git repositories, written in go

This is a fun project to learn [golang](https://golang.org/)


## How does this work?

TODO

## Dependencies

go >= 1.6

## Development set-up

1. Follow the [official documentation](https://golang.org/doc/install) on setting up your `go workspace
2. Fork the repository 
3. Install `synchrogit` with `go get github.com/the-big-three/synchrogit`
4. Change into dir `cd $GOPATH/src/github.com/the-big-three/synchrogit`
5. Add your fork repo as origin `git remote add fork git@github.com:<your_username>/synchrogit.git`

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