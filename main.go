package main

import (
	"database/sql"
	"encoding/csv"
	// "fmt"
	"log"
	"net/http"
	// "os"
	"strconv"
	"text/template"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	_ "github.com/lib/pq"
	
)


func getMySql() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=root port=5432 dbname=training sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

var temp *template.Template
var db *sql.DB

func uploadCSVHandler(c *gin.Context) {


	file, _, err := c.Request.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fp := csv.NewReader(file)
	data, err := fp.ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var successfulRows []map[string]interface{}
	var failedRows []map[string]interface{}

	for i, line := range data {
		if (i == 0)){
			continue
		}
		db = getMySql()
		_, err := db.Exec(`insert into "trainingchart" ("firstname", "lastname", "dept", "cloud", "trainingattended", "trainingpath", "email") values($1, $2, $3, $4, $5, $6, $7)`, line[0], line[1], line[2],line[3], line[4],line[5], line[6])
		db.Close()
		if err != nil {
			failedRows = append(failedRows, map[string]interface{}{
				"firstname":   line[0],
				"lastname": line[1],
				"dept":   line[2],
				"cloud": line[3],
				"trainingattended":   line[4],
				"trainingpath": line[5],
				"email": line[6],
				"error":  err.Error(),
			})
			continue
		}
		successfulRows = append(successfulRows, map[string]interface{}{
			"firstname":   line[0],
			"lastname": line[1],
			"dept":   line[2],
			"cloud": line[3],
			"trainingattended":   line[4],
			"trainingpath": line[5],
			"email": line[6],
		})
	}

	response := map[string]interface{}{
		"successfulRows": successfulRows,
		"failedRows":     failedRows,
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.POST("/uploadcsv", uploadCSVHandler)
	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8000", handler)
}
