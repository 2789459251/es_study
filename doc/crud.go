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
			ID:        19,
			Name:      "Zy",
			Nickname:  "asdfghjkl;",
			CreatedAt: time.Now(),
		},
		{
			ID:        120,
			Name:      "Sy",
			Nickname:  "qwertyuio",
			CreatedAt: time.Now(),
		},
		{
			ID:        21,
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

func FindDoc() {
	limit := 2
	page := 4
	from := (page - 1) * limit
	query := elastic.NewBoolQuery()
	reslist, err := utils.Client.Search(models.User{}.Index()).Query(query).From(from).Size(limit).Do(context.Background())
	if err != nil {
		fmt.Println("查询文档列表错误：", err)
		return
	}
	count := reslist.Hits.TotalHits.Value
	fmt.Println("查到的数量：", count)
	for _, hit := range reslist.Hits.Hits {
		fmt.Println(string(hit.Source))
	}
}

// 精确匹配是指keyword来匹配
func FindDocExact() {
	limit := 2
	page := 1
	from := (page - 1) * limit
	query := elastic.NewTermQuery("Nickname.keyword", "Zy爱吃蘑菇和鸡蛋")
	reslist, err := utils.Client.Search(models.User{}.Index()).Query(query).From(from).Size(limit).Do(context.Background())
	if err != nil {
		fmt.Println("查询文档列表错误：", err)
		return
	}
	count := reslist.Hits.TotalHits.Value
	fmt.Println("查到的数量：", count)
	for _, hit := range reslist.Hits.Hits {
		fmt.Println(string(hit.Source))
	}
}

func UpdateDoc() {
	updateRes, err := utils.Client.Update().Index(models.User{}.Index()).Id("ZLUeD48BsTBsV_jmurI9").Doc(map[string]interface{}{
		"Name": "ty",
	}).Do(context.Background())
	if err != nil {
		fmt.Println("更新文档错误：", err)
		return
	}
	fmt.Println(updateRes)
}
