package main

import (
	"fmt"

	"github.com/ektagarg/gin-gorm-todo-app/Config"
	"github.com/ektagarg/gin-gorm-todo-app/Models"
	"github.com/ektagarg/gin-gorm-todo-app/Routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 添加 MySQL dialect
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
		// return
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	// running
	r.Run()
}
