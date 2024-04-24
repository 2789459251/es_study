package main

import (
	"es_study/utils"
	"fmt"
)

func main() {
	utils.EsConnect()
	fmt.Println(utils.Client)
}
