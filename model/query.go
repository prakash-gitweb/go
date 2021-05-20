package model

import{

}

func dbConnect(){
	// Open database connection
	db, err := sql.Open("mysql", "root:Prkayy_0651@/pfh")
	if err != nil {
		panic(err.Error())  
	}
	defer db.Close()
}

func updateAgent (account uint, commission float32, volume float32) bool{
	dbConnect() // Connection to database
	update, err := db.Query("UPDATE ib SET available=available+?, volume=volume+? WHERE account_no=?",commission,closed.volume/100, agent)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer update.Close()
}
func updateTx (agent uint, account uint, commission float32, volume float32, ticket uint, symbol string) bool{
	
}