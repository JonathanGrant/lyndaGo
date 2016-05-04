package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	//How do I open a file?
	f, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println(f)

	//How do I read the file?
	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		fmt.Println(row)
	}
}
