package main

import (
    "flag"
	"fmt"
	"os"
	"time"
	"strconv"
    
    _ "github.com/mattn/go-sqlite3"
)

func main() {

    // Handle flags
        
    var action [2]string
    
    flag.StringVar(&action[0], "n", "", "Add flag for action, e.g. -n for new")
    flag.StringVar(&action[0], "d", "", "Add flag for action, e.g. -d for delete")
    flag.Parse()
    
    // Set up configuration, i.e. database path only currently, set in CALLY_DB env,
    // followed by database creation (if not exists) based on that path.
    
    var config Configuration
    config = Configure()
    db := config.databaseURI
    
    createDb(db)
    
    // Handle args or lack there of, main process of program. 

    // if there are no args, default output
    if action[0] == "" && len(os.Args) <= 1 {
    	t := time.Now()
    	fmt.Printf("\nHi! It's %s the %d-%d %d, %d:%d\nand here are your current notes:\n", t.Weekday(), t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
		readNotes(db)
		fmt.Println("\nAnd you current events:")
		printEvents(db)
    } else {
    	for arg := range os.Args {
    		// Skip the first arg, i.e. program path... 
    		if arg == 0 {
    			continue
    		}
    		 // for "new" items
    		if os.Args[arg] == "-n" {
    			if os.Args[arg + 1] != "" {
	    			switch {
	    			// Handle new notes
	    				case os.Args[arg + 1] == "n":
	    					if os.Args[arg + 2] != "" {
	    						createNote(os.Args[arg + 2], db)
	    						continue
	    					} else {
	    					fmt.Println("Invalid input.\nUsage:\n")
	    					printUsage()
	    					break
	    					}
	    				 // Handle new events
	    				case os.Args[arg + 1] == "e":
	    					if os.Args[arg + 2] != "" && os.Args[arg + 3] != "" && os.Args[arg + 4] != "" {
	    						createEvent(os.Args[arg + 2], os.Args[arg + 3], os.Args[arg + 4], db)
	    						continue
	    					} else {
	    					fmt.Println("Invalid input.\nUsage:\n")
	    					printUsage()
	    					break
	    					}
	    				// invalid input.. 
	    				default: 
	    				    fmt.Println("Invalid input.\nUsage:\n")
	    					printUsage()
	    					break
	    			}
    			}  else {
				printUsage()
				break
			}
    		// For deleting items	
			} else if os.Args[arg] == "-d" {
				switch {
					// For deleting notes, takes note id and db
					case os.Args[arg + 1] == "n":
						if os.Args[arg + 2] != "" {
							val, _ := strconv.Atoi(os.Args[arg + 2])
							deleteNote(val, db)
							continue
						} else {
							printUsage()
							break
						}
					// For deleting events, takes title and db
					case os.Args[arg + 1] == "e":
						if os.Args[arg + 2] != "" {
							deleteEvent(os.Args[arg + 2], db)
						}
				}
			}
    	}
    }
}

// Simple function for outputting usage, specifically called when only arg is 'help' 
// or upon invalid input. 

func printUsage() {
	fmt.Println("cally -n n \"use this to write a new note\"")
	fmt.Println("cally -n e \"christmas eve\" \"use this to create an event for chrstimas eve!\" \"2020-12-24\"")
	fmt.Println("This deleted note number 24: ")
	fmt.Println("cally -d n 24")
	fmt.Println("And this deleted christmas eve (the event, not actually christmas eve):")
	fmt.Println("cally -d e \"christmas eve\"")
	fmt.Println("Finally, simply run 'cally' without any arguments to print all notes and proximate events. Have fun!")
}
