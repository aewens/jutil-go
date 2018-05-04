package main

import (
    "encoding/json"
    "path/filepath"
    "io/ioutil"
    "strings"
    "flag"
    "fmt"
    "os"
)

func main() {
    // TODO
    // [*]: -f (file)
    // [ ]: -r (read)
    // [ ]: -e (edit)
    // [ ]: -a (append)
    // [ ]: -d (delete)
    // [ ]: -i (ini)
    // [ ]: -c (create)
    // [ ]: -j (json)
    // [ ]: -v (value)
    // [*]: -s (see)
    
    // Catch panics here
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
    var seeAction bool
    
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
    flag.BoolVar(&seeAction, "s", false, "See contents of file")
    
    // Parse all flags
    flag.Parse()
    
    // Handle by context
    rootActions := make(map[string]bool)
    rootActions["file"] = len(fileAction) > 0
    rootActions["create"] = len(createAction) > 0
    
    switch getContext(rootActions) {
    case "file":
        var fileMode Mode = createMode("file", fileAction)
        var contents map[string]interface{} = fileMode.getContents()
        
        subActions := make(map[string]bool)
        subActions["read"] = len(readAction) > 0
        subActions["edit"] = len(editAction) > 0
        subActions["append"] = len(appendAction) > 0
        subActions["delete"] = len(deleteAction) > 0
        subActions["see"] = seeAction
        
        switch getContext(subActions) {
        case "read":
            fmt.Println("Read is coming soon!")
        case "edit":
            fmt.Println("Edit is coming soon!")
        case "append":
            fmt.Println("Append is coming soon!")
        case "delete":
            fmt.Println("Delete is coming soon!")
        case "see":
            seeContents(contents, 0)
        }
        
        fmt.Println("Success [0x00]")
    case "create":
        fmt.Println("Create is coming soon!")
    }
}

// Count how many bools are true
func countBool(tests map[string]bool) (int, []string) {
    var count int = 0
    var names []string
    for name, test := range tests {
        if test {
            count = count + 1
            names = append(names, name)
        }
    }
    return count, names
}

// Determine context of script
func getContext(rootActions map[string]bool) string {
    truthCount, trueNames := countBool(rootActions)
    keys := make([]string, len(rootActions))
    var i int = 0
    for key, _ := range rootActions {
        keys[i] = key
        i = i + 1
    }
    
    if (truthCount > 1) {
        allFlags := make([]string, len(trueNames))
        for n, name := range trueNames {
            allFlags[n] = fmt.Sprintf("-%s", string(name[0]))
        }
        
        var andFlags string = strings.Join(allFlags, " and ")
        panic(fmt.Sprintf("[0x02]: %s cannot be used together", andFlags))
    } else if (truthCount == 0) {
        allFlags := make([]string, len(keys))
        for k, key := range keys {
            allFlags[k] = fmt.Sprintf("-%s", string(key[0]))
        }
        
        var norFlags string = strings.Join(allFlags, " nor ")
        panic(fmt.Sprintf("[0x03]: Neither %s was selected", norFlags))
    }
    
    return trueNames[0]
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

func printIndent(indent int) {
    for i := 0; i < indent; i++ {
        fmt.Print("\t")
    }
}


func seeLists(key, value interface{}, ind int) {
    printIndent(ind)
    fmt.Println(key, "[list]")
    for i, v := range value.([]interface{}) {
        switch v.(type) {
        case map[string]interface{}:
            printIndent(ind + 1)
            fmt.Println(fmt.Sprintf("%d ", i))
            seeContents(v, ind + 2)
        case []interface{}:
            seeLists(i, v, ind + 1)
        default:
            printIndent(ind + 1)
            fmt.Println(fmt.Sprintf("%d ", i))
            printIndent(ind + 2)
            fmt.Println(v, fmt.Sprintf("[%T]", v))
        }
    }
}

func seeContents(contents interface{}, indent int) {
    for key, val := range contents.(map[string]interface{}) {
        switch value := val.(type) {
        case map[string]interface{}:
            printIndent(indent)
            fmt.Println(key, "[dict]")
            seeContents(value, indent + 1)
        case []interface{}:
            seeLists(key, value, indent)
        default:
            printIndent(indent)
            fmt.Println(key, value, fmt.Sprintf("[%T]", value))
        }
    }
}
