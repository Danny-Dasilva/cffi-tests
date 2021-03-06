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
	Headers         map[string]string `json:"headers"`
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
	var basket Options
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		log.Println(err)
	fmt.Println(basket.URL, basket.Ja3, basket.UserAgent)
	// return basket.Name
	jsonData, err = json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	return C.CString(string(jsonData))
}	
}


func main() {
	name := getVertex(`
	{
		"url": "https://google.com",
		"ja3": "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0",
		"userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36"
	}`)
	fmt.Println(name)
}
