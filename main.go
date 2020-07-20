package main

import (
    "log"  
    "net/http"
    "goCRUD/config"
    "goCRUD/models/employees"
)
 
func main() {
    
    http.HandleFunc("/employees", employees.Employees)

    env := config.InitEnv()
    log.Println("Server started on: http://localhost:" + env.Port)
    http.ListenAndServe(":" + env.Port, nil)
}
