# Pet-gotchi (Go + Ebiten + WASM)

Un juego **super cute** tipo mascota virtual construido en **Go** usando **Ebiten**. Pensado para Hacktoberfest y para lucir en tu CV.

## Demo (WASM)
> Después de activar GitHub Pages, sube `web/index.html` y la build `main.wasm` al branch configurado. (Opcional: automatiza con GitHub Actions).

## Requisitos
- Go 1.22+
- Módulo Ebiten v2

## Ejecutar (Desktop)
```bash
go run ./cmd/game
```

## Compilar (Desktop)
```bash
go build -o pet-gotchi ./cmd/game
```

## Compilar a WASM
```bash
GOOS=js GOARCH=wasm go build -o web/main.wasm ./cmd/game
# Copiar wasm_exec.js de tu instalación de Go:
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/
# Abrir web/index.html en un servidor (por CORS):
# python3 -m http.server -d web 8080
```

## Estructura
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

## Contribuir
Lee `CONTRIBUTING.md` y mira los issues marcados como `good first issue` y `help wanted`.

---

© MIT
