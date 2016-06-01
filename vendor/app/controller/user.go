package controller

import (

	"net/http"
	"app/shared/database"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	result := database.Cassandra.Query(`SELECT * FROM user`)
	result.Iter()
}
