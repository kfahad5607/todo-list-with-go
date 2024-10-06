package todo

import (
	"encoding/json"
	"os"
	"time"
)


type JsonStore struct {
	StoreName string
}

func NewJsonStore(storeName string) *JsonStore{
	store := JsonStore{StoreName: storeName + ".json"}
	store.ensureStoreExists()

	return &store
}

func (store JsonStore) ensureStoreExists() error {
	if checkFileExists(store.StoreName) {
		return nil
	}

	file, err := os.OpenFile(store.StoreName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if CheckError(err, "While opening a new file") {
		return err
	}
	defer file.Close()

	w := json.NewEncoder(file)
	err = w.Encode([]DataItem{})

	if CheckError(err, "While writing to a json file") {
		return err
	}

	return nil
}

func (store JsonStore) ReadItems(showAll bool) (items []DataItem) {
	items = []DataItem{}
	data, err := os.ReadFile(store.StoreName)
	if CheckError(err, "While opening file to read") {
		return
	}
	
	allItems := []DataItem{}
	err = json.Unmarshal(data, &allItems)
	if CheckError(err, "While unmarshaling json file") {
		return
	}

	if showAll {
		items = allItems
		return
	}

	for _, item := range allItems {
		if *item.IsComplete {
			continue
		}
		items = append(items, item)
	}

	return
}

func (store JsonStore) ReadItem(id int) (item DataItem) {
	records := store.ReadItems(true)

	for _, record := range records {
		if record.ID == id {
			item = record
			return
		}
	}

	return
}

func (store JsonStore) CreateItem(itemDescription string) (item DataItem) {
	items := store.ReadItems(true)

	lastRecordId := 0
	if len(items) > 0 {
		lastRecord := items[len(items)-1]
		lastRecordId = lastRecord.ID
	}

	isComplete := false
	createdAt := time.Now().Format(time.RFC3339)
	item = DataItem{Description: &itemDescription, CreatedAt: &createdAt, IsComplete: &isComplete}
	item.ID = lastRecordId + 1

	items = append(items, item)

	updatedData, err := json.Marshal(items)
	if CheckError(err, "While marshalling json data") {
		return
	}

	err = os.WriteFile(store.StoreName, updatedData, os.ModePerm)
	if CheckError(err, "While writing json data to file") {
		return
	}

	return
}

func (store JsonStore) UpdateItem(id int, item DataItem) (updatedItem DataItem) {
	items := store.ReadItems(true)

	for i := range items {
		if items[i].ID == id {
			if item.Description != nil {
				items[i].Description = item.Description
			}
			if item.IsComplete != nil {
				items[i].IsComplete = item.IsComplete
			}
			updatedItem = items[i]

			break
		}
	}

	updatedData, err := json.Marshal(items)
	
	if CheckError(err, "While marshalling json data") {
		return
	}

	err = os.WriteFile(store.StoreName, updatedData, os.ModePerm)
	if CheckError(err, "While writing json data to file") {
		return
	}

	return
}

func (store JsonStore) DeleteItem(id int) {
	items := store.ReadItems(true)

	filteredItems := []DataItem{}
	for _, item := range items {
		if item.ID != id {
			filteredItems = append(filteredItems, item)
		}
	}

	updatedData, err := json.Marshal(filteredItems)
	if CheckError(err, "While marshalling json data") {
		return
	}

	err = os.WriteFile(store.StoreName, updatedData, os.ModePerm)
	if CheckError(err, "While writing json data to file") {
		return
	}
}