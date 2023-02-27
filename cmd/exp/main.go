package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		DBName:   "lenslocked",
		SSLMode:  "disable",
	}

	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected")

	// Create tables...
	_, err = db.Exec(`
		create table if not exists users (
			id serial primary key,
			name text,
			email text unique not null
		);
	
		create table if not exists orders (
			id serial primary key,
			user_id int not null,
			amount int,
			description text
		);
	`)
	if err != nil {
		panic(err)
	}

	// Create a user...
	/*
		row := db.QueryRow(
			`insert into users (name, email) values ($1, $2) returning id;`,
			"Bob",
			"bob3@example.com",
		)

		var id int
		err = row.Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println("New user ID:", id)
	*/

	// Get a user...
	/*
		row := db.QueryRow(`select name, email from users where id = $1`, 1)

		var name, email string
		err = row.Scan(
			&name,
			&email,
		)
		if err == sql.ErrNoRows {
			fmt.Println("error: no rows")
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("User is:", name, email)
	*/

	// Create some junk data for the orders table...
	/*
		userID := 1
		for i := 1; i <= 5; i++ {
			amount := i * 100
			desc := fmt.Sprintf("Fake order #%d", i)
			_, err := db.Exec(`
			  INSERT INTO orders(user_id, amount, description)
			  VALUES($1, $2, $3)`, userID, amount, desc)
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("Created fake orders.")
	*/

	type Order struct {
		ID          int
		UserID      int
		Amount      int
		Description string
	}

	var orders []Order
	userID := 1

	// Query multiple records...
	rows, err := db.Query(`select id, amount, description from orders where user_id = $1`, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		order.UserID = userID

		err := rows.Scan(
			&order.ID,
			&order.Amount,
			&order.Description,
		)
		if err != nil {
			panic(err)
		}

		orders = append(orders, order)
	}

	if rows.Err() != nil {
		panic(rows.Err())
	}

	fmt.Println(orders)
}
