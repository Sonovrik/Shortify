package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB
const dbConfig =
	"host=postgress_container " +
	"port=5432 " +
	"dbname=postgres_db " +
	"sslmode=disable " +
	"user=root " +
	"password=root"

func	getDB() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("postgres", dbConfig)
		if err != nil {
			log.Fatalf("Unable to connect to database: %s", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("Unable to ping to database: %s", err)
		}
		fmt.Println("DB OK")
	}
	return db
}

func test(w http.ResponseWriter, r *http.Request){

	db = getDB()

	uri := r.URL.Query().Get("uri")
	if len(uri) == 0 {
		w.WriteHeader(http.StatusForbidden)
	}


}

func	getEnv(key string) string {
	env, exists := os.LookupEnv(key)
	if !exists {
		return ""
	}
	return env
}

func main() {
	port := getEnv("SHORTIFY_CONTAINER_PORT")
	if len(port) == 0 {
		panic("Envs not found :(")
	}

	db = getDB()


	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir("./html/")))
	http.HandleFunc("/asd", test)

	log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))
}
