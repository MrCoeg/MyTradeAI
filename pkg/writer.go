package pkg

import (
	"encoding/csv"
	"os"
)

func WriteCSV(records [][]string) {
	f, err := os.Create("users.csv")
	LogIfError(err)
	defer f.Close()
	LogIfError(err)

	w := csv.NewWriter(f)
	w.WriteAll(records)
}
