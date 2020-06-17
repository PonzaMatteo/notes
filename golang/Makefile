BIN = ./bin

make:
	go build -o $(BIN)/include_code  -i ./src/examples/include_code.go
	rm readme.md
	ls ./src/*.md | xargs cat | $(BIN)/include_code -d ./src/ > readme.md
	pandoc -s --toc -c asset/pandoc.css readme.md  -o readme.html
bin:
	mkdir -p $(BIN)
