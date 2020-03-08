package main

import (
	"fmt"
	"github.com/jschaare/kanji-browser/api/internal/app"
)

func main() {
	fmt.Println("Starting..")
	app := &app.App{}
	app.Init()
	app.Run(":3000")
}
