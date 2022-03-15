package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Order struct {
	OrderId      string `json:"orderId"`
	CustomerName string `json:"customerName"`
	OrderAt      string `json:"orderAt"`
	Items        []Item `json:"items"`
}
type Item struct {
	ItemId      string `json:"itemId"`
	Description string `json:"description"`
	quantity    int    `json:"quantity"`
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var orders []Order
	json.NewDecoder(r.Body).Decode(&orders)
	orders = append(orders, Order)
	
	// orders = append(orders, Order{OrderId: "1", CustomerName: "John", OrderAt: "2018-01-01", Items: []Item{Item{ItemId: "1", Description: "Item 1", quantity: 1}, Item{ItemId: "2", Description: "Item 2", quantity: 2}}})
	// orders = append(orders, Order{OrderId: "2", CustomerName: "John1", OrderAt: "2018-01-02", Items: []Item{Item{ItemId: "1", Description: "Item 1", quantity: 1}, Item{ItemId: "2", Description: "Item 2", quantity: 2}}})
	// orders = append(orders, Order{OrderId: "3", CustomerName: "John2", OrderAt: "2018-01-03", Items: []Item{Item{ItemId: "1", Description: "Item 1", quantity: 1}, Item{ItemId: "2", Description: "Item 2", quantity: 2}}})

}

var orders []Order

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", getOrders).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}
