package main

import (
	"fmt"
	"strings"
)

type MemoryStore struct {
	Id         string
	Name       string
	Loc1       string
	Loc2       string
	Loc3       string
	Loc4       string
	Loc5       string
	Loc6       string
	Price      string
	Model      string
	FinalPrice string
}

var memoryStores []MemoryStore

func LoadMemoryStore(filename string) {
	readParse(filename)
}

func SearchViaName(name string) []MemoryStore {
	var res []MemoryStore
	for _, store := range memoryStores {
		if strings.Contains(strings.ToLower(store.Name), strings.ToLower(name)) {
			res = append(res, store)
		}
	}
	return res
}

func SearchViaSN(id string) (MemoryStore, error) {
	for _, store := range memoryStores {
		if store.Id == id {
			return store, nil
		}
	}
	return MemoryStore{}, fmt.Errorf("store with ID %s not found", id)
}

func Pagination(page int) []MemoryStore {
	var res []MemoryStore
	if page == 0 {
		return memoryStores
	}
	if page < 0 {
		return res
	}
	if page > len(memoryStores) {
		return res
	}
	start := (page - 1) * 30
	end := start + 30
	if end > len(memoryStores) {
		end = len(memoryStores)
	}
	return memoryStores[start:end]
}
