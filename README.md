### golearn — Rustlings-like Go learning CLI

## Quickstart

- Install Go 1.22+
- Initialize exercises:
  - `go run ./cmd/golearn init`
- List:
  - `go run ./cmd/golearn list`
- Verify one:
  - `go run ./cmd/golearn verify 01_hello`
  - Edit `exercises/01_hello/hello.go` to return "Hello, Go!"
- Show progress:
  - `go run ./cmd/golearn progress`

## Docker

- Build: `docker build -t golearn .`
- Run: `docker run --rm -it -v "$PWD:/workspace" -w /workspace golearn init`

## Licensing and Attribution

- Code: MIT (see `LICENSE`).
- Non-code lesson content and any included gopher artwork: CC BY 3.0 (see `CONTENT_LICENSE`).
- Inspired by [rust-lang/rustlings](https://github.com/rust-lang/rustlings) and [mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample). If adapting content from these sources, include the required CC BY 3.0 attribution statements as noted in `NOTICE` and `CONTENT_LICENSE`.
