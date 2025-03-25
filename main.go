package main

import (
	"fmt"
	"github.com/Sheyiyuan/ChronoMind/api"
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
	api.NewAPICore(8080, true)
}
