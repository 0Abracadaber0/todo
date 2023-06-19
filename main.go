package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id    int
	Name  string
	Text  string
	Start string
	End   string
}

var database *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Query("SELECT * FROM todo")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	todos := []Todo{}

	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.Id, &todo.Name, &todo.Text, &todo.Start, &todo.End)
		if err != nil {
			fmt.Println(err)
			continue
		}
		todos = append(todos, todo)
	}

	tmpl, _ := template.ParseFiles("view/index.html")
	tmpl.Execute(w, todos)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		name := r.FormValue("name")
		text := r.FormValue("text")
		start := r.FormValue("start")
		end := r.FormValue("end")

		_, err = database.Exec("INSERT INTO todo (name, text, start, end) VALUES (?, ?, ?, ?)", name, text, start, end)
		if err != nil {
			fmt.Println(err)
		}

		http.Redirect(w, r, "/", 301)

	} else {
		http.ServeFile(w, r, "view/create.html")
	}
}

func EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	row := database.QueryRow("SELECT * FROM todo WHERE todo_id = ?", id)
	todo := Todo{}
	err := row.Scan(&todo.Id, &todo.Name, &todo.Text, &todo.Start, &todo.End)

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		tmpl, _ := template.ParseFiles("view/edit.html")
		tmpl.Execute(w, todo)
	}

}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	text := r.FormValue("text")
	start := r.FormValue("start")
	end := r.FormValue("end")

	_, err = database.Exec("UPDATE todo SET name = ?, text = ?, start = ?, end = ? WHERE todo_id = ?", name, text, start, end, id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	_, err := database.Exec("DELETE FROM todo WHERE todo_id = ?", id)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", os.Getenv("LOGIN")+":"+os.Getenv("PASS")+"@/"+os.Getenv("DB_NAME"))

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/create", CreateHandler)
	router.HandleFunc("/delete/{id:[0-9]+}", DeleteHandler)
	router.HandleFunc("/edit/{id:[0-9]+}", EditPage).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", EditHandler).Methods("POST")

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", router)
}
