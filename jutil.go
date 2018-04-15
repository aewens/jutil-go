package main

import (
    // "encoding/json"
    // "io/ioutil"
    "flag"
    "fmt"
    "os"
)

func main() {
    // TODO
    // [ ]: -f (file)
    // [ ]: -r (read)
    // [ ]: -e (edit)
    // [ ]: -a (append)
    // [ ]: -d (delete)
    // [ ]: -i (ini)
    // [ ]: -c (create)
    // [ ]: -j (json)
    // [ ]: -v (value)
    
    // Make sure arguments were passed
    if len(os.Args[1:]) == 0 {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Error", r)
            }
        }()
        panic("[0x01]: No arguments passed")
    }
    
    // Define flag value variables
    var fileAction string
    var readAction string
    var editAction string
    var appendAction string
    var deleteAction string
    var iniAction string
    var createAction string
    var jsonAction string
    var valueAction string
    
    // NOTE
    // flag.String(param, default, description)
    // flag.StringVar(pointer*, param, default, description)
    
    // Flag defitions
    flag.StringVar(&fileAction, "f", "", "Path to JSON file")
    flag.StringVar(&readAction, "r", "", "Path to read")
    flag.StringVar(&editAction, "e", "", "Path to edit")
    flag.StringVar(&appendAction, "a", "", "Path to append to")
    flag.StringVar(&deleteAction, "d", "", "Path to delete")
    flag.StringVar(&iniAction, "i", "", "INI output file name")
    flag.StringVar(&createAction, "c", "", "File to create from")
    flag.StringVar(&jsonAction, "j", "", "JSON output file name")
    flag.StringVar(&valueAction, "v", "", "Value to use")
    
    // Parse all flags
    flag.Parse()
    
    // Get output
    fmt.Println("fileAction:", fileAction)
    fmt.Println("readAction:", readAction)
    fmt.Println("editAction:", editAction)
    fmt.Println("appendAction:", appendAction)
    fmt.Println("deleteAction:", deleteAction)
    fmt.Println("iniAction:", iniAction)
    fmt.Println("createAction:", createAction)
    fmt.Println("jsonAction:", jsonAction)
    fmt.Println("valueAction:", valueAction)
    fmt.Println("remainder:", flag.Args())
}
