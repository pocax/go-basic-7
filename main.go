package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "db-go-sql"
)

var (
	db  *sql.DB
	err error
)

type Employee struct {
	ID       int
	Fullname string
	Email    string
	Age      int
	Division string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	//CreateEmployee()
	//GetEmployee()
	//UpdateEmployee()
	DeleteEmployee()
}

func CreateEmployee() {
	var employee = Employee{}
	sqlStatement := `
		INSERT INTO employee (fullname, email, age, division)
		VALUES ($1, $2, $3, $4)
		RETURNING *
	`
	err = db.QueryRow(sqlStatement, "John Doe", "john@doe.com", 23, "IT").
		Scan(&employee.ID, &employee.Fullname, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee ID is: %+v\n", employee)
}

func GetEmployee() {
	var results = []Employee{}

	sqlStatement := `SELECT * FROM employee`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}
		err = rows.Scan(&employee.ID, &employee.Fullname, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}
	fmt.Println("Employees:", results)
}

func UpdateEmployee() {
	sqlStatement := `UPDATE employee SET fullname = $2, email = $3, division = $4, age = $5 WHERE id = $1;`
	res, err := db.Exec(sqlStatement, 1, "John Linus", "john@linus.com", "IT", 25)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Update data amount:", count)
}

func DeleteEmployee() {
	sqlStatement := `DELETE FROM employee WHERE id = $1`

	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete data amount:", count)
}
