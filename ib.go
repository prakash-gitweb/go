package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Pip struct {
	Symbol string  `json:"symbol"`
	Value  float32 `json:"value"`
}

func IbDeposit(ticket *int, account *int, symbol *string, vol *int) bool {
	var volume float32 = float32(*vol) / 100
	// Logging transaction
	const layout = "02-01-2006"
	t := time.Now()
	file := fmt.Sprintf("logs/%s.txt", t.Format(layout))
	f, err := os.OpenFile(file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	logger.Println("ticket:", *ticket, "account:", *account, "symbol:", *symbol, "volume:", volume)
	
	a := GetAccount(account) // Rebate distribution starts
	if a.Agent > 0 {
		var totalRebate float32
		var y float32
		agent := a.Agent
		for {
			agentData := GetAgentOfAgent(&agent)
			x := agentData.Rebate
			rebate := x - y
			pip := getPip(*symbol)
			commission := pip * volume * rebate / 100
			UpdateAgent(&agent, &commission, &volume)
			CreateOrUpdateAgentTx(&agent, account, &commission, &volume)
			y = x
			totalRebate += rebate
			fmt.Println(totalRebate)
			logger.Printf("agent: %d, rebate: %f, commission: %f,", agent, rebate, commission)

			if agentData.Agent == 0 || totalRebate >= 100 {
				break
			}
			agent = agentData.Agent

		}
	}
	logger.Println("-----------------------------------------------------------------")

	return true
}


func getPip(symbol string) float32 {
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
	var data []Pip
	json.Unmarshal(body, &data)

	var val float32
	for _, v := range data {
		// fmt.Printf("", i)
		if v.Symbol == symbol {
			val = v.Value
		}
	}
	return val

}
