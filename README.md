# GoToDo
GoToDo is "to do" app that is developed for learning golang.

### Tips for Golang
* Handlers as being a bit like controllers in MVC
* `log.fatal()` function also call `os.Exit(1)`
* `http.FileServer` handler which can you use serve files over HTTP from a specifi directory
* The **cmd** directory contains the application-specific code for the executable applicaitons in the project
* The **pkg** directory contains the non-application specific code used in the project
* The **ui** directory contains the user-interface assets used by web applicaton
* Custom loggers created by `http.FileServer` are **concurrency-safe**.
* model can be considered as service layer or data access layer
* Every reuest is handled in it's own goroutine
* `DB.Query()` returns multiple rows
* `DB.QueryRow()` returns a single row
* `DB.Exec()` doesn't return any rows
* [Alice](https://github.com/justinas/alice) provides a convenient way to chain your HTTP middleware functions and the app handler
* [nosurf](https://github.com/justinas/nosurf) is an HTTP package for Go that helps you prevent Cross-Site Request Forgery attacks. It acts like a middleware and therefore is compatible with basically any Go HTTP application.
* Go servemux doesn't support method based routing or semantic URLs with variables in them.
* In essence every `http.Request` that our handlers process has a `context.Context` object embedded in it, which we can use to store information during the lifetime of the request
* It's important to emphasize that request context should only be used store information relevant to the lifetime of specific request. The Go documentation for `context.Context` warns. **_Use context Values only for request-scoped data transits processes and APIs_** That means you should not use it to pass dependencies that exist outside of the lifetime of a request - like loggers, template caches and your database connection pool - to your middleware and handlers
* You can run only specific test by using the `-run` flag
* The `html/template` package automatically escapes any data that is yielded between `{{}}` tags such as avoiding XSS attacks
* Servemux has two types of URL patters
    * Subtree paths
    * Fixed paths

### SQL Table Scripts
GoToDO app has two db tables.

* tbl_todoitems
```sql
CREATE TABLE gotodo_db.tbl_todoitems (
  id VARCHAR(36) NOT NULL,
  completed TINYINT(1) NOT NULL DEFAULT 0,
  created_date DATETIME NULL DEFAULT NOW(),
  description VARCHAR(1024) NULL,
  deleted TINYINT(1) NOT NULL DEFAULT 0,
PRIMARY KEY (id));
```
* tbl_users
```sql
CREATE TABLE gotodo_db.tbl_users (
  id VARCHAR(36) NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  hashed_password VARCHAR(60) NOT NULL,
  created_date DATETIME NULL DEFAULT NOW(),
PRIMARY KEY (id));

ALTER TABLE gotodo_db.tbl_users ADD CONSTRAINT users_uc_email UNIQUE (email);
```

### References

Let's Go! Learn to build professional web application with Go | [Alex Edvards](https://lets-go.alexedwards.net/)
