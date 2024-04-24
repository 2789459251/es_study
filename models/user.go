package models

import "time"

type User struct {
	ID        uint      `jsom:"id"`
	Name      string    `jsom:"name"`
	Nickname  string    `jsom:"nick_name"`
	CreatedAt time.Time `jsom:"created_at"`
}

func (u User) Index() string {
	return "user_index"
}
func (u User) Mapping() string {
	return `{
			"mappings":{
				"properties":{
					"nick_name":{
						"type":"text" //分词
					},
					"name":{
							"type":"keyword" //完全匹配
					},
					"id":{
							"type":"integer"
					},
					"created_at":{
							"type":"date",
							"null_value":"null",
							"format":"[yyyy-MM-dd HH:mm:ss]"
					}
				}
			}	
		}`
}
