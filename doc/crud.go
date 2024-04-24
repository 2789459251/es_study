package doc

import (
	"context"
	"es_study/models"
	"es_study/utils"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

func CreateDoc() {
	user := models.User{
		ID:        10,
		Name:      "Zy",
		Nickname:  "Zy爱吃蘑菇和鸡蛋",
		CreatedAt: time.Now(),
	}

	IndexResponse, err := utils.Client.Index().Index(user.Index()).BodyJson(user).Do(context.Background())
	if err != nil {
		fmt.Println("创建文档错误：", err)
		return
	}
	fmt.Printf("%#v", IndexResponse)
}

func DeleteDoc() {
	deleteResponse, err := utils.Client.Delete().Index(models.User{}.Index()).Id("YLUbD48BsTBsV_jmmbJQ").Refresh("true").Do(context.Background())
	if err != nil {
		fmt.Println("创建文档错误：", err)
		return
	}
	fmt.Printf("%#v", deleteResponse)
}
func DeleteDocs() {
	idList := []string{
		"YrUeD48BsTBsV_jmhbI5",
		"Y7UeD48BsTBsV_jml7If",
	}
	bulk := utils.Client.Bulk().Index(models.User{}.Index()).Refresh("true")
	for _, s := range idList {
		req := elastic.NewBulkDeleteRequest().Id(s)
		bulk.Add(req)
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println("创建文档错误：", err)
		return
	}
	fmt.Println(res.Succeeded())
}

func CreateDocs() {
	userList := []models.User{
		{
			ID:        10,
			Name:      "Zy",
			Nickname:  "asdfghjkl;",
			CreatedAt: time.Now(),
		},
		{
			ID:        11,
			Name:      "Sy",
			Nickname:  "qwertyuio",
			CreatedAt: time.Now(),
		},
		{
			ID:        12,
			Name:      "Qz",
			Nickname:  "zxcbnm,",
			CreatedAt: time.Now(),
		},
	}
	bulk := utils.Client.Bulk().Index(models.User{}.Index()).Refresh("true")
	for _, user := range userList {
		req := elastic.NewBulkCreateRequest().Doc(user)
		bulk.Add(req)
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println("创建文档错误：", err)
		return
	}
	fmt.Println(res.Succeeded())
}
