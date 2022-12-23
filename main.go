package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func authMiddleHandler(context *gin.Context) {
	cookie, err := context.Cookie("hyj")
	if err != nil {
		return
	}
	if cookie == "qwerpvp" {
		context.Next()
	} else {
		context.Abort()
	}
}
func main() {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/new")
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping err:", err)
		return
	}

	rows, err := db.Query("select * from student where id=?", 1)
	if err != nil {
		log.Println(rows)
		panic(err)
	}
	defer rows.Close()

}
