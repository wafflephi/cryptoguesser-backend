package configs

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func exportTransactionsToFile(fileName string, transactions []Transaction) error {

	file, err := os.Create("./archive/" + fileName + ".csv")
	if err != nil {
		return err
	}

	var fileData [][]string
	label := []string{"Name", "Price", "Hour", "Action"}
	fileData = append(fileData, label)

	for i := range transactions {
		transaction := transactions[i]

		//? Is there a better way to this ?
		var line []string
		line = append(line, transaction.Name)
		line = append(line, fmt.Sprint(transaction.Price))
		line = append(line, transaction.Hour)
		line = append(line, fmt.Sprint(transaction.Action))

		fileData = append(fileData, line)
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll(fileData)
	if err != nil {
		return err
	}

	return nil
}

func SetupArchive() {
	_, err := os.Stat("./archive")
	if os.IsNotExist(err) {
		log.Println("Dir does not exist, creating one.")
		err := os.Mkdir("./archive", 0755)
		if err != nil {
			log.Panic("Could not create archive")
		}
	}
}

func ArchiveToday() error {
	//* Create the archive from Redis transactions
	transactions, err := GetAllTransactions()
	if err != nil {
		return err
	}
	if transactions == nil {
		log.Println("No entries to the archive")
		return nil
	}

	date := Today.Format("02-01-2006")
	log.Println("Date:", date)

	err = exportTransactionsToFile(date, transactions)
	if err != nil {
		return err
	}

	err = ClearAllTransactions()
	if err != nil {
		return err
	}

	return nil
}
