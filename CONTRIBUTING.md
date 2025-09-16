# Contributing to GoLearn

Thanks for your interest in contributing! This project aims to be a friendly, welcoming place to learn and teach Go.

## Ways to contribute

- Report bugs and propose enhancements
- Improve exercises, hints, and documentation
- Add new exercises or solutions (see below)
- Improve the CLI UX and accessibility

## Development setup

```bash
# Clone your fork
git clone https://github.com/<you>/golearn
cd golearn

# Build the CLI
make build        # or: go build ./cmd/golearn

# Run locally
./bin/golearn help

# Initialize embedded exercises into a workspace
mkdir -p /tmp/golearn-ws && cd /tmp/golearn-ws
/path/to/bin/golearn init
```

## Running tests

This repository embeds template tests. To run tests for an exercise from the CLI repo:

```bash
# From repo root
make verify NAME=01_hello  # or: go run ./cmd/golearn verify 01_hello
```

## Code style

- Go 1.22+
- Match existing formatting; run `gofmt` and `go vet`
- Prefer clear, readable code; add concise comments for non-obvious logic
- Avoid adding heavy deps; keep the binary small and portable

## Git workflow

1. Create a feature branch
2. Make focused commits with descriptive messages
3. Ensure `make build` and `make verify` pass
4. Open a PR with a clear title and description (screenshots welcome)

## Adding or updating exercises

- Edit templates under `internal/exercises/templates/<slug>`
- Keep tests self-contained in the exercise folder
- Provide hints in `internal/exercises/catalog.yaml`
- Do not add solution code into templates; see solutions below

## Adding solutions

- Place solutions in `internal/exercises/solutions/<slug>`
- Only include implementation files; do not include tests
- Validate locally:

```bash
# Validate embedded solution against embedded tests
./bin/golearn verify <slug> --solution
```

## Docs and UX

- Update `README.md` if you add a new command or flag
- Keep output accessible; follow `internal/cli/theme` guidance

## License

- Code is MIT. Non-code lesson content and artwork are under CC BY 3.0 (see CONTENT_LICENSE).

Please also read the Code of Conduct.
