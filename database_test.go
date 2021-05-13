package belajargolangdatabase

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql" //! ada underscore harapannya memanggil method initnya
)

/*
! PASTIKAN HARUS MENGIMPORT DRIVER
*/

func TestDatabaseOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:user1234@tcp(localhost:3306)/belajar_golang_database")

	if err != nil {
		panic(err)
	}
	fmt.Println("Success Connected")
	defer db.Close()
}