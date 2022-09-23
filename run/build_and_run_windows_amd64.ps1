cd "$Env:GOPATH/src/go-wasm-template";
$Env:GOOS = "js"; $Env:GOARCH = "wasm"; go build -o static/main.wasm wasm/main.go;
Copy-Item "$Env:GOROOT/misc/wasm/wasm_exec.js" -Destination static/wasm_exec.js;
$Env:GOOS = "windows"; $Env:GOARCH = "amd64"; go run ./main.go;