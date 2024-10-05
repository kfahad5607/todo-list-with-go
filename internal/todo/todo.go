package todo

type DataStore interface {
	ReadItems(showAll bool) []DataItem
	ReadItem(id int) DataItem
	CreateItem(itemDescription string) DataItem
	UpdateItem(id int, item DataItem) DataItem
	DeleteItem(id int)
}