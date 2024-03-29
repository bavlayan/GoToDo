package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/bavlayan/GoToDo/pkg/models"
	"github.com/bavlayan/GoToDo/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

type contextKey string

var contextKeyUser = contextKey("user")

type application struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	todoitems interface {
		Save(string) (int, error)
		GetDaily() ([]*models.TodoItems, error)
		Get(string) (*models.TodoItems, error)
	}
	templateCache map[string]*template.Template
	session       *sessions.Session
	user          interface {
		Insert(string, string, string) error
		Authenticate(string, string) (string, error)
		Get(string) (*models.User, error)
	}
}

func main() {

	addr := flag.String("addr", ":4444", "HTTP network address")
	dsn := flag.String("dsn", "todoapp_dbuser:password@/gotodo_db?parseTime=true", "MySQL database connection string")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret for session")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteStrictMode

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		todoitems: &mysql.TodoItemModel{
			DB: db,
		},
		templateCache: templateCache,
		session:       session,
		user: &mysql.UserModel{
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
