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
