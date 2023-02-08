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
    GearitemID      int64           `json:"gearitemid"`
    Customname      sql.NullString  `json:"customname"`
    Category        sql.NullString  `json:"category"`
    Brandname       sql.NullString  `json:"brandname"`
    Productname     sql.NullString  `json:"productname"`
    WeightKG        sql.NullFloat64 `json:"weightkg"`
    Quantity        sql.NullInt32   `json:"quantity"`
    UserID          sql.NullInt32   `json:"userid"`
    Picture         sql.NullString  `json:"picture"`
    Cost            sql.NullString  `json:"cost"`
    TempfromCELC    sql.NullInt16   `json:"tempfromcelc"`
    TempuptoCELC    sql.NullInt16   `json:"tempuptocelc"`
    Datepurchased   sql.NullString  `json:"datepurchased"`
    Dateadded       sql.NullString  `json:"dateadded"`
    Notes           sql.NullString  `json:"notes"`
}

type Checklist struct {
    Checklistid     int64           `json:"checklistid"`
    Checklistname   sql.NullString  `json:"checklistname"`
    Atctivitytype   sql.NullString  `json:"atctivitytype"`
    Gearweight      sql.NullFloat64 `json:"gearweight"`
    Datestart       sql.NullString  `json:"datestart"`
    Dateend         sql.NullString  `json:"dateend"`
    Tempfrom        sql.NullFloat64 `json:"tempfrom"`
    Tempupto        sql.NullFloat64 `json:"tempupto"`
}

type Gearchecklistrelations struct {
    Relationid      int64 `json:"relationid"`
    Gearitemid      int64 `json:"gearitemid"`
    Checklistid     int64 `json:"checklistid"`
    Packagecheck    bool  `json:"packagecheck"`
}
/*
"checklistname": "",
"atctivitytype": "",
"gearweight": "",
"datestart": "",
"dateend": "",
"tempfrom": "",
"tempupto": ""

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
