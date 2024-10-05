package todo

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)


type CsvStore struct {
	StoreName string
}



func (store CsvStore) ensureStoreExists() error {
	if checkFileExists(store.StoreName) {
		return nil
	}

	file, err := os.OpenFile(store.StoreName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if CheckError(err, "While opening a new file") {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)

	err = w.Write([]string{"ID", "Description", "CreatedAt", "IsCompleted"})
	if CheckError(err, "While writing to a new file") {
		return err
	}
	w.Flush()
	err = w.Error()
	if CheckError(err, "General writing error") {
		return err
	}

	return nil
}

func (store CsvStore) ReadItems(showAll bool) (items []DataItem) {
	items = []DataItem{}

	err := store.ensureStoreExists()
	if CheckError(err, "While ensuring store exists") {
		return
	}

	file, err := os.Open(store.StoreName)
	if CheckError(err, "While opening file to read") {
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if CheckError(err, "While reading file") {
		return
	}

	for i := 1; i < len(records); i++ {
		record := records[i]

		id, err := strconv.Atoi(record[0])
		if CheckError(err, "While converting ID to int") {
			return
		}

		isComplete := record[3] == "true"
		if !showAll && isComplete {
			continue
		}

		items = append(items, DataItem{ID: id, Description: &record[1], CreatedAt: &record[2], IsComplete: &isComplete})
	}

	return items
}

func (store CsvStore) ReadItem(id int) (item DataItem) {
	records := store.ReadItems(true)

	for _, record := range records {
		if record.ID == id {
			item = record
			return
		}
	}

	return
}

func (store CsvStore) CreateItem(itemDescription string) (item DataItem) {
	err := store.ensureStoreExists()
	if CheckError(err, "While ensuring store exists") {
		return
	}

	file, err := os.OpenFile(store.StoreName, os.O_RDWR, os.ModePerm)
	if CheckError(err, "While opening file to read") {
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if CheckError(err, "While reading file") {
		return
	}

	lastRecord := records[len(records)-1]
	lastRecordId, err := strconv.Atoi(lastRecord[0])
	if CheckError(err, "While converting ID to int") {
		lastRecordId = 0
	}

	isComplete := false
	createdAt := time.Now().Format(time.RFC3339)
	item = DataItem{Description: &itemDescription, CreatedAt: &createdAt, IsComplete: &isComplete}
	item.ID = lastRecordId + 1

	records = append(records, []string{strconv.Itoa(item.ID), *item.Description, *item.CreatedAt, strconv.FormatBool(*item.IsComplete)})

	file.Truncate(0)
	file.Seek(0, 0)
	csvWriter := csv.NewWriter(file)
	err = csvWriter.WriteAll(records)
	if CheckError(err, "While writing to new file") {
		return
	}

	csvWriter.Flush()
	err = csvWriter.Error()
	if CheckError(err, "General writing error") {
		return
	}

	return
}

func (store CsvStore) UpdateItem(id int, item DataItem) (updatedItem DataItem) {
	if !checkFileExists(store.StoreName) {
		return
	}

	file, err := os.OpenFile(store.StoreName, os.O_RDWR, os.ModePerm)
	if CheckError(err, "While opening file to read") {
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if CheckError(err, "While reading file") {
		return
	}

	idStr := strconv.Itoa(id)
	for _, record := range records {
		if record[0] == idStr {
			if item.Description != nil {
				record[1] = *item.Description
			}
			if item.CreatedAt != nil {
				record[2] = *item.CreatedAt
			}
			if item.IsComplete != nil {
				record[3] = strconv.FormatBool(*item.IsComplete)
			}
			// updatedItem = DataItem{Description: &record[1], IsComplete: }

			break
		}
	}

	file.Truncate(0)
	file.Seek(0, 0)
	csvWriter := csv.NewWriter(file)
	err = csvWriter.WriteAll(records)
	if CheckError(err, "While writing to new file") {
		return
	}

	csvWriter.Flush()

	err = csvWriter.Error()
	if CheckError(err, "General writing error") {
		return
	}

	return DataItem{}
}

func (store CsvStore) DeleteItem(id int) {
	if !checkFileExists(store.StoreName) {
		return
	}

	file, err := os.OpenFile(store.StoreName, os.O_RDWR, os.ModePerm)
	if CheckError(err, "While opening file to read") {
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if CheckError(err, "While reading file") {
		return
	}

	filteredRecords := [][]string{}
	idStr := strconv.Itoa(id)
	for _, record := range records {
		if record[0] != idStr {
			filteredRecords = append(filteredRecords, record[0:])
		}
	}

	file.Truncate(0)
	file.Seek(0, 0)
	csvWriter := csv.NewWriter(file)

	err = csvWriter.WriteAll(filteredRecords)
	if CheckError(err, "While writing to new file") {
		return
	}

	csvWriter.Flush()
	err = csvWriter.Error()
	if CheckError(err, "General writing error") {
		return
	}
}