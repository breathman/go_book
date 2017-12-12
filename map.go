package main

import "fmt"

func main() {
	currencies:= make(map[string]string)

	currencies["btc"] = "Bitcoin"
	currencies["eth"] = "Etherium"
	currencies["bch"] = "Bitcoin cash"

	if name, ok := currencies["bch"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("Error")
	}

	currencies2 := map[string]string{
		"BTC": "Bitcoin",
		"LTC": "Litecoin",
		"ETH": "Etherium",
	}

	fmt.Println(currencies2["BTC"])

	details := map[string]map[string]string {
		"BTC": {
			"name": "Bitcoin",
			"price": "16000",
		},
		"ETH": {
			"name": "Etherium",
			"price": "469",
		},
	}

	if el, ok := details["BTC"]; ok {
		fmt.Println(el["name"], el["price"])
	}
}
