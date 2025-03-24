package main

import (
	"fmt"
	"github.com/Sheyiyuan/ChronoMind/api_core"
	"github.com/Sheyiyuan/ChronoMind/logos"
)

func init() {
	err := initCore()
	if err != nil {
		panic(err)
	}
	logos.InitLog()
}

func main() {
	fmt.Printf("Ciallo～(∠・ω< )⌒☆, World!\n")
	api_core.NewAPICore(8080, true)
}
