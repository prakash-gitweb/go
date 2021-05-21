package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func dbConnect() *sql.DB {
	// Open database connection
	db, err := sql.Open("mysql", "root:Prkayy0651@/pfhtest")
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Account struct {
	Type  string `json:"type"`
	Agent uint   `json:"agent"`
}

type AgentData struct {
	Rebate    float32
	Agent     uint
}

func GetAccount(account *int) Account {
	db := dbConnect() // Connection to database
	var data Account
	err := db.QueryRow("SELECT type, agent FROM live_accounts where account_no = ?", account).Scan(&data.Type, &data.Agent)
	if err != nil {
		panic(err.Error())
	}
	return data
}

func UpdateAgent(agent *uint, commission *float32, volume *float32) bool {
	db := dbConnect() // Connection to database
	update, err := db.Query("UPDATE ib SET available=available+?, volume=volume+? WHERE account_no=?", commission, volume, agent)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer update.Close()
	return true
}
func CreateOrUpdateAgentTx(agent *uint, account *int, commission *float32, volume *float32) bool {
	db := dbConnect() // Connection to database
	rows, err := db.Query("SELECT COUNT(*) as count FROM ib_tx where account = ? AND agent = ?", account,agent)
	checkErr(err)
	defer rows.Close()
	if checkCount(rows) > 0 {
		update, err := db.Query("UPDATE ib_tx SET profit=profit+?, volume=volume+? WHERE account=? AND agent=?", commission, volume, account, agent)
		checkErr(err)
		defer update.Close()
	} else {
		insert, err := db.Query("INSERT INTO ib_tx SET agent=?, account=?, profit=?, volume=?", agent, account, commission, volume)
		checkErr(err)
		defer insert.Close()
	}
	return true
}
func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
	   err:= rows.Scan(&count)
	   checkErr(err)	
   }   
   return count
}
func GetAgentOfAgent(account *uint) AgentData {
	db := dbConnect() // Connection to database
	var data AgentData
	err := db.QueryRow("SELECT rebate, agent FROM ib where account_no = ?", account).Scan(&data.Rebate, &data.Agent)
	checkErr(err)
	return data
}

func CreateTradeTx(ticket *uint, account *uint, closeTime *string, symbol *string, volume *float32, commission *float32 ){
	db := dbConnect() // Connection to database
	insert, err := db.Query("INSERT INTO trade_tx SET ticket=?, account=?, closeTime=?, symbol=?, volume=?, profit=?",
							ticket, account, closeTime, symbol, volume, commission)
	checkErr(err)
	defer insert.Close()
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}