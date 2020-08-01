package main

import(
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type Record struct {
	ID          uint64  `csv:"id"`
	Name        string  `csv:"name"`
	Description string  `csv:"description"`
	Cost        float64 `csv:"cost"`
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nameToCost := map[string]float64{}

	csvReader := csv.NewReader(file)
	csvReader.LazyQuotes = true
	_, err = csvReader.Read()

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		id, _ := strconv.ParseUint(row[0], 10, 64)
		cost, _ := strconv.ParseFloat(row[3], 64)
		record := &Record{
			ID: id,
			Name: row[1],
			Description: row[2],
			Cost: cost,
		}
		if _, ok := nameToCost[record.Name]; ok {
			nameToCost[record.Name] += record.Cost
		} else {
			nameToCost[record.Name] = record.Cost
		}
	}

	keys := make([]string, 0, len(nameToCost))
	for key := range nameToCost {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s\t%.2f\n", key, nameToCost[key])
	}
	
}
