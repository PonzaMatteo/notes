package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var fileName = flag.String("i", "", "name of the input file")
var srcDir = flag.String("d", "./", "directory where to look for the source code examples")

func main() {
	flag.Parse()
	input, err := Input(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	content, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	codePattern := regexp.MustCompile("!code\\((.*\\..*)\\)")
	filePattern := regexp.MustCompile("\\((.*)\\.(.*)\\)")
	includes := codePattern.FindAllString(string(content), -1)
	var substituteMap = make(map[string]string)
	for _, include := range includes {
		tokens := filePattern.FindStringSubmatch(include)
		fileName := fmt.Sprintf("%s.%s", tokens[1], tokens[2])
		content, err := ioutil.ReadFile(*srcDir + fileName)
		if err != nil {
			content = []byte("failed to read file to include: " + err.Error())
		}
		substituteMap[include] = CodeBlock(tokens[2], string(content))
	}

	var out = string(content)
	for key, value := range substituteMap {
		out = regexp.MustCompile(regexp.QuoteMeta(key)).ReplaceAllString(out, value)
	}
	fmt.Print(out)
}

func CodeBlock(ext, content string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("```%s\n", ext))
	builder.WriteString(content)
	builder.WriteString(fmt.Sprintf("```\n"))
	return builder.String()
}

func Input(fileName *string) (io.ReadCloser, error) {
	if fileName != nil && *fileName != "" {
		input, err := os.Open(*fileName)
		if err != nil {
			return nil, fmt.Errorf("failed to open input file: %v", err)
		}
		return input, nil
	}
	return os.Stdin, nil
}
