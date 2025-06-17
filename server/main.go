package main

import (
	"fmt"
	"github.com/jessehorne/endlessskyonline/server/pkg"
)

func main() {
	gs, err := pkg.NewGameServer("localhost:5050")
	if err != nil {
		fmt.Println(err)
		return
	}

	gs.SetLogging(true)
	err = gs.Start()
	if err != nil {
		fmt.Println(err)
	}
}
