# Contributing Guide

Thank you for your interest in contributing to **Pet-gotchi**! We are looking for small, focused PRs.

## How to get started:

1. Fork the repository and create a branch from `main`.
2. Install dependencies and run the game locally:

   ```bash
   go run ./cmd/game
   ```

3. Choose an issue with `good first issue` or `help wanted`.

## Style:

- Use Go 1.22+ and run `go fmt ./...` before committing.
- Commits: imperative short type (e.g. `feat: add wink animation`).

## PRs:

- Focus your PR on **one thing**.
- Add screenshots or GIFs if you change the UI.
- Describe the `why` and the `how`.

## Ideas to get started:

- New emojis (e.g. 😺, 💤, 🍖).
- SFX when feeding/brushing.
- Accessibility features such as keyboard reading and text size.
- Translations (es/en) using .json files.
- Balance stats in `pet.go` (constants).

Thanks! ✨
