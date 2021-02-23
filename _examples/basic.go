package main

import (
	"fmt"

	"github.com/suckerpunched/goblin"
)

type gold struct {
	Weight int
}

var (
	db *goblin.Database

	x = gold{2}
	y = gold{}
)

func init() {
	var err error

	db, err = goblin.New("./data", nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error

	err = db.Write("bank", "tom", &x)
	if err != nil {
		panic(err)
	}

	err = db.Read("bank", "tom", &y)
	if err != nil {
		panic(err)
	}

	fmt.Println(y)
}
