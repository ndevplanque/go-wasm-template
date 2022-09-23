cd "$(go env GOPATH)/src/go-wasm-template"
GOOS=js GOARCH=wasm go build -o static/main.wasm wasm/main.go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" static/
GOOS=linux GOARCH=amd64 go run ./main.go
