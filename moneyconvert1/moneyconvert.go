package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

type MoneyConvertResult struct {
	Amount    float64
	CurrencyA string
	CurrencyB string
}

type CurrencyRate struct {
	Country     string
	CountryCode string
	Currency    string
	Code        string
	rate        float64
}

type conversionRate struct {
	key   string
	value float64
}

func readExchangeRateJson() []conversionRate {
	var result map[string]interface{}
	file, _ := ioutil.ReadFile("./file/exchangerate-api.json")
	_ = json.Unmarshal([]byte(file), &result)

	var converse map[string]interface{} = result["conversion_rates"].(map[string]interface{})

	conversion := []conversionRate{}
	keys := reflect.ValueOf(converse).MapKeys()
	for i := 0; i < len(converse); i++ {
		conversion = append(conversion, conversionRate{keys[i].String(), converse[keys[i].String()].(float64)})
	}
	return conversion
}

func getExchangeRate(currencyCode string) float64 {
	ExchangeRate := readExchangeRateJson()
	for i := 0; i < len(ExchangeRate); i++ {
		if strings.ToLower(ExchangeRate[i].key) == strings.ToLower(currencyCode) {
			return ExchangeRate[i].value
		}
	}
	return 0
}

func readCSVCurrencyRate() []CurrencyRate{
	filename := "./file/currency.csv"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to read input file "+filename, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filename, err)
	}
	//fmt.Println(len(records[0]))
	listCurrencyRate := []CurrencyRate{}
	for i := 1; i < len(records); i++ {
		check :=true
		currencyRate := CurrencyRate{records[i][0],
			records[i][1],
			records[i][2],
			records[i][3],
			getExchangeRate(records[i][3])}
		for j:=0;j < len(listCurrencyRate);j++ {
			if strings.ToLower(listCurrencyRate[j].Code) == strings.ToLower(currencyRate.Code){
				check = false
				break
			}
		}
		if check == true {
			listCurrencyRate = append(listCurrencyRate,currencyRate)
		}
	}
	fmt.Println(listCurrencyRate)
	return listCurrencyRate
}

func getResult(amount float64,currencyB string,currencyA string) float64 {
	return amount * getExchangeRate(currencyA) /getExchangeRate(currencyB)
}
