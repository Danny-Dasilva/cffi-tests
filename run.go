package main
/*
#include <stdlib.h>
typedef struct {
	int X;
	int Y;
} Nested;

*/
import "C"
// import "unsafe"
import (
	"log"
	"encoding/json"
	"fmt"
	"unsafe"
)


type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64 `json:"ref"`
}

type Response struct {
	RequestID string
	Status    int
	Body      string
	Headers   map[string]string
}




//export Add
func Add(a, b int) int {
	return a + b
}


//export Sort
func Sort(vals []int) {
	fmt.Println(vals)

}

//export Log
func Log(msg string) {
	fmt.Println(msg)
}

type Options struct {
	URL             string            `json:"url"`
	Method          string            `json:"method"`
	Headers         []C.Nested 
	Body            string            `json:"body"`
	Ja3             string            `json:"ja3"`
	UserAgent       string            `json:"userAgent"`
	Proxy           string            `json:"proxy"`
	// Cookies         []Cookie          `json:"cookies"`
	Timeout         int               `json:"timeout"`
	DisableRedirect bool              `json:"disableRedirect"`
	HeaderOrder     []string          `json:"headerOrder"`
	OrderAsProvided bool              `json:"orderAsProvided"` //TODO
}
// func Do(URL string, Method string, Headers []C.Nested, Body string, Ja3 string, UserAgent string, Proxy string, Timeout int, DisableRedirect bool, HeaderOrder []string, OrderAsProvided bool){

//export Do
func Do(URL string, Method string){
	fmt.Println(URL)
	fmt.Println(Method)
	// fmt.Println(Headers)
	// fmt.Println(Body)
	// fmt.Println(Ja3)
	// fmt.Println(UserAgent)
	// fmt.Println(Proxy)
	// fmt.Println(Timeout)
	// fmt.Println(DisableRedirect)
	// fmt.Println(HeaderOrder)
	// fmt.Println(OrderAsProvided)

}
//export freeString
func freeString(cs *C.char) {
	C.free(unsafe.Pointer(cs))
}

//export getVertex
func getVertex(data string) *C.char {
	jsonData := []byte(data)
	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(basket.Name, basket.Fruit, basket.Id)
	// return basket.Name
	jsonData, err = json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	return C.CString(string(jsonData))
}

func main() {
	name := getVertex(`
	{
		"Name": "Standard",
		"Fruit": [
			"Apple",
			"Banana",
			"Orange"
		],
		"ref": 999
	}`)
	fmt.Println(name)
}
