package main

import (
	"es_study/index"
	"es_study/utils"
)

func main() {
	utils.EsConnect()
	index.CreateIndex()
}
