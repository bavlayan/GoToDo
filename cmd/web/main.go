package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/bavlayan/GoToDo/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	todoitems *mysql.TodoItemModel
}

func main() {

	addr := flag.String("addr", ":4444", "HTTP network address")
	dsn := flag.String("dsn", "todoapp_dbuser:password@/gotodo_db?parseTime=true", "MySQL database connection string")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		todoitems: &mysql.TodoItemModel{
			DB: db,
		},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
