package middleware

import (
	"all-my-gear-backend-go/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllGearCheckRelations(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    relations, err := getAllGearCheckRelations()

    if err != nil {
        log.Fatalf("Unable to get all relations. %v", err)
    }

    json.NewEncoder(w).Encode(relations)
}

func GetRelationsByChecklistID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    params := mux.Vars(r)

    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    relations, err := getRelationsByChecklistID(int64(id))

    if err != nil {
        log.Fatalf("Unable to get relations. %v", err)
    }

    json.NewEncoder(w).Encode(relations)
}

//--------------


func getAllGearCheckRelations() ([]models.Gearchecklistrelations, error) {
    db := createConnection()

    defer db.Close()

    var relations []models.Gearchecklistrelations

    sqlStatement := `SELECT * FROM gearchecklistrelations`

    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    defer rows.Close()

    for rows.Next() {
        var relation models.Gearchecklistrelations

        err = rows.Scan(&relation.Relationid, &relation.Gearitemid, &relation.Checklistid, &relation.Packagecheck)
        
        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }
        relations = append(relations, relation)
    }

    return relations, err
}


func getRelationsByChecklistID(id int64) ([]models.Gearchecklistrelations, error) {
    db := createConnection()

    defer db.Close()
	
    var relations []models.Gearchecklistrelations

    sqlStatement := `SELECT * FROM gearchecklistrelations WHERE checklistid=$1`

    rows, err := db.Query(sqlStatement, id)

	if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    defer rows.Close()

    for rows.Next() {
        var relation models.Gearchecklistrelations

        err = rows.Scan(&relation.Relationid, &relation.Gearitemid, &relation.Checklistid, &relation.Packagecheck)
        
        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }
        relations = append(relations, relation)
    }

    return relations, err
}