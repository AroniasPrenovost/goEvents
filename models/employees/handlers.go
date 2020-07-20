package employees

import (
    "log"
    "net/http"
    "encoding/json"
    "goCRUD/config"
    _ "github.com/go-sql-driver/mysql"
)
 
func Employees(w http.ResponseWriter, r *http.Request) {
    
    switch r.Method {

        case "GET": 
            
            db := config.InitDB()
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

            db := config.InitDB()

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
            
            http.Redirect(w, r, "/", 301)
 
            log.Println("POST: added new employee")

        case "PUT":

            db := config.InitDB()

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

            db := config.InitDB()
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
