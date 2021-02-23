# goblin

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/suckerpunched/goblin)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/suckerpunched/goblin?label=Version)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/suckerpunched/goblin?label=Release)


[![Go Reference](https://pkg.go.dev/badge/github.com/suckerpunched/goblin.svg)](https://pkg.go.dev/github.com/suckerpunched/goblin)
[![Go Report Card](https://goreportcard.com/badge/github.com/nanobox-io/golang-scribble)](https://goreportcard.com/report/github.com/nanobox-io/golang-scribble)
--------

short description here...

## Getting Started

### Installation
To start using `goblin`, install Go and run `go get`:
```
$ go get github.com/suckerpunched/goblin
```

### Usage
##### Open a Database
```go
db, err = goblin.New("./data", nil)
if err != nil {
  panic(err)
}
```

##### Available Options
```go
&Options{
  Format: ["json" | "gob"],
  Compression: ["gzip"],
  Backend: ["local"],
}
```

##### Write
```go
err = db.Write("bank", "tom", &x)
if err != nil {
  panic(err)
}
```

##### Read
```go
err = db.Read("bank", "tom", &y)
if err != nil {
  panic(err)
}
```