package main

import (
	"fmt"
	"time"
	"os"

	"github.com/gin-gonic/gin"
)

func (app *application) indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func (app *application) dbtestHandler(c *gin.Context) {
	var connectStatus string
    dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// 格式: user:password@tcp(host:port)/dbname?parseTime=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)

	db, err := openDB(dsn)
	if err != nil {
		fmt.Println("无法连接到数据库:", err)
	}else{
		fmt.Println("成功连接到数据库!")
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		connectStatus = fmt.Sprintf("数据库连接失败: %v", err)
	} else {
		connectStatus = "数据库连接成功"
		fmt.Println(time.Now(),"数据库连接成功!")
	}
	c.HTML(200, "index.html", gin.H{
		"status": connectStatus,
	})
}