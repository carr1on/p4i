package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Functions interface {
	ReadOnServiceSMS(fullContent []byte) (err error)
	ReportSMS(slice1 string) (err error)
}

var fu Functions

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

type DataService map[int]*SMSData

var ddb *DataService

func main() {

	smsdata := "tools.go.data"
	file, err := os.Open(filepath.Join("test/diploma/" + smsdata))
	if err != nil {
		fmt.Println("err")
	}
	defer file.Close()

	err = readerCSV(file)
	if err != nil {
		fmt.Println("err reader")
	}

	if err != nil {
		fmt.Print("err")
	}

}

func readerCSV(file io.Reader) error {
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // "-1 - считывание всего из файла"
	reader.Comma = ';'

	for {
		record, err := reader.Read()
		if err != nil {
			err = fmt.Errorf("err record")
			break
		}
		err = ddb.ReportSMS(record)
		if err != nil {
			fmt.Println("err")
		}

	}
	return nil
}

func (d *DataService) ReportSMS(record []string) (err error) {

	for i := 0; i < len(record); i++ {

		if len(record) == 4 {

			if record[3] == "Topolo" || record[3] == "Rond" || record[3] == "Kildy" {
				fmt.Print("Country ", record[0])
				fmt.Print(" Bandwidth ", record[1])
				fmt.Print(" ResponseTime ", record[2])
				fmt.Print(" Provider ", record[3], "\n")

				record = nil
			} else {
				fmt.Println("Provider currupt!")
				break
			}
		} else {
			record = nil
			return err
		}
	}
	return nil
}
