package belajargolangdatabase

import (
	"context"
	"fmt"
	"testing"
	"time"
)


func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('eko', 'EKO', 'eko@mail.com', 1000000, 90.0, '1999-10-10', true), ('budi', 'BUDI', 'budi@mail.com', 10000000, 85.5, '1998-10-10', true),('joko', 'JOKO', 'joko@mail.com', 2000000, 70.0, '1997-05-10', false);"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() { //! untuk mengambil data selanjutnya
		var id, name string
		err = rows.Scan(&id, &name) //! kenapa menggunakan pointer, karena nanti di dalanm scan dia akan set dari parameternya
		if err != nil {
			panic(err)
		}

		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() { //! untuk mengambil data selanjutnya
		var id, name, email string
		var balance int32
		var rating float64
		var birthDate, created_at time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &created_at) //! kenapa menggunakan pointer, karena nanti di dalanm scan dia akan set dari parameternya
		if err != nil {
			panic(err)
		}

		fmt.Println("==============================")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		fmt.Println("Email: ", email)
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		fmt.Println("Birth Date: ", birthDate)
		fmt.Println("Married: ", married)
		fmt.Println("Created_at: ", created_at)
	}
}
