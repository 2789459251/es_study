package index

import (
	"context"
	"es_study/models"
	"es_study/utils"
	"fmt"
)

func CreateIndex() {
	index := "user_index"
	//todo 自定义index的name，以及mapping
	if Isexist(index) {
		fmt.Println("已存在", index, "索引，执行删除操作！")
		//索引删除
		DeleteIndex(index)
	}
	createIndex, err := utils.Client.CreateIndex(index).BodyString(models.User{}.Mapping()).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(createIndex, "索引创建成功！")
}

// 是否存在
func Isexist(indexName string) bool {
	exist, err := utils.Client.IndexExists("user_index").Do(context.Background())
	if err != nil {
		fmt.Println("判断索引是否存在的方法出错：", err)
	}
	return exist
}

func DeleteIndex(indexName string) {
	deleteindex, err := utils.Client.DeleteIndex("user_index").Do(context.Background())
	if err != nil {
		fmt.Println("删除索引：", err.Error())
		return
	}
	fmt.Println(deleteindex, "删除索引成功")
}
