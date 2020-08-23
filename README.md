# embedr

Embedr allows you to embed static files in Go code and read them at runtime. That's it.

If you need more advanced capabilties, see [Pkger](https://github.com/markbates/pkger).

# Usage

1. `go get github.com/victorhurdugaci/embedr/cmd/embedr`
2. Add a `//go:generate embedr -include <glob>` directive in the package where you want to embed the files
3. Run `go generate .` in the package folder

See the `/example` folder for a working example.

Note: All static files must be in the same directory or a subdirectory of the package in which you embed them.

## API

- `Open(path string) (io.Reader, error)`: returns a reader that can be used to read the file specified by `path`. If the file doesn't exist returns an error.
- `Walk(walkFn WalkFunc) error`: calls `walkFn` for every embeded file. If `walkFn` returns an error, the iteration stops and `Walk` will returns the error

# I would like a new feature

Feature requests and contributions are welcome! Open an issue for discussions.
