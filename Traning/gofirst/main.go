package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"gofirst/compoment"
	_ "gofirst/compoment"
	"gofirst/modules/restaurant/restauranttransport/ginrestaurant"
	_ "gofirst/modules/restaurant/restauranttransport/ginrestaurant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	_ "net/http"
	"os"
	"strconv"

	_ "strconv"
)

type Note struct {
	Id      int    `json: "id, omitempty" gorm:"column:id;"`
	Title   string `json: "title" gorm:"column:title;"`
	Content string `json: "content" gorm:"column:content;"`
}

type NoteUpdate struct {
	Title *string `json: "title" gorm:"column:title;"`
}

func (Note) TableName() string {
	return "note"
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	var dsn = os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	var notes []Note
	db.Where("status = ?", 1).Find(&notes)
	fmt.Println(notes)

	if err := db.Where("id=3").Find(&notes); err != nil {
		log.Println(err)
	}

	var note Note
	if err := db.Where("id=3").First(&note); err != nil {
		log.Println(err)
	}

	//note.Title = ""
	//db.Table(Note{}.TableName()).where("id =3").Updates(&note)
	//fmt.Println(note)

	newTitle := "demo 3"
	db.Table(Note{}.TableName()).Where("id=3").Updates(&NoteUpdate{Title: &newTitle})
	fmt.Println(note)
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	c.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	appCtx := compoment.NewAppContext(db)
	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	}
	return r.Run()
}

// POST, GET, PUT, PATCH, DELETE
// POST: create ở server,
//get:Lấy dữ liêu,
//server đẩy dữ liệu về cho mình,
//Put: thay đổi toàn bộ object,
// patch thay đổi một phần, delete: muốn xóa đi ở server
