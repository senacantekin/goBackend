package main

import (
	"fmt"
	"net/http"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type PostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}
type PostResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
}

const (
	host     = "localhost"
	port     = 5439
	user     = "postgres"
	password = "depixen-pass"
	dbname   = "postgres"
)

func OpenConnection() *sql.DB {

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

func main() {

	//now := time.Now()
	///

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")

}
func getAlbums(c *gin.Context) {
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM tb_casestudy")
	if err != nil {
		fmt.Printf("there was an error in select query %v", err)
	}
	var response []PostResponse
	for rows.Next() {
		row := PostResponse{}
		err := rows.Scan(&row.ID, &row.Title, &row.Description, &row.ImageUrl, &row.CreatedAt)
		if err != nil {
			fmt.Printf("there was an error in scan operation %v", err)
			c.IndentedJSON(http.StatusInternalServerError, nil)

		}
		response = append(response, row)
	}
	c.IndentedJSON(http.StatusOK, response)

	defer db.Close()

}

func postAlbums(c *gin.Context) {
	db := OpenConnection()

	var PostRequest PostRequest

	if err := c.BindJSON(&PostRequest); err != nil {
		return
	}
	epoch := time.Now().Format("2006-01-02 15:04:05.000000")

	insertDynStmt := `insert into "tb_casestudy"("title", "description","image_url","created_at") values($1,$2,$3,$4)`
	_, err := db.Exec(insertDynStmt, PostRequest.Title, PostRequest.Description, PostRequest.ImageUrl, epoch)

	if err != nil {
		fmt.Printf("there was an error %v", err)

	}

	c.IndentedJSON(http.StatusCreated, PostRequest)
	defer db.Close()

}
