package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lithammer/shortuuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type users struct {
	UniqueID    string `gorm:"primaryKey"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber int    `json:"phone_number"`
}

//NewUser function ..
func NewUser(w http.ResponseWriter, r *http.Request) {
	connect()
	var user users
	json.NewDecoder(r.Body).Decode(&user)
	id := shortuuid.New()
	user.UniqueID = id
	DB.Create(&user)
	json.NewEncoder(w).Encode(&user)
}

//GetUser function..
func GetUser(w http.ResponseWriter, r *http.Request) {
	connect()
	var user []users
	DB.Find(&user)
	json.NewEncoder(w).Encode(&user)

}

//SearchUser function..
func SearchUser(w http.ResponseWriter, r *http.Request) {
	connect()
	vars := mux.Vars(r)
	userID := vars["id"]
	var user users
	DB := DB.Where("unique_id = ?", userID).Find(&user)
	DB.Find(&user)
	json.NewEncoder(w).Encode(&user)
}

//LoginEndPoint function..
func LoginEndPoint(w http.ResponseWriter, r *http.Request) {
	connect()
	vars := mux.Vars(r)
	usrname := vars["username"]
	pasword := vars["password"]
	var user users
	DB.Where("username= ? AND password= ?", usrname, pasword).Find(&user)
	DB.Find(&user)
	json.NewEncoder(w).Encode(&user)
}

//EditUser function..
func EditUser(w http.ResponseWriter, r *http.Request) {
	connect()
	vars := mux.Vars(r)
	userID := vars["id"]
	var getUser users
	DB.First(&getUser)
	json.NewDecoder(r.Body).Decode(&getUser)
	DB.Model(&users{}).Where("unique_id=?", userID).Update("name", getUser.Name)
	DB.Model(&users{}).Where("unique_id=?", userID).Update("email", getUser.Email)
	DB.Model(&users{}).Where("unique_id=?", userID).Update("username", getUser.Username)
	DB.Model(&users{}).Where("unique_id=?", userID).Update("password", getUser.Password)
	DB.Model(&users{}).Where("unique_id=?", userID).Update("date_of_birth", getUser.DateOfBirth)
	DB.Model(&users{}).Where("unique_id=?", userID).Update("phone_number", getUser.PhoneNumber)

	DB := DB.Where("unique_id = ?", userID).Find(&getUser)
	DB.Find(&getUser)
	userDetails, _ := &getUser, DB
	res, _ := json.Marshal(&userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
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
