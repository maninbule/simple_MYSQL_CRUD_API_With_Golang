package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/simple_MYSQL_CRUD_API_With_Golang/pkg/models"
	"github.com/simple_MYSQL_CRUD_API_With_Golang/pkg/utils"
	"net/http"
	"strconv"
)

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allBook := models.GetAllBookDB()
	json.NewEncoder(w).Encode(allBook)
}

func GetABookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idstr, ok := vars["bookId"]
	if !ok {
		json.NewEncoder(w).Encode("没有传递id")
		return
	}
	id, err := strconv.ParseUint(idstr, 0, 0)
	if err != nil {
		json.NewEncoder(w).Encode("id非法")
		return
	}
	book, ok := models.GetBookByID(uint(id))
	if !ok {
		json.NewEncoder(w).Encode("不存在指定id的书")
		return
	}
	json.NewEncoder(w).Encode(book)
}

func CreateABook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book := &models.Book{}
	if ok := utils.ParseBody(r, book); !ok {
		json.NewEncoder(w).Encode("json数据格式不正确")
		return
	}
	book.CreateBookDB()
	json.NewEncoder(w).Encode("创建成功")
}

func DeleteABookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idstr, ok := vars["bookId"]
	if !ok {
		json.NewEncoder(w).Encode("没有传递id")
		return
	}
	id, err := strconv.ParseUint(idstr, 0, 0)
	if err != nil {
		json.NewEncoder(w).Encode("id非法")
		return
	}
	ok = models.DeleteBookById(uint(id))
	if !ok {
		json.NewEncoder(w).Encode("不存在指定的id")
		return
	} else {
		json.NewEncoder(w).Encode("删除成功")
		return
	}

}

func UpdateABookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idstr, ok := vars["bookId"]
	if !ok {
		json.NewEncoder(w).Encode("没有传递id")
		return
	}
	id, err := strconv.ParseUint(idstr, 0, 0)
	if err != nil {
		json.NewEncoder(w).Encode("id非法")
		return
	}
	book := &models.Book{}
	if ok := utils.ParseBody(r, book); !ok {
		json.NewEncoder(w).Encode("json数据格式不正确")
		return
	}
	ok = models.UpdateById(uint(id), *book)
	if !ok {
		json.NewEncoder(w).Encode("指定id不存在")
	} else {
		json.NewEncoder(w).Encode("更新成功")
	}

}
