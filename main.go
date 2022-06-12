package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id        int
	Name      string
	Nickname  string
	Homepage  string
	Introduce string
}

var DB, _ = Init()

func Init() (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/WHCTF?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	return
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", func(c *gin.Context) {
		filter := make(map[string]interface{})
		users := []User{}
		json.Unmarshal([]byte(c.Query("filter")), &filter)
		fmt.Println(filter)
		DB.Select("*").Table("user").Where(filter).Find(&users)
		c.HTML(http.StatusOK, "index.html", users)
	})
	r.Run("0.0.0.0:1337")
}
