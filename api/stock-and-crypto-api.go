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

	//Read from the Thread Table

	// Read from the Favorites Table

	// Create new User in User Table

	// Create new Thread in Thread Table

	// Create new Favorite in Favorites Table

}

func mwCheck(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func readUserTable() (int, error) {}

func readThreadTable() (int, error) {}

func readFavoritesTable() (int, error) {}

func createNewUser() (model.UserTable, error) {}

func createNewThread() (model.ThreadTable, error) {}

func createNewFavorite() (model.FavoritesTable, error) {}
