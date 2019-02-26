GOROOT=`go env GOROOT`

all: main.wasm wasm_exec.js index.html

main.wasm: main.go
	GOOS=js GOARCH=wasm go build -o main.wasm main.go

wasm_exec.js:
	cp "${GOROOT}/misc/wasm/wasm_exec.js" .

index.html: index.py
	python index.py > index.html

clean:
	rm index.html wasm_exec.js main.wasm

serve:
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
