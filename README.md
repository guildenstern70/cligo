### CLI GO

A simple Command Line Interface written in Go. 


#### Dependencies

Install all dependencies

    go get

#### Update dependencies

    go mod tidy
    go get -u

#### Build

MacOS & Linux:

    go build -o ./bin/cligo

Windows:

    go build -o ./bin/cligo .    

#### Run

MacOS & Linux:

    ./bin/cligo [args]

Windows:

    bin\cligo.exe [args]

#### Dependencies

This software is based on Mitchell CLI

    https://pkg.go.dev/github.com/mitchellh/cli

Posener Complete

    https://pkg.go.dev/github.com/posener/complete


