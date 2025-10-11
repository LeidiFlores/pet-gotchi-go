# Pet-gotchi (Go + Ebiten + WASM)

A **super cute** virtual pet game built using **Go** and **Ebiten**. Great for Hacktoberfest and to show in your CV.

## Demo (WASM)
> Next to an active GitHub Pages, use `web/index.html` and the build `main.wasm` in the configured branch. (Optional: automated with GitHub Actions).

## Requirements
- Go 1.22+
- Module Ebiten v2

## Run (Desktop)
```bash
go run ./cmd/game
```

## Build (Desktop)
```bash
go build -o pet-gotchi ./cmd/game
```

## Compile to WASM
```bash
GOOS=js GOARCH=wasm go build -o web/main.wasm ./cmd/game
# Copiar wasm_exec.js de tu instalación de Go:
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/
# Abrir web/index.html en un servidor (por CORS):
# python3 -m http.server -d web 8080
```

## Structure
```
pet-gotchi-go/
├ cmd/game/main.go # Start
├ internal/game/game.go # Main loop + scene manager
├ internal/scene/menu.go # Menu scene
├ internal/scene/pet.go # Main "pet" scene
├ internal/ui/button.go # Minimal UI
├ web/index.html # WASM host
├ .github/workflows/build.yml # Basic CI
├─ CONTRIBUTING.md
├─ ROADMAP.md
└─ LICENSE
```

## Contribute
Read `CONTRIBUTING.md` and show the issues marked as `good first issue` and `help wanted`.

---

© MIT
