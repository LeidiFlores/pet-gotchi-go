# Pet-gotchi (Go + Ebiten + WASM)

A **super cute** game like virtual pet build with **Go** using **Ebiten**. Thinking for Hacktoberfest and showed in your CV.

## Demo (WASM)
> Next to active GitHub Pages, use `web/index.html` and the build `main.wasm` in the configurated branch. (Optional: automated with Github Actions).

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
├─ cmd/game/main.go           # arranque
├─ internal/game/game.go      # loop principal + gestor de escenas
├─ internal/scene/menu.go     # escena del menú
├─ internal/scene/pet.go      # escena principal "mascota"
├─ internal/ui/button.go      # UI mínima
├─ web/index.html             # host WASM
├─ .github/workflows/build.yml# CI básico
├─ CONTRIBUTING.md
├─ ROADMAP.md
└─ LICENSE
```

## Contribute
Read `CONTRIBUTING.md` and show the issues marked as `good first issue` and `help wanted`.

---

© MIT
