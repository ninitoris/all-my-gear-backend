package middleware

import (
	"all-my-gear-backend-go/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllChecklists(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    checklists, err := getAllChecklists()

    if err != nil {
        log.Fatalf("Unable to get all checklists. %v", err)
    }

    json.NewEncoder(w).Encode(checklists)
}

func GetChecklist(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    params := mux.Vars(r)

    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    checklist, err := getChecklist(int64(id))

    if err != nil {
        log.Fatalf("Unable to get checklist. %v", err)
    }

    fmt.Println("sosiiiii suuukkaa", checklist.Checklistid)

    json.NewEncoder(w).Encode(checklist)
}

func CreateChecklist(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    var checklist models.Checklist

    err := json.NewDecoder(r.Body).Decode(&checklist)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    insertID := insertChecklist(checklist)

    res := response{
        ID:      insertID,
        Message: "checklist created successfully",
    }

    json.NewEncoder(w).Encode(res)
}


func UpdateChecklist(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    params := mux.Vars(r)

    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    var checklist models.Checklist

    err = json.NewDecoder(r.Body).Decode(&checklist)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    updatedRows := updateChecklist(int64(id), checklist)

    msg := fmt.Sprintf("checklist updated successfully. Total rows/record affected %v", updatedRows)

    res := response{
        ID:      int64(id),
        Message: msg,
    }

    json.NewEncoder(w).Encode(res)
}

func DeleteChecklist(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    params := mux.Vars(r)

    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    deletedRows := deleteChecklist(int64(id))

    msg := fmt.Sprintf("Checklist deleted successfully. Total rows/record affected %v", deletedRows)

    res := response{
        ID:      int64(id),
        Message: msg,
    }

    json.NewEncoder(w).Encode(res)
}

//------------------------- handler functions ----------------
func getAllChecklists() ([]models.Checklist, error) {
    db := createConnection()

    defer db.Close()

    var checklists []models.Checklist

    sqlStatement := `SELECT * FROM checklist`

    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    defer rows.Close()

    for rows.Next() {
        var checklist models.Checklist

        err = rows.Scan(&checklist.Checklistid, &checklist.Checklistname, &checklist.Atctivitytype, &checklist.Gearweight, &checklist.Datestart, &checklist.Dateend, &checklist.Tempfrom, &checklist.Tempupto)
        
        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }
        checklists = append(checklists, checklist)
    }

    return checklists, err
}


func getChecklist(id int64) (models.Checklist, error) {
    db := createConnection()

    defer db.Close()

    var checklist models.Checklist

    sqlStatement := `SELECT * FROM checklist WHERE checklistid=$1`

    row := db.QueryRow(sqlStatement, id)

    err := row.Scan(&checklist.Checklistid, &checklist.Checklistname, &checklist.Atctivitytype, &checklist.Gearweight, &checklist.Datestart, &checklist.Dateend, &checklist.Tempfrom, &checklist.Tempupto)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return checklist, nil
    case nil:
        return checklist, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    return checklist, err
}

func getChecklistWithItems(id int64) (models.Checklist, error) {
    db := createConnection()

    defer db.Close()

    var checklist models.Checklist

    sqlStatement := `SELECT * FROM checklist WHERE checklistid=$1`

    row := db.QueryRow(sqlStatement, id)

    err := row.Scan(&checklist.Checklistid, &checklist.Checklistname, &checklist.Atctivitytype, &checklist.Gearweight, &checklist.Datestart, &checklist.Dateend, &checklist.Tempfrom, &checklist.Tempupto)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return checklist, nil
    case nil:
        break;
        // return checklist, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }
    
    fmt.Println("sosiiiii suuukkaa", checklist.Checklistid)

    return checklist, err
}

func insertChecklist(checklist models.Checklist) int64 {

    db := createConnection()

    defer db.Close()

    sqlStatement := `INSERT INTO checklist (checklistname, atctivitytype, gearweight, datestart, dateend, tempfrom, tempupto) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING checklistid`

    var id int64

    err := db.QueryRow(sqlStatement, checklist.Checklistname, checklist.Atctivitytype, checklist.Gearweight, checklist.Datestart, checklist.Dateend, checklist.Tempfrom, checklist.Tempupto).Scan(&id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", id)

    return id
}

func updateChecklist(id int64, checklist models.Checklist) int64 {
    db := createConnection()

    defer db.Close()

    sqlStatement := `UPDATE checklist SET checklistname=$2, atctivitytype=$3, gearweight=$4, datestart=$5, dateend=$6, tempfrom=$7, tempupto=$8 WHERE checklistid=$1`

    res, err := db.Exec(sqlStatement, id, checklist.Checklistname, checklist.Atctivitytype, checklist.Gearweight, checklist.Datestart, checklist.Dateend, checklist.Tempfrom, checklist.Tempupto)

    if err != nil {
        log.Fatalf("Unable to execute the UPDATE query. %v", err)
    }

    rowsAffected, err := res.RowsAffected()

    if err != nil {
        log.Fatalf("Error while checking the affected rows. %v", err)
    }

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}

func deleteChecklist(id int64) int64 {
    db := createConnection()

    defer db.Close()

    sqlStatement := `DELETE FROM checklist WHERE checklistid=$1`

    res, err := db.Exec(sqlStatement, id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    rowsAffected, err := res.RowsAffected()

    if err != nil {
        log.Fatalf("Error while checking the affected rows. %v", err)
    }

    fmt.Printf("Total rows/record affected on DELETING %v", rowsAffected)

    return rowsAffected
}