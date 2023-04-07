package main

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

const (
	host     = "8080"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func OpenConnection() *sql.DB {
	fmt.Println("successsssssssssssssssssssssssssssssddsssssssssssssssssssssssssssssssssssss")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db

}

var albums = []album{
	{ID: "1", Title: "blue train", Artist: "john", Price: 434.2},
	{ID: "2", Title: "blue train", Artist: "john", Price: 434.2},
	{ID: "3", Title: "blue train", Artist: "john", Price: 434.2},
}

func main() {
	//now := time.Now()
	///fmt.Println(now.Format("2006-02-01"))

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")

}
func getAlbums(c *gin.Context) {
	db := OpenConnection()

	c.IndentedJSON(http.StatusOK, albums)

	defer db.Close()

}
func postAlbums(c *gin.Context) {
	db := OpenConnection()

	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	fmt.Println(newAlbum)

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
	defer db.Close()

}
