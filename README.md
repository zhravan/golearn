# GoLearn

Rustlings‑style Go exercises in a tiny CLI.

## Why this exists

This project is my attempt to learn Go by building as I learn, making the journey more engaging while exploring the language in practice. I'm sharing the exercises and tooling so others can learn alongside me.

## Install (Go 1.22+)

```bash
go install github.com/shravan20/golearn/cmd/golearn@latest
```

## Use

```bash
# In any folder, set up exercises here
golearn init

# Or clone a remote exercises repo
# golearn init <repo-url> <dir>
golearn init https://github.com/your-org/your-exercises my-exercises
cd my-exercises

# Learn
golearn list
golearn verify 01_hello
golearn progress
```

Need commands?

```bash
golearn help
```

## Licensing and Attribution

- Code: [MIT](./LICENSE)
- Non-code lesson content and any included gopher artwork: [CC BY 3.0](./CONTENT_LICENSE).
- Inspired by rustlings and Go by Example. [NOTICE](./NOTICE) for attributions.

Built with love for learning and knowledge sharing.
