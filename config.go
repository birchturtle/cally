package main

import (
    "fmt"
    "os"
)

func Configure() Configuration {
    var calEnv Configuration
    var ok bool
    calEnv.databaseURI, ok = os.LookupEnv("CALLY_DB")
    
    if !ok {
        // set env CALLY_DB
        fmt.Println("Path to database not set, please enter one.")
                
        var path string
        fmt.Scanln(&path)
        os.Setenv("CALLY_DB", path)
        calEnv.databaseURI, _ = os.LookupEnv("CALLY_DB")
    }

	home,_ := os.UserHomeDir()
    calEnv.databaseURI = home + "/" + calEnv.databaseURI
    
    return calEnv
}

type Configuration struct {
    databaseURI string
}
