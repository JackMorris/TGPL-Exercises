package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.handleList)
	http.HandleFunc("/price", db.handlePrice)
	http.HandleFunc("/add", db.handleAdd)
	http.HandleFunc("/update", db.handleUpdate)
	http.HandleFunc("/delete", db.handleDelete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // Use global mux
}

// Dollars definition //

type dollars int

func (d dollars) String() string {
	return fmt.Sprintf("$%d", d)
}

// Database definition //

type database map[string]dollars

func (db database) handleList(w http.ResponseWriter, req *http.Request) {
	// List all of the items in the db, and their price, to w.
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) handlePrice(w http.ResponseWriter, req *http.Request) {
	// Print the price of the specified item to w if it exists, otherwise print an error.
	item := req.URL.Query().Get("item")

	// Check that the item exists.
	price, exists := db[item]
	if !exists {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "%s\n", price)
}

func (db database) handleAdd(w http.ResponseWriter, req *http.Request) {
	// Add the specified item to the db if it doesn't already exist, otherwise print an error.
	item := req.URL.Query().Get("item")

	// Check if the item already exists.
	_, exists := db[item]
	if exists {
		w.WriteHeader(http.StatusConflict) //409
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}

	// Ensure the new price is valid.
	price, err := strconv.Atoi(req.URL.Query().Get("price"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}

	// Add the new item.
	db[item] = dollars(price)
	fmt.Fprintf(w, "item added: %q\n", item)
}

func (db database) handleUpdate(w http.ResponseWriter, req *http.Request) {
	// Update the price of the specified item in the db if it exists, otherwise print an error.
	item := req.URL.Query().Get("item")

	// Check that the item exists.
	_, exists := db[item]
	if !exists {
		w.WriteHeader(http.StatusConflict) //409
		fmt.Fprintf(w, "item doesn't exist: %q\n", item)
		return
	}

	// Ensure the new price is valid.
	price, err := strconv.Atoi(req.URL.Query().Get("price"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}

	// Update the price.
	db[item] = dollars(price)
	fmt.Fprintf(w, "price of %s updated to %s\n", item, db[item])
}

func (db database) handleDelete(w http.ResponseWriter, req *http.Request) {
	// Delete the item from the db if it exists, otherwise print an error
	item := req.URL.Query().Get("item")

	// Check that the item exists.
	_, exists := db[item]
	if !exists {
		w.WriteHeader(http.StatusConflict) //409
		fmt.Fprintf(w, "item doesn't exist: %q\n", item)
		return
	}

	// Delete the item from the database.
	delete(db, item)
	fmt.Fprintf(w, "item deleted: %q\n", item)
}
