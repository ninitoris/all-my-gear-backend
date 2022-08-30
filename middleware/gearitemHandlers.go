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

func CreateGearitem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    var gearitem models.Gearitem

    err := json.NewDecoder(r.Body).Decode(&gearitem)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    insertID := insertGearitem(gearitem)

    res := response{
        ID:      insertID,
        Message: "Gearitem created successfully",
    }

    json.NewEncoder(w).Encode(res)
}

func GetGearItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    params := mux.Vars(r)

    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    gearitem, err := getGearitem(int64(id))

    if err != nil {
        log.Fatalf("Unable to get gearitem. %v", err)
    }

    json.NewEncoder(w).Encode(gearitem)
}

func GetAllGearitems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    gearitems, err := getAllGearitems()

    if err != nil {
        log.Fatalf("Unable to get all gearitems. %v", err)
    }

    json.NewEncoder(w).Encode(gearitems)
}

func UpdateGearitems(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    params := mux.Vars(r)

    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    var gearitem models.Gearitem

    err = json.NewDecoder(r.Body).Decode(&gearitem)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    updatedRows := updateGearitem(int64(id), gearitem)

    msg := fmt.Sprintf("Gearitem updated successfully. Total rows/record affected %v", updatedRows)

    res := response{
        ID:      int64(id),
        Message: msg,
    }

    json.NewEncoder(w).Encode(res)
}

func DeleteGearitem(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    params := mux.Vars(r)

    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    deletedRows := deleteGearitem(int64(id))

    msg := fmt.Sprintf("Gearitem deleted successfully. Total rows/record affected %v", deletedRows)

    res := response{
        ID:      int64(id),
        Message: msg,
    }

    json.NewEncoder(w).Encode(res)
}

//------------------------- handler functions ----------------
func insertGearitem(gearitem models.Gearitem) int64 {

    db := createConnection()

    defer db.Close()

    sqlStatement := `INSERT INTO gearitem (Customname, Category, Brandname, Productname, WeightKG, Quantity, UserID, Picture, Cost, TempfromCELC, TempuptoCELC, Datepurchased, Dateadded, Notes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING gearitemid`

    var id int64

    err := db.QueryRow(sqlStatement, gearitem.Customname, gearitem.Category, gearitem.Brandname, gearitem.Productname, gearitem.WeightKG, gearitem.Quantity, gearitem.UserID, gearitem.Picture, gearitem.Cost, gearitem.TempfromCELC, gearitem.TempuptoCELC, gearitem.Datepurchased, gearitem.Dateadded, gearitem.Notes ).Scan(&id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", id)

    return id
}

func getGearitem(id int64) (models.Gearitem, error) {
    db := createConnection()

    defer db.Close()

    var gearitem models.Gearitem

    sqlStatement := `SELECT * FROM gearitem WHERE gearitemid=$1`

    row := db.QueryRow(sqlStatement, id)

    err := row.Scan(&gearitem.GearitemID, &gearitem.Customname, &gearitem.Category, &gearitem.Brandname, &gearitem.Productname, &gearitem.WeightKG, &gearitem.Quantity, &gearitem.UserID, &gearitem.Picture, &gearitem.Cost, &gearitem.TempfromCELC, &gearitem.TempuptoCELC, &gearitem.Datepurchased, &gearitem.Dateadded, &gearitem.Notes)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return gearitem, nil
    case nil:
        return gearitem, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    return gearitem, err
}

func getAllGearitems() ([]models.Gearitem, error) {
    db := createConnection()

    defer db.Close()

    var gearitems []models.Gearitem

    sqlStatement := `SELECT * FROM gearitem`

    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    defer rows.Close()

    for rows.Next() {
        var gearitem models.Gearitem

        err = rows.Scan(&gearitem.GearitemID, &gearitem.Customname, &gearitem.Category, &gearitem.Brandname, &gearitem.Productname, &gearitem.WeightKG, &gearitem.Quantity, &gearitem.UserID, &gearitem.Picture, &gearitem.Cost, &gearitem.TempfromCELC, &gearitem.TempuptoCELC, &gearitem.Datepurchased, &gearitem.Dateadded, &gearitem.Notes)
        
        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }
        gearitems = append(gearitems, gearitem)
    }

    return gearitems, err
}

func updateGearitem(id int64, gearitem models.Gearitem) int64 {
    db := createConnection()

    defer db.Close()

    sqlStatement := `UPDATE gearitem SET customname=$2, category=$3, brandname=$4, productname=$5, weightkg=$6, quantity=$7, userid=$8, picture=$9, cost=$10, tempfromcelc=$11, tempuptocelc=$12, datepurchased=$13, dateadded=$14, notes=$15 WHERE gearitemid=$1`

    res, err := db.Exec(sqlStatement, id, gearitem.Customname, gearitem.Category, gearitem.Brandname, gearitem.Productname, gearitem.WeightKG, gearitem.Quantity, gearitem.UserID, gearitem.Picture, gearitem.Cost, gearitem.TempfromCELC, gearitem.TempuptoCELC, gearitem.Datepurchased, gearitem.Dateadded, gearitem.Notes)

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

func deleteGearitem(id int64) int64 {
    db := createConnection()

    defer db.Close()

    sqlStatement := `DELETE FROM gearitem WHERE gearitemid=$1`

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