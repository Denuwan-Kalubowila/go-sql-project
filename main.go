package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Student struct {
	ID    int
	Fname string
	Lname string
	Age   int
}

func main() {
	cfg := mysql.Config{
		User:                 "user name",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "db name",
		AllowNativePasswords: true,
	}
	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Printf("Connected.\n")
	StdtName, err := studentByLname("Lname")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Details found %+v\n", StdtName)

	StdtId, err := getByID(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Details found %+v\n", StdtId)

	addStd, err := addStudent(Student{
		Fname: "Fname",
		Lname: "Lname",
		Age:   00,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Index Is %d", addStd)
}
func studentByLname(name string) ([]Student, error) {
	var std []Student
	rows, err := db.Query("SELECT * FROM <Database Name> WHERE lname=?", name)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var stdnew Student
		if err := rows.Scan(&stdnew.ID, &stdnew.Fname, &stdnew.Lname, &stdnew.Age); err != nil {
			return nil, fmt.Errorf("StudentByName %q:%v", name, err)
		}
		std = append(std, stdnew)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("StudentByName %q:%v", name, err)
	}

	return std, nil
}

func getByID(id int) (Student, error) {
	var std Student

	row := db.QueryRow("SELECT * FROM <Database Name> WHERE id = ?", id)
	if err := row.Scan(&std.ID, &std.Fname, &std.Lname, &std.Age); err != nil {
		if err == sql.ErrNoRows {
			return std, fmt.Errorf("getByID %d: no such student", id)
		}
		return std, fmt.Errorf("getByID %d: %v", id, err)
	}
	return std, nil
}

func addStudent(stdData Student) (int64, error) {
	row, err := db.Exec("INSERT INTO <DataBase Name>(fname,lname,age) VALUES(?,?,?)", stdData.Fname, stdData.Lname, stdData.Age)
	if err != nil {
		return 0, fmt.Errorf("addstudent %v ", err)
	}
	resualt, err := row.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addstudent %v ", err)
	}

	return resualt, nil

}
