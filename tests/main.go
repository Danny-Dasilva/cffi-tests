
package main


import (
	"log"
	"encoding/json"
	"fmt"
)


type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64 `json:"ref"`
}


func getVertex(data string) string {
	jsonData := []byte(data)
	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(basket.Name, basket.Fruit, basket.Id)
	return basket.Name
}

func main() {
	getVertex(`
	{
		"Name": "Standard",
		"Fruit": [
			"Apple",
			"Banana",
			"Orange"
		],
		"ref": 999
	}`)
}
