# simple_MYSQL_CRUD_API_With_Golang
简单的使用mysql进行一个图书管理的crud api设计


api:
```
GET /book 获取全部书  json格式
GET /book/{bookId} 获取指定id书 json格式
PUT /book/{bookId} 更新指定id书 json格式
DELETE /book/{bookId} 删除指定id书 json格式
POST /book/ 创建一本书 json格式
```