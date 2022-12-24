package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	items []Item
)

func init() {
	items = []Item{{Id: "1", Title: "Item1", Text: "Description"}}
}

func getItems(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(items)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{error: "Error marshalling the posts array"}`))
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func addItem(resp http.ResponseWriter, req *http.Request) {
	var item Item
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{error: "Error unmarshalling the request"}`))
		return
	}
	item.Id = strconv.Itoa(len(items) + 1)
	items = append(items, item)
	result, _ := json.Marshal(item)
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func updateItem(resp http.ResponseWriter, req *http.Request) {
	var item Item
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{error: "Error unmarshalling the request"}`))
		return
	}
	for k, v := range items {
		if v.Id == item.Id {
			items[k] = item
		}
	}
	result, _ := json.Marshal(items)
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}
func removeItem(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	for index, item := range items {
		if item.Id == id {
			items = append(items[:index], items[index+1:]...)
		}
	}
	result, _ := json.Marshal(items)
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}
