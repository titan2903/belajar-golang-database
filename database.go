package belajargolangdatabase

import (
	"database/sql"
	"time"
)


func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:user1234@tcp(localhost:3306)/belajar_golang_database?parseTime=true")

	if err != nil {
		panic(err)
	}
	
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

/*
! tidak perlu lagi open connection baru di tiap unit test
? parseTime=true, agar otomatis melakukan conversi dari []uint8 ke time
*/