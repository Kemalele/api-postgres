package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "123"
// 	dbname   = "mydb"
// )

type author struct {
	ID    int
	Name  string
	Alias string
	Spec  string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/author", authors)
	mux.HandleFunc("/", main_handler)
	psqlInfo := fmt.Sprintf("host=db user=postgres password=123 dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	// host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("ping err:", err)
	}
	fmt.Println("http://localhost:3000")
	http.ListenAndServe(":3000", mux)
}

func main_handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello buddy!"))
}

func authors(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/author" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	var avtor author
	err := json.NewDecoder(r.Body).Decode(&avtor)
	if err != nil {
		log.Println(err)
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println(err)
	fmt.Fprint(w, "Author added successfully")

	// sqlStatement := `
	// INSERT INTO users (age, email, first_name, last_name)
	// VALUES (30, 'jon@calhoun.io', 'Jonathan', 'Calhoun')`
	// _, err = db.(sqlStatement)
	// if err != nil {
	// 	panic(err)
	// }
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal("ping err:", err)
	// }

}
