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
* Go servemux doesn't support method based routing or semantic URLs with variables in them.
* The `html/template` package automatically escapes any data that is yielded between `{{}}` tags such as avoiding XSS attacks
* Servemux has two types of URL patters
    * Subtree paths
    * Fixed paths


### References
