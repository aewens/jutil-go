package main

import (
    "encoding/json"
    "path/filepath"
    "io/ioutil"
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
    
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Error", r)
        }
    }()
    
    // Make sure arguments were passed
    if len(os.Args[1:]) == 0 {
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
    
    var context string = getContext(fileAction, createAction)
    
    switch context {
    case "file":
        var fileMode Mode = createMode("file", fileAction)
        var contents map[string]interface{} = fileMode.getContents()
        
        for key, val := range contents {
            switch value := val.(type) {
            case []interface{}:
                fmt.Println(key, fmt.Sprintf("json[%d]", len(value)))
                for i, v := range value {
                    fmt.Println(i, v)
                }
            default:
                fmt.Println(key, value)
            }
        }
        
        fmt.Println("Got contents")
    case "create":
        fmt.Println("Coming soon!")
    }
}

// Determine context of script
func getContext(fileAction, createAction string) string {
    var fileAssigned bool = len(fileAction) > 0
    var createAssigned bool = len(createAction) > 0
    
    if fileAssigned && createAssigned {
        panic("[0x02]: -f and -c cannot be used together")
    } else if fileAssigned {
        return "file"
    } else if createAssigned {
        return "create"
    } else {
        panic("[0x03]: Neither -f nor -c was selected")
    }
}

type Mode struct {
    action string
    path string
    value string
}

// Create mode structure
func createMode(params ...string) Mode {
    var action string
    var path string
    var value string
    
    if len(params) < 2 {
        panic("[0x04]: Not enough parameters passed!!")
    } else if len(params) > 3 {
        panic("[0x05]: Too many parameters passed!!")
    } else if len(params) == 2 {
        action = params[0]
        path = params[1]
        value = ""
    } else if len(params) == 3 {
        action = params[0]
        path = params[1]
        value = params[2]
    } else {
        panic("[0x06]: Logic error!!")
    }
    
    return Mode{action, path, value}
}

// Returns decoded contents of file
func (mode *Mode) getContents() map[string]interface{} {
    if mode.action != "file" {
        panic(fmt.Sprintf("[0x07]: '%s' is not a file mode", mode.action))
    }
    
    path, err := filepath.Abs(mode.path)
    if err != nil {
        panic(fmt.Sprintf("[0x08]: %s", err.Error()))
    }
    
    raw, err := ioutil.ReadFile(path)
    if err != nil {
        panic(fmt.Sprintf("[0x09]: %s", err.Error()))
    }
    
    var contents interface{}
    err = json.Unmarshal(raw, &contents)
    
    if err != nil {
        panic(fmt.Sprintf("[0x0A]: %s", err.Error()))
    }
    
    return contents.(map[string]interface{})
}
