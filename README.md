tinygo build -o plugin.wasm -scheduler=none -target=wasi --no-debug plugin.go
go test -v ./