package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func main() {
	sqlQuery()
	// sqlPrepare()
	// sqlExec()
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "learner:secret@tcp(localhost:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, err
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select id, name, age, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
	}

	var result []student
	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.age, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	stmt.QueryRow("E001").Scan(&result1.name, &result1.age)
	fmt.Printf("name: %v\ngrade: %v\n", result1.name, result1.age)

	var result2 = student{}
	stmt.QueryRow("W001").Scan(&result2.name, &result2.age)
	fmt.Printf("name: %v\ngrade: %v\n", result2.name, result2.age)

	var result3 = student{}
	stmt.QueryRow("B001").Scan(&result3.name, &result3.age)
	fmt.Printf("name: %v\ngrade: %v\n", result3.name, result3.age)
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values(?, ?, ?, ?)", "G001", "Galahad", 20, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")

	_, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")
}
