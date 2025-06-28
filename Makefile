run:
	go run cmd/main.go
wasm:
	env GOOS=js GOARCH=wasm go build -o main.wasm main.go
dist: wasm
	cp main.wasm dist/main.wasm
	zip -r unibun.zip dist