# Contributing Guide

¡Gracias por querer contribuir a **Pet-gotchi**! Buscamos PRs pequeñas y enfocadas.

## Cómo empezar
1. Forkea el repo y crea una rama desde `main`.
2. Instala dependencias y corre el juego localmente:
   ```bash
   go run ./cmd/game
   ```
3. Elige un issue con `good first issue` o `help wanted`.

## Estilo
- Go 1.22+, `go fmt ./...` antes de commitear.
- Commits: tipo breve en imperativo (e.g., `feat: add wink animation`).

## PRs
- Enfoca tu PR en **una cosa**.
- Incluye capturas/GIFs si cambias UI.
- Describe el **por qué** y el **cómo**.

## Ideas para empezar
- Nuevos emotes (😺💤🍖).
- SFX al alimentar/cepillar.
- Accesibilidad (lectura de teclas, tamaño de texto).
- Traducciones (es/en) mediante archivos `.json`.
- Balance de stats en `pet.go` (constantes).

¡Gracias!
