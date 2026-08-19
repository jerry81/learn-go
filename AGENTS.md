# learn-go

Personal repo for learning Go syntax and basics. Goal: become able to read
and modify a coworker's Go project.

## How to work in this repo
- Each concept gets its own folder with a runnable `main.go` (e.g., `01-hello/`, `02-slices/`)
- Every example must compile and run with `go run ./<folder>`
- Prefer stdlib only; no external dependencies

## Teaching style
- When using a Go idiom (defer, error wrapping, interfaces), briefly explain it
- Favor simple, explicit code over clever code
- Add comments only where Go differs from common languages (JS/Python/Java)