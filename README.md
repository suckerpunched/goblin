# goblin

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/suckerpunched/goblin)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/suckerpunched/goblin?label=Version)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/suckerpunched/goblin?label=Release)
[![GitHub Actions](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fatrox%2Fsync-dotenv%2Fbadge&label=build&logo=none)](https://actions-badge.atrox.dev/suckerpunched/goblin/goto)

[![Go Reference](https://pkg.go.dev/badge/github.com/suckerpunched/goblin.svg)](https://pkg.go.dev/github.com/suckerpunched/goblin)
[![Go Report Card](https://goreportcard.com/badge/github.com/nanobox-io/golang-scribble)](https://goreportcard.com/report/github.com/suckerpunched/goblin)

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

##### Delete
```go
// delete single resource
err = db.Delete("bank", "tom")
if err != nil {
  panic(err)
}

// delete entire collection
err = db.Delete("bank", "")
if err != nil {
  panic(err)
}
```