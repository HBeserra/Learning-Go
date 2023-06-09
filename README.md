# Learning Go

Leaning resources:

- [Udemy: Go (Golang) Programming](https://www.udemy.com/course/master-go-programming-complete-golang-bootcamp)

- [Book: The Go Programming Language Book by Alan A. A. Donovan and Brian Kernighan](https://www.amazon.com.br/Go-Programming-Language-Brian-Kernighan/dp/0134190440)



# Installing the Go

Setup the Go Compiler and basic tooling:

Ubuntu 16.04+
```
$ sudo apt update
$ sudo apt install golang-go -y
```

[Go Extension for VSCode](https://marketplace.visualstudio.com/items?itemName=golang.Go)


# Running and Compiling

`go run <program_name>` 

`go build <program_name>`

Windows: `GOOS=windows GOARCH=amd64 go build`

Linux: `GOOS=linux GOARCH=amd64 go build`

Mac: `GOOS=darwin GOARCH=amd64` use `arm64` for m1 or newest processors

`-o` changes the output program name.


