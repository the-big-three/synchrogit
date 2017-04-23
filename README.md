# Synchrogit


Parallelized synchronization of git repositories

## Dependencies

go 1.6

## Development set-up

- Create a go workspace with 3 folder src, pkg, bin (as we follow this guideline: https://golang.org/doc/code.html#Workspaces)
- Check out this repo into src
- Set GOPATH as environment var to the go workspace root folder
- Extend PATH to include $GOPATH/bin

## Installation

When workspace is correctly configured, simply change to the root of this project and type

    go install

to install

## Usage

- Change target and source repos in synchroGitSync.json
- Execute this programm: go run main.go OR use the executable generated in installation

## Contributing

Aliaksandr Bulyha

Erik Auer

Christoph Kappel

## History

TODO

## Credits

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