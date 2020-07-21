package main

import (
    "log"  
    "net/http" 
    "os"
    "fmt"
    "encoding/json"
    "github.com/joho/godotenv"
    "time"
    "database/sql" 
    _ "github.com/go-sql-driver/mysql"
)

/* 
*
    Environment variables
*
*/
 

type Env struct {
	Port string // int? 
	DB_port string // int? 
	DB_name string
	DB_user string
    DB_password string
    API_endpoint string
	Admin_password string
}

func InitEnv() (Env) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ENV := Env{}
	ENV.Port = os.Getenv("PORT")
	ENV.DB_port = os.Getenv("DB_PORT")
	ENV.DB_user = os.Getenv("DB_USER")
	ENV.DB_name = os.Getenv("DB_NAME")
    ENV.DB_password = os.Getenv("DB_PASSWORD")
    ENV.API_endpoint = os.Getenv("API_ENDPOINT")
	ENV.Admin_password = os.Getenv("ADMIN_PASSWORD")

	return ENV
}

/* 
*
    Database connection
*
*/
 

// var DB *sql.DB
func InitDB() (db *sql.DB) {

    env := InitEnv()	
    var err error

    connection_string := env.DB_user + ":" + env.DB_password + "@tcp(" + env.API_endpoint + ":" + env.DB_port + ")/" + env.DB_name
    db, err = sql.Open("mysql", connection_string)
    if err != nil {
        log.Panic(err)
    }

    db.SetConnMaxLifetime(120*time.Second)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(5)

    if err = db.Ping(); err != nil {
        log.Panic(err)
    }

    return db
}


/* 
*
    API models 
*
*/


type Employee struct {
    Id    int
    Name  string
    City string
}


func Employees(w http.ResponseWriter, r *http.Request) {
    
    switch r.Method {

        case "GET": 

        fmt.Fprintf(w, "hello GET")
            
            db := InitDB()
            selDB, err := db.Query("SELECT * FROM Employees ORDER BY id ASC")
            if err != nil {
                panic(err.Error())
            }
            emp := Employee{}
            res := []Employee{}
            for selDB.Next() {
                var id int
                var name, city string
                err = selDB.Scan(&id, &name, &city)
                if err != nil {
                    panic(err.Error())
                }
                emp.Id = id
                emp.Name = name
                emp.City = city
                res = append(res, emp)
            }
            
            defer db.Close()
            
            jsonBytes, err := json.Marshal(res)
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte(err.Error()))
            }
        
            w.Header().Add("content-type", "application/json")
            w.WriteHeader(http.StatusOK)
            w.Write(jsonBytes)

            log.Println("GET: get all employees")

        case "POST":      

            db := InitDB()

            decoder := json.NewDecoder(r.Body)
            emp := Employee{}
            err := decoder.Decode(&emp)
            if err != nil {
                panic(err)
            }

            name := emp.Name
            city := emp.City
            insForm, err := db.Prepare("INSERT INTO Employees(name, city) VALUES(?,?)")
            if err != nil {
                panic(err.Error())
            }

            insForm.Exec(name, city)

            defer db.Close()


            jsonBytes, err := json.Marshal(insForm)
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte(err.Error()))
            }
        
            w.Header().Add("content-type", "application/json")
            w.WriteHeader(http.StatusOK)
            w.Write(jsonBytes)
 
            log.Println("POST: added new employee")

        case "PUT":

            db := InitDB()

            decoder := json.NewDecoder(r.Body)
            emp := Employee{}
            err := decoder.Decode(&emp)
            if err != nil {
                panic(err)
            }
        
            name := emp.Name
            city := emp.City
            id := emp.Id
            insForm, err := db.Prepare("UPDATE Employees SET name=?, city=? WHERE id=?")
            if err != nil {
                panic(err.Error())
            }
            insForm.Exec(name, city, id)
        
            defer db.Close()

            jsonBytes, err := json.Marshal(emp)
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte(err.Error()))
            }
        
            w.Header().Add("content-type", "application/json")
            w.WriteHeader(http.StatusOK)
            w.Write(jsonBytes)

            // http.Redirect(w, r, "/", 301)   
        
            log.Println("PUT: updated employee record")

        case "DELETE":

            db := InitDB()
            delForm, err := db.Query("DELETE FROM Employees")
            if err != nil {
                panic(err.Error())
            }

            log.Println(delForm)

            defer db.Close()
            http.Redirect(w, r, "/", 301)

            log.Println("DELETE: deleted    all employees")

    }
}


/* 
*
    Main
*
*/


func main() {
    
    http.HandleFunc("/employees", Employees)

    env := InitEnv()
    log.Println("Server started on: http://localhost:" + env.Port)
    http.ListenAndServe(":" + env.Port, nil)
}
