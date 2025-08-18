# GoLearn

A Rustlings-like Go learning CLI.

## Install

Install the CLI globally (requires Go 1.22+):

```
go install github.com/shravan20/golearn/cmd/golearn@latest
```

This installs a `golearn` binary in your `$GOBIN` (typically `~/go/bin`). Ensure it is on your `PATH`.

## Quick start

- Initialize with embedded templates in current folder:

```
golearn init
```

- Or clone a remote exercises repository to a directory:

```
golearn init https://github.com/your-org/your-exercises my-exercises
cd my-exercises
```

- Work through exercises:

```
golearn list
golearn verify 01_hello
golearn hint 01_hello
```

- Track progress:

```
golearn progress
```

Progress is stored in `.golearn/progress.json` in the current workspace when writable, otherwise in the user config directory under `golearn/progress.json`.

## Docker

- Build: `docker build -t golearn .`
- Run: `docker run --rm -it -v "$PWD:/workspace" -w /workspace golearn init`

## Licensing and Attribution

- Code: MIT (see `LICENSE`).
- Non-code lesson content and any included gopher artwork: CC BY 3.0 (see `CONTENT_LICENSE`).
- Inspired by [rust-lang/rustlings](https://github.com/rust-lang/rustlings) and [mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample). If adapting content from these sources, include the required CC BY 3.0 attribution statements as noted in `NOTICE` and `CONTENT_LICENSE`.
