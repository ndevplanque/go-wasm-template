const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch("./static/main.wasm"), goWasm.importObject)
    .then((result) => {
        goWasm.run(result.instance)

        document.getElementById("get-html").addEventListener("click", () => {
            document.body.innerHTML += getHtml()
        })
    })