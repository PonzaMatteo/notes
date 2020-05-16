BIN=./bin

make:
	go build -o $BIN/include_code  -i ./src/examples/include_code.go
	ls ./src/*.md |  xargs cat | $BIN/include_code -d ./src/ > readme.md
	tree
bin:
	mkdir -p $BIN
