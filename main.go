package main

import (
	"database/sql"
	"os"
	"log"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prakash-gitweb/go/model"
)
type LiveAccount struct {
    Type string `json:"type"`
	Agent uint `json:"agent"`
}
type order struct {
	ticket uint
	volume float32
	symbol string
	account uint
	price float32
}
type agent struct{
	rebate float32
	available float32
	volume float32
	agent uint
}
type Test struct{
	Symbol string `json:"symbol"`
	Value float32 `json:"value"`
	
}
func main() {
	// Open database connection
	db, err := sql.Open("mysql", "root:Prkayy_0651@/pfh")
	if err != nil {
		panic(err.Error())  
	}
	defer db.Close()

	closed := order{
		ticket: 100000,
		volume: 100,
		symbol: "AUDHKD.ecn",
		account: 444499,
		price: 1.120098,
	} 

	fmt.Println(closed)
	var tag LiveAccount
	err = db.QueryRow("SELECT type, agent FROM live_accounts where account_no = ?", closed.account).Scan(&tag.Type, &tag.Agent)
	if err != nil {
    	panic(err.Error()) 
	}

	const layout = "02-01-2006"
	t:= time.Now()
	file := fmt.Sprintf("logs/%s.txt", t.Format(layout))
	f, err := os.OpenFile(file,
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	logger.Println(closed)
	// log the transactions in logs folder

	if tag.Agent > 0 {
		var totalRebate float32
		var y float32
		agent := tag.Agent
		for{
			agentData := getAgent(agent)
			x := agentData.rebate
			rebate := x-y
			pip := getPip(closed.symbol)
			commission := pip*closed.volume*rebate/10000
			

			y = x
			totalRebate += rebate

			fmt.Println(totalRebate)
			logger.Printf("agent: %d, rebate: %d, commission: %f,", agent, rebate, commission)
			
			if agentData.agent == 0{
				break
			}
			if totalRebate >= 100 {
				break
			}
			agent = agentData.agent
			
			
		}
	}
	logger.Println("-----------------------------------------------------------------")

	// insert, err := db.Query("UPDATE ib SET total_withdrawal=120 WHERE email='pk@pk.com'")
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// defer insert.Close()
}
func getAgent(account uint) agent{
	db, err := sql.Open("mysql", "root:Prkayy_0651@/pfh")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	var tag agent
	err = db.QueryRow("SELECT rebate, available, volume, agent FROM ib where account_no = ?", account).Scan(&tag.rebate, &tag.available, &tag.volume, &tag.agent)
	if err != nil {
    	panic(err.Error()) // proper error handling instead of panic in your app
	}
	return agent{
		rebate: tag.rebate,
		available: tag.available,
		volume: tag.volume,
		agent: tag.agent,
	}
}

func getPip(symbol string) float32{
	symbol = strings.Trim(symbol, ".ecn")
	url := "http://161.97.101.141/fxrates/pips.txt"
	
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)

	}
	body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        panic(err.Error())
    }
	
	defer resp.Body.Close()
	var data []Test
	json.Unmarshal(body, &data)

	var val float32
	for _, v := range data {
		// fmt.Printf("", i)
		if v.Symbol==symbol{
			val = v.Value
		}
	}
	return val

}
