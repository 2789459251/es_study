package utils

import (
	"fmt"
	"github.com/olivere/elastic/v7"
)

var Client *elastic.Client

func EsConnect() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("elastic", "123456"))

	if err != nil {
		fmt.Println("连接失败:", err.Error())
		return
	}
	Client = client
}
