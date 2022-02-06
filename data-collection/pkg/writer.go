package pkg

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func WriteCSV(records [][]string) {
	f, err := os.Create("users.csv")
	LogIfError(err)
	defer f.Close()
	LogIfError(err)

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		} else {
			fmt.Println("Success")
		}
	}
}
