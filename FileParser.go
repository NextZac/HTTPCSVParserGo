package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readParse(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file:" + err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		memoryStore := MemoryStore{
			Id:         strings.Trim(parts[0], `"`),
			Name:       strings.Trim(parts[1], `"`),
			Loc1:       strings.Trim(parts[2], `"`),
			Loc2:       strings.Trim(parts[3], `"`),
			Loc3:       strings.Trim(parts[4], `"`),
			Loc4:       strings.Trim(parts[5], `"`),
			Loc5:       strings.Trim(parts[6], `"`),
			Loc6:       strings.Trim(parts[7], `"`),
			Price:      strings.Trim(parts[8], `"`),
			Model:      strings.Trim(parts[9], `"`),
			FinalPrice: strings.Trim(parts[10], `"`),
		}

		memoryStores = append(memoryStores, memoryStore)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error readin file:" + err.Error())
		return
	}

}
