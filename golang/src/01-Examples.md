# 01 - Some Examples
Before starting to dive into the language let's give a look to some simple program.

## Hello World
!code(examples/hello.go)

## HTTP Server
!code(example/01_http_server.go)

## HTTP Client
!code(example/01_http_client.go)

## Another HTTP Server
!code(example/02_http_server.go)

## Another HTTP Client
!code(example/02_http_client.go)

## A CLI Tool
Let's see a program that allow to process a markdown file and include some source code files as code blocks. The placeholder will be in the at the beginning of a new line with the format `!code(file_name)`
!code(examples/include_code.go)

