package main

import (
	"database/sql"
	"log"
)

func getOne(id int) (a app, err error) {
	a = app{}
	db.QueryRow("select Id,Name,Status,[Level],[Order] from dbo.App where Id=@Id", sql.Named("Id", id)).Scan(&a.ID, &a.name, &a.status, &a.level, &a.order)
	return
}

func getMany(id int) (apps []app, err error) {
	rows, err := db.Query("select Id,Name,Status,[Level],[Order] from dbo.App where Id>@Id", sql.Named("Id", id))
	for rows.Next() {
		a := app{}
		err = rows.Scan(&a.ID, &a.name, &a.status, &a.level, &a.order)
		if err != nil {
			log.Fatalln(err.Error())
		}
		apps = append(apps, a)
	}
	return
}

func (a *app) Update() (err error) {
	_, err = db.Exec("update dbo.App set Name=@Name,[Order]=@Order where Id=@Id", sql.Named("Name", a.name), sql.Named("Order", a.order), sql.Named("Id", a.ID))
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}

func (a *app) Delete() (err error) {
	_, err = db.Exec("DELETE FROM dbo.App where Id=@Id", sql.Named("Id", a.ID))
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}

func (a *app) Insert() (err error) {
	sqlStr := "INSERT INTO dbo.App(Id,Name,Status,Level,[Order]) VALUES(@Id,@Name,@Status,@Level,@Order;);"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer stmt.Close()
	stmt.QueryRow(
		sql.Named("Id", a.ID),
		sql.Named("Name", a.name),
		sql.Named("Status", a.status),
		sql.Named("Level", a.level),
		sql.Named("Order", a.order))

	return
}
