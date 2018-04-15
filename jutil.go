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
    
    if len(os.Args[1:]) == 0 {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Error", r)
            }
        }()
        panic("[0x01]: No arguments passed")
    }
    
    // flag.String(param, default, description)
    // flag.StringVar(pointer*, param, default, description)
    
    var fileAction string
    var readAction string
    var editAction string
    var appendAction string
    var deleteAction string
    var iniAction string
    var createAction string
    var jsonAction string
    var valueAction string
    
    filePtr := flag.StringVar(&fileAction, "-f", "", "Path to JSON file")
    readPtr := flag.StringVar(&readAction, "-r", "", "Path to read")
    editPtr := flag.StringVar(&editAction, "-e", "", "Path to edit")
    appendPtr := flag.StringVar(&appendAction, "-a", "", "Path to append to")
    deletePtr := flag.StringVar(&deleteAction, "-d", "", "Path to delete")
    iniPtr := flag.StringVar(&iniAction, "-i", "", "INI output file name")
    createPtr := flag.StringVar(&createAction, "-c", "", "File to create from")
    jsonPtr := flag.StringVar(&jsonAction, "-j", "", "JSON output file name")
    valuePtr := flag.StringVar(&valueAction, "-v", "", "Value to use")
    
    
}
