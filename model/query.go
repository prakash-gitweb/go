package model

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	
)

func dbConnect() *sql.DB{
	// Open database connection
	db, err := sql.Open("mysql", "root:Prkayy_0651@/pfh")
	if err != nil {
		panic(err.Error())  
	}
	return db
}
type Account struct{
	Type string `json:"type"`
	Agent uint `json:"agent"`
}
var data Account
func GetAccount(account *uint) Account {
	db := dbConnect() // Connection to database
	err := db.QueryRow("SELECT type, agent FROM live_accounts where account_no = ?",account).Scan(&data.Type, &data.Agent)
	if err != nil {
    	panic(err.Error()) 
	}
	return data
}

func UpdateAgent (agent *uint, commission *float32, volume *float32) bool{
	db := dbConnect() // Connection to database
	update, err := db.Query("UPDATE ib SET available=available+?, volume=volume+? WHERE account_no=?",commission,volume, agent)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer update.Close()
	return true
}
func updateTx (agent uint, account uint, commission float32, volume float32, ticket uint, symbol string) bool{
return true
}