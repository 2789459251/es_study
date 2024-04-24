package main

import (
	"es_study/doc"
	"es_study/utils"
)

func main() {
	utils.EsConnect()
	//doc.CreateDoc()
	//doc.CreateDocs()
	//doc.FindDocExact()
	doc.UpdateDoc()
}
