package models

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

type NullString string

func (s *NullString) Scan(value interface{}) error {
    if value == nil {
        *s = ""
        return nil
    }
    strVal, ok := value.(string)
    if !ok {
        return errors.New("column is not a string")
    }
    *s = NullString(strVal)
    return nil
}

func (s NullString) Value() (driver.Value, error) {
    if len(s) == 0 { // if nil or empty string
        return nil, nil
    }
    return string(s), nil
}

// User schema of the user table
type User struct {
    ID       int64  `json:"id"`
    Name     string `json:"name"`
    Location string `json:"location"`
    Age      int64  `json:"age"`
}

// Gearitem schema
type Gearitem struct {
    GearitemID      int64           `json:"GearitemID"`
    Customname      NullString      `json:"Customname"`
    Category        NullString      `json:"Category"`
    Brandname       NullString      `json:"Brandname"`
    Productname     NullString      `json:"Productname"`
    WeightKG        sql.NullFloat64 `json:"WeightKG"`
    Quantity        sql.NullInt32   `json:"Quantity"`
    UserID          sql.NullInt32   `json:"UserID"`
    Picture         NullString      `json:"Picture"`
    Cost            NullString      `json:"Cost"`
    TempfromCELC    sql.NullInt16   `json:"TempfromCELC"`
    TempuptoCELC    sql.NullInt16   `json:"TempuptoCELC"`
    Datepurchased   NullString      `json:"Datepurchased"`
    Dateadded       NullString      `json:"Dateadded"`
    Notes           NullString      `json:"Notes"`
}


/*

    "GearitemID": "",
    "Customname": "",
    "Category": "",
    "Brandname": "",
    "Productname": "",
    "WeightKG": "",
    "Quantity": "",
    "UserID": "",
    "Picture": "",
    "Cost": "",
    "TempfromCELC": "",
    "TempuptoCELC": "",
    "Datepurchased": "",
    "Dateadded": "",
    "Notes": ""

*/
