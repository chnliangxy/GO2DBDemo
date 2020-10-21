package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

const (
	server   = "localhost"
	port     = 1433
	user     = "admin"
	password = "admin"
	database = "TestLxy"
)

func main() {
	connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)

	var err error
	db, err = sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Connected!")

	//查询
	one, err := getOne(1001)
	// apps, err := getMany(1)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(one)
	// fmt.Println(apps)

	//更新
	one.name += " new"
	one.order++
	err = one.Update()
	if err != nil {
		log.Fatalln(err.Error())
	}
	a1, _ := getOne(1001)
	fmt.Println(a1)

	//插入
	newOne := app{
		ID:     1004,
		name:   "TestInsert",
		order:  1124,
		level:  10,
		status: 1,
	}

	err = newOne.Insert()
	if err != nil {
		log.Fatalln(err.Error())
	}

	a2, _ := getOne(newOne.ID)
	fmt.Println(a2)
}
