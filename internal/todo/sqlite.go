package todo

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteStore struct {
	StoreName string
	DB *sql.DB
}

func NewSqliteStore(storeName string) *SqliteStore{
	store := SqliteStore{DB: nil, StoreName: storeName + ".db"}

	db := store.connectDB()
	defer store.closeDB()

	query := `
		CREATE table IF NOT EXISTS tasks (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, description TEXT NOT NULL, createdAt DATETIME DEFAULT CURRENT_TIMESTAMP, isComplete BOOLEAN DEFAULT 0);
	`
	db.Exec(query)

	store.DB = db

	return &store
}

func (store *SqliteStore)connectDB() *sql.DB{
	db, err := sql.Open("sqlite3", store.StoreName)
	if CheckError(err, "While connecting to sqlite DB") {
		log.Fatal(err)
	}
	store.DB = db

	return store.DB
}

func (store *SqliteStore)closeDB(){
	err := store.DB.Close()

	if CheckError(err, "While closing DB") {
		log.Fatal(err)
	}
}

func (store SqliteStore) ReadItems(showAll bool) (items []DataItem) {
	db := store.connectDB()
	defer store.closeDB()

	query := `SELECT id, description, createdAt, isComplete from tasks`
	if showAll{
		query = query + ";"
	}else{
		query = query + " where isComplete = $1;"
	}

	rows, err := db.Query(query, 0)
	if CheckError(err, "While getting tasks list"){
		log.Fatal(err)
	}

	for rows.Next(){
		item := DataItem{}
		err := rows.Scan(&item.ID, &item.Description, &item.CreatedAt, &item.IsComplete)
		if CheckError(err, "while scanning row") {
			log.Fatal(err)
		}

		items = append(items, item)
	}

	return
}

func (store SqliteStore) ReadItem(id int) (item DataItem) {
	db := store.connectDB()
	defer store.closeDB()

	query := `
		SELECT id, description, createdAt, isComplete from tasks where id = $1;
	`
	row := db.QueryRow(query, id)
	err := row.Scan(&item.ID, &item.Description, &item.CreatedAt, &item.IsComplete)
	if CheckError(err, "while scanning row") {
		log.Fatal(err)
	}

	return
}

func (store SqliteStore) CreateItem(itemDescription string) (item DataItem) {
	db := store.connectDB()
	defer store.closeDB()

	query := `
		INSERT into tasks (description) VALUES ($1);
	`

	_, err := db.Exec(query, itemDescription)
	if CheckError(err, "while inserting row") {
		log.Fatal(err)
	}

	query = `SELECT id, description, createdAt, isComplete FROM tasks WHERE id = last_insert_rowid();`
	row := db.QueryRow(query)
	err = row.Scan(&item.ID, &item.Description, &item.CreatedAt, &item.IsComplete)
	if CheckError(err, "while scanning row") {
		log.Fatal(err)
	}
	
	return
}

func (store SqliteStore) UpdateItem(id int, item DataItem) (updatedItem DataItem) {
	db := store.connectDB()
	defer store.closeDB()

	query := `UPDATE tasks set`

	var err error
	if item.Description != nil && item.IsComplete != nil {
		query += " description=$1, isComplete=$2 where id=$3;"
		_, err = db.Exec(query, *item.Description, *item.IsComplete, id)
	}else if item.Description != nil{
		query += " description=$1 where id=$2;"
		_, err = db.Exec(query, *item.Description, id)
	}
	if item.IsComplete != nil {
		query += " isComplete=$1 where id=$2;"
		_, err = db.Exec(query, *item.IsComplete, id)
	}

	if CheckError(err, "while updating row") {
		log.Fatal(err)
	}

	query = `SELECT id, description, createdAt, isComplete FROM tasks WHERE id = $1;`
	row := db.QueryRow(query, id)
	err = row.Scan(&updatedItem.ID, &updatedItem.Description, &updatedItem.CreatedAt, &updatedItem.IsComplete)
	if CheckError(err, "while scanning row") {
		log.Fatal(err)
	}

	return
}

func (store SqliteStore) DeleteItem(id int) {
	db := store.connectDB()
	defer store.closeDB()

	query := `DELETE from tasks where id=$1`
	_, err := db.Exec(query, id)
	if CheckError(err, "while deleting row") {
		log.Fatal(err)
	}
}