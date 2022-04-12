package main

import (
	"encoding/json"
	"log"
	"net/http"
	model "stock-and-crypto-api/models"
	"github.com/gorilla/mux"
	"database/sql"
	"context"
	_"github.com/denisenkom/go-mssqldb"
	"fmt"
	"time"
)

var db *sql.DB
var server = ""
var user = ""
var password = ""
var database = ""

func main() {
	// Build connection string
	connectionString := fmt.Sprintf("server%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password. port, database)
	var err error

	// Create connection
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Error connecting:", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected to the database!\n")

	// Read from the User Table
	userCount, userErr := readUserTable()
	if userErr != nil {
		log.Fatal("Error reading user table:", userErr.Error())
	}
	fmt.Printf("Read %d row(s) successfully", userCount)

	//Read from the Thread Table
	threadCount, threadErr := readThreadTable()
	if threadErr != nil {
		log.Fatal("Error reading thread table:", threadErr.Error())
	}
	fmt.Printf("Read %d row(s) successfully", threadCount)

	// Read from the Favorites Table
	favoritesCount, favoritesErr := readFavoritesTable()
	if favoritesErr != nil {
		log.Fatal("Error reading favorites table:", favoritesErr.Error())
	}
	fmt.Printf("Read %d row(s) successfully", favoritesCount)

	// Create new User in User Table
	newUser, newUserErr := createNewUser()
	if newUserErr != nil {
		log.Fatal("Error creating new user:", newUserErr.Error())
	}
	fmt.Printf("Created new user successfully", newUser)

	// Create new Thread in Thread Table
	newThread, newThreadErr := createNewThread()
	if newThreadErr != nil {
		log.Fatal("Error creating new thread:", newThreadErr.Error())
	}
	fmt.Printf("Created new thread succcessfully", newThread)

	// Create new Favorite in Favorites Table
	newFavorite, newFavoriteErr := createNewFavorite()
	if newFavoriteErr != nil {
		log.Fatal("Error creating new favorite:", newFavoriteErr.Error())
	}
	fmt.Printf("Created new favorite successfully", newFavorite)
}

func mwCheck(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func readUserTable() (int, error) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsqlQuery := fmt.Sprintf("SELECT * FROM Users")

	// Execute query
	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int
	for rows.Next() {
		var userid int32
		var email string
		var username string
		var password string

		err := rows.Scan(&userid, &email, &username, &password)
		if err != nil {
			return -1, err
		}
		// do work here
		count++
	}
	return count, nil
}

func readThreadTable() (int, error) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsqlQuery := fmt.Sprintf("SELECT * FROM Thread")

	// Execute query
	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int
	for rows.Next() {
		var threadid int32
		var userid int32
		var name string
		var description string
		var datecreated time.Time

		err := rows.Scan(&threadid, &userid, &name, &description, &datecreated)
		if err != nil {
			return -1, err
		}
		// do work here
		count++
	}
	return count, nil
}

func readFavoritesTable() (int, error) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsqlQuery := fmt.Sprintf("SELECT * FROM Favorites")

	// Execute query
	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int
	for rows.Next() {
		var favoritesid int32
		var userid string
		var stockid string
		var apiurl string

		err := rows.Scan(&favoritesid, &userid, &stockid, &apiurl)
		if err != nil {
			return -1, err
		}
		// do work here
		count++
	}
	return count, nil
}

func createNewUser() (model.UserTable, error) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsqlMutation := fmt.Sprintf("INSERT INTO Users VALUES()") // finish mutation

	// Execute query
	newUser, err := db.QueryContext(ctx, tsqlMutation)
	if err != nil {
		return nil, err
	}

	defer newUser.Close()

	// finisn here
}

func createNewThread() (model.ThreadTable, error) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsqlMutation := fmt.Sprintf("INSERT INTO Thread VALUES()") // finish mutation

	// Execute query
	newThread, err := db.QueryContext(ctx, tsqlMutation)
	if err != nil {
		return nil, err
	}

	defer newThread.Close()

	// finisn here
}

func createNewFavorite() (model.FavoritesTable, error) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsqlMutation := fmt.Sprintf("INSERT INTO Favorites VALUES()") // finish mutation

	// Execute query
	newFavorite, err := db.QueryContext(ctx, tsqlMutation)
	if err != nil {
		return nil, err
	}

	defer newFavorite.Close()

	// finisn here
}
