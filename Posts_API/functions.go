package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lithammer/shortuuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type posts struct {
	UniqueID  string `gorm:"primaryKey"`
	AuthorId  string `json:"author_id"`
	PostedOn  string `json:"posted_on"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Thumbnail string `json:"thumbnail"`
}

//NewPost function ..
func NewPost(w http.ResponseWriter, r *http.Request) {
	connect()
	var post posts
	json.NewDecoder(r.Body).Decode(&post)
	id := shortuuid.New()
	post.UniqueID = id
	post.PostedOn = time.Now().Format("2006-01-02 15:04:05")
	DB.Create(&post)
	json.NewEncoder(w).Encode(&post)
}

//GetPost function..
func GetPost(w http.ResponseWriter, r *http.Request) {
	connect()
	var post []posts
	DB.Order("posted_on ASC").Find(&post)
	json.NewEncoder(w).Encode(&post)

}

//EditPost function..
func EditPost(w http.ResponseWriter, r *http.Request) {
	connect()
	vars := mux.Vars(r)
	userID := vars["id"]
	var getpost posts
	err := DB.First(&getpost).Error
	if err != nil {
		fmt.Println("Post not found")
	}
	json.NewDecoder(r.Body).Decode(&getpost)
	DB.Model(&posts{}).Where("unique_id=?", userID).Update("author_id", getpost.AuthorId)
	DB.Model(&posts{}).Where("unique_id=?", userID).Update("posted_on", getpost.PostedOn)
	DB.Model(&posts{}).Where("unique_id=?", userID).Update("title", getpost.Title)
	DB.Model(&posts{}).Where("unique_id=?", userID).Update("body", getpost.Body)
	DB.Model(&posts{}).Where("unique_id=?", userID).Update("thumbnail", getpost.Thumbnail)

	DB := DB.Where("unique_id = ?", userID).Find(&getpost)
	DB.Find(&getpost)
	postDetails, _ := &getpost, DB
	res, _ := json.Marshal(&postDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

//DeletePost function..
func DeletePost(w http.ResponseWriter, r *http.Request) {
	connect()
	vars := mux.Vars(r)
	postId := vars["id"]
	var post []posts
	err := DB.Where("unique_id =?", postId).Delete(&post).Error
	if err != nil {
		fmt.Println("Db error")
	}
	json.NewEncoder(w).Encode(post)
}

//DB Connection funcion..
func connect() {
	dsn := "host=localhost user=postgres password=yuvi@123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can not connect to the database")
	}
	DB = db

}
