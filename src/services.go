package main

import (
	"database/sql"
)

func getOne(id int) (a app, err error) {
	a = app{}
	db.QueryRow("select Id,Name,Status,[Level],[Order] from dbo.App where Id=@Id", sql.Named("Id", id)).Scan(&a.ID, &a.name, &a.status, &a.level, &a.order)
	return
}
