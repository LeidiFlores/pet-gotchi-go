# Pet Gotchi Go 🐾

> A virtual pet game built with Go + Ebiten, playable in the browser via WebAssembly.

Raise, feed, and interact with your pixel companion — all in real time. This project was built to explore 2D game development in Go and WebAssembly compilation for browser delivery.

---

## Demo

> **Try it live:** [GitHub Pages link](#) _(deploy via the WASM build — see below)_

---

## Features

- **Real-time game loop** — powered by Ebiten's tick-based engine
- **Scene system** — menu and pet scenes managed by a central scene manager
- **WASM compilation** — play directly in the browser, no install needed
- **CI/CD pipeline** — GitHub Actions workflow for automated builds
- **Open to contributions** — Hacktoberfest-friendly with labeled issues

---

## Tech stack

| Layer | Technology |
|---|---|
| Language | Go 1.22+ |
| Game engine | Ebiten v2 |
| Rendering target | Desktop + WebAssembly |
| CI | GitHub Actions |

---

## Getting started

### Prerequisites
- Go 1.22+
- Ebiten v2 (`go get github.com/hajimehoshi/ebiten/v2`)

### Run on desktop
```bash
go run ./cmd/game
```

### Build for desktop
```bash
go build -o pet-gotchi ./cmd/game
```

### Build for browser (WASM)
```bash
GOOS=js GOARCH=wasm go build -o web/main.wasm ./cmd/game
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/
python3 -m http.server -d web 8080
# then open http://localhost:8080
```

---

## Project structure

```
pet-gotchi-go/
├── cmd/game/           # Entry point
├── internal/
│   ├── game/           # Main loop + scene manager
│   ├── scene/          # Menu and pet scenes
│   └── ui/             # Button components
├── web/                # WASM host (index.html)
├── assets/             # Sprites and game assets
├── .github/workflows/  # CI pipeline
├── CONTRIBUTING.md
└── ROADMAP.md
```

---

## Roadmap

See [ROADMAP.md](ROADMAP.md) for planned features. Check [open issues](../../issues) for contribution opportunities tagged `good first issue`.

---

## Contributing

Read [CONTRIBUTING.md](CONTRIBUTING.md) to get started. All skill levels welcome — this project participates in Hacktoberfest.

---

## License

MIT © [LeidiFlores](https://github.com/LeidiFlores)