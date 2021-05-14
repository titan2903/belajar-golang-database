package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var created_at time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &created_at) //! kenapa menggunakan pointer, karena nanti di dalanm scan dia akan set dari parameternya
		if err != nil {
			panic(err)
		}

		fmt.Println("==============================")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date: ", birthDate.Time)
		}
		fmt.Println("Married: ", married)
		fmt.Println("Created_at: ", created_at)
	}
}


func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password +"' LIMIT 1"
	fmt.Println("sript: ", script)
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username) //! pada bagian func Scan harus berupa pointer
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	sqlQuery := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1" //! aman dari sql injection
	fmt.Println("sript: ", sqlQuery)
	rows, err := db.QueryContext(ctx, sqlQuery, username, password) //! bisa menambahkan sesuai degnan jumlah parameter tanda tanya

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username) //! pada bagian func Scan harus berupa pointer
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login", username)
	} else {
		fmt.Println("gagal login")
	}
}


func TestExecSqlSafeParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "titan'; DROP TABLE user; #"
	password:= "titan"

	script := "INSERT INTO user(username, password) VALUES(?, ?);" //! jika tidak menggunakan query sql seperti ini, user bisa saja melakukan SQL injection
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}


func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "user@gmail.com"
	comment := "test comment"

	script := "INSERT INTO comments(email, comment) VALUES(?, ?);" //! jika tidak menggunakan query sql seperti ini, user bisa saja melakukan SQL injection
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	inserId, err := result.LastInsertId() //! Jika ingin memasukkan id yang auto increment dan di konversi menjadi int64

	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new commment with id", inserId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?, ?);"

	stmt, err := db.PrepareContext(ctx, script) //! jika ingin melakukan query berkali kali yang sama terus dan parameter berbeda beda, lebih baik menggunakan prepare statement

	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	/*
		! statement ketika di buat sudah bimbing ke sql scriptnya dan tidak bisa di rubah lagi
		! tidak perlu menyakan lagi ke DB poolnya lagi karena sudah include koneksi ke DB
		! prepare statement bisa di gunakan untuk query dan exec
	*/


	for i := 0; i < 10; i++ {
		email := "eko@" + strconv.Itoa(i) + "@gmai.com"
		comment := "komentar ke " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId() //! tidak perlu lagi memasukkan sql scriptnya karena sudah otomatis di bimbing
		if err != nil {
			panic(err)
		}

		fmt.Println("comment id ", lastInsertId)
	}
}


func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin() //! memulai transaction
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?, ?);"
	//! do transaction
	for i := 0; i < 10; i++ {
		email := "budi@" + strconv.Itoa(i) + "@gmai.com"
		comment := "komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment) //! dari tx menggunakan execContext, prepare statement atau pun query

		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId() //! tidak perlu lagi memasukkan sql scriptnya karena sudah otomatis di bimbing
		if err != nil {
			panic(err)
		}

		fmt.Println("comment id ", lastInsertId)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
