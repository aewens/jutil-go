package main

import (
    "encoding/json"
    "path/filepath"
    "io/ioutil"
    "strconv"
    "strings"
    "flag"
    "fmt"
    "os"
)

func main() {
    // TODO
    // [*]: -f (file)
    // [*]: -r (read)
    // [*]: -e (edit)
    // [ ]: -a (append)
    // [*]: -d (delete)
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
        var contents Contents = fileMode.Get()
        
        subActions := make(map[string]bool)
        subActions["read"] = len(readAction) > 0
        subActions["edit"] = len(editAction) > 0
        subActions["append"] = len(appendAction) > 0
        subActions["delete"] = len(deleteAction) > 0
        subActions["see"] = seeAction
        
        switch getContext(subActions) {
        case "read":
            contents.Read(parsePath(readAction), 0)
        case "edit":
            if valueAction == "" {
                panic(fmt.Sprintf("[0x11]: Value is missing"))
            }
            contents.Edit(parsePath(editAction), valueAction)
            fileMode.Save(contents)
        case "append":
            if valueAction == "" {
                panic(fmt.Sprintf("[0x11]: Value is missing"))
            }
            fmt.Println("Append is coming soon!")
        case "delete":
            // fmt.Println("Delete is coming soon!")
            contents.Delete(parsePath(deleteAction))
            // contents.See(0)
            fileMode.Save(contents)
        case "see":
            contents.See(0)
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

type Contents struct {
    body interface{}
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
func (mode *Mode) Get() Contents {
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
    
    return Contents{contents.(map[string]interface{})}
}

// Returns decoded contents of file
func (mode *Mode) Save(contents Contents) {
    if mode.action != "file" {
        panic(fmt.Sprintf("[0x0B]: '%s' is not a file mode", mode.action))
    }

    path, err := filepath.Abs(mode.path)
    if err != nil {
        panic(fmt.Sprintf("[0x0C]: %s", err.Error()))
    }
    
    bytes, err := json.MarshalIndent(contents.body, "", "    ")
    if err != nil {
        panic("[0xD]: Failed to encode contents")
    }
    
    err = ioutil.WriteFile(path, bytes, 0644)
    if err != nil {
        panic(fmt.Sprintf("[0x0E]: %s", err.Error()))
    }
    
    fmt.Println("Saved file!")
}

// Handle indentation for displaying JSON
func printIndent(indent int) {
    for i := 0; i < indent; i++ {
        fmt.Print("\t")
    }
}

// Converts path to array for searching
func parsePath(path string) []string {
    var parsed []string = strings.Split(path, "/")
    return parsed
}

// Renders the JSON of the file
func (c *Contents) See(indent int) {
    var body interface{} = c.body
    if dict, ok := body.(map[string]interface{}); ok {
        for key, val := range dict {
            switch value := val.(type) {
            case map[string]interface{}:
                printIndent(indent)
                fmt.Println(key, "[dict]")
                newContents := &Contents{value}
                newContents.See(indent + 1)
            case []interface{}:
                printIndent(indent)
                fmt.Println(key, "[list]")
                newContents := &Contents{value}
                newContents.See(indent + 1)
            default:
                printIndent(indent)
                fmt.Println(key, value, fmt.Sprintf("[%T]", value))
            }
        }
    } else if list, ok := body.([]interface{}); ok {
        for i, val := range list {
            switch value := val.(type) {
            case map[string]interface{}:
                printIndent(indent)
                fmt.Println(fmt.Sprintf("%d ", i))
                newContents := &Contents{value}
                newContents.See(indent + 1)
            case []interface{}:
                printIndent(indent)
                fmt.Println(fmt.Sprintf("%d ", i))
                newContents := &Contents{value}
                newContents.See(indent + 1)
            default:
                printIndent(indent)
                fmt.Println(fmt.Sprintf("%d ", i))
                printIndent(indent + 1)
                fmt.Println(value, fmt.Sprintf("[%T]", value))
            }
        }
    }
}

// View value from JSON file based on provided path
func (c *Contents) Read(path []string, indent int) {
    var search string
    var remaining []string
    var found bool = false
    
    if path != nil {
        search = path[0]
        if len(path) > 1 {
            remaining = path[1:]
        }
    }
    
    var body interface{} = c.body
    if dict, ok := body.(map[string]interface{}); ok {
        for key, val := range dict {
            if path != nil && key != search {
                continue
            }
            found = true
            
            switch value := val.(type) {
            case map[string]interface{}:
                printIndent(indent)
                fmt.Println(key, "[dict]")
                newContents := &Contents{value}
                newContents.Read(remaining, indent + 1)
            case []interface{}:
                printIndent(indent)
                fmt.Println(key, "[list]")
                newContents := &Contents{value}
                newContents.Read(remaining, indent + 1)
            default:
                printIndent(indent)
                fmt.Println(key, value, fmt.Sprintf("[%T]", value))
            }
        }
    } else if list, ok := body.([]interface{}); ok {
        for i, val := range list {
            if path != nil && strconv.Itoa(i) != search {
                continue
            }
            found = true
            
            switch value := val.(type) {
            case map[string]interface{}:
                printIndent(indent)
                fmt.Println(fmt.Sprintf("%d ", i))
                newContents := &Contents{value}
                newContents.Read(remaining, indent + 1)
            case []interface{}:
                printIndent(indent)
                fmt.Println(fmt.Sprintf("%d ", i))
                newContents := &Contents{value}
                newContents.Read(remaining, indent + 1)
            default:
                printIndent(indent)
                fmt.Println(fmt.Sprintf("%d ", i))
                printIndent(indent + 1)
                fmt.Println(value, fmt.Sprintf("[%T]", value))
            }
        }
    }
    
    if !found {
        panic(fmt.Sprintf("[0x0F]: '%s' was not found", search))
    }
}

// Delete value from JSON file based on provided path
func (c *Contents) Delete(path []string) interface{} {
    var search string
    var remaining []string
    var found bool = false
    
    if path != nil {
        search = path[0]
        if len(path) > 1 {
            remaining = path[1:]
        }
    }
    
    var body interface{} = c.body
    if dict, ok := body.(map[string]interface{}); ok {
        for key, val := range dict {
            if path != nil && key != search {
                continue
            }
            found = true
            
            if remaining == nil {
                delete(dict, key)
                fmt.Println(fmt.Sprintf("Deleting %s", key))
                return body
            }
            
            switch value := val.(type) {
            case map[string]interface{}:
                newContents := &Contents{value}
                dict[key] = newContents.Delete(remaining)
            case []interface{}:
                newContents := &Contents{value}
                dict[key] = newContents.Delete(remaining)
            default:
                continue
            }
        }
    } else if list, ok := body.([]interface{}); ok {
        for i, val := range list {
            if path != nil && strconv.Itoa(i) != search {
                continue
            }
            found = true
            
            if remaining == nil {
                list = append(list[:i], list[i+1:]...)
                body = list
                fmt.Println(fmt.Sprintf("Deleting %d", i))
                return body
            }
            
            switch value := val.(type) {
            case map[string]interface{}:
                newContents := &Contents{value}
                list[i] = newContents.Delete(remaining)
            case []interface{}:
                newContents := &Contents{value}
                list[i] = newContents.Delete(remaining)
            default:
                continue
            }
        }
    }
    
    if !found {
        panic(fmt.Sprintf("[0x10]: '%s' was not found", search))
    }
    
    return body
}

// Edit value from JSON file based on provided path
func (c *Contents) Edit(path []string, entry string) {
    var search string
    var remaining []string
    var found bool = false
    
    if path != nil {
        search = path[0]
        if len(path) > 1 {
            remaining = path[1:]
        }
    }
    
    var body interface{} = c.body
    if dict, ok := body.(map[string]interface{}); ok {
        for key, val := range dict {
            if path != nil && key != search {
                continue
            }
            found = true
            
            if remaining == nil {
                switch dict[key].(type) {
                case float64:
                    if f, err := strconv.ParseFloat(entry, 64); err == nil {
                        dict[key] = f
                    } else {
                        dict[key] = entry
                    }
                case bool:
                    if entry == "true" {
                        dict[key] = true
                    } else {
                        dict[key] = false
                    }
                default:
                    dict[key] = entry
                }
                fmt.Println(fmt.Sprintf("Editing %s", key))
                continue
            }
            
            switch value := val.(type) {
            case map[string]interface{}:
                newContents := &Contents{value}
                newContents.Edit(remaining, entry)
            case []interface{}:
                newContents := &Contents{value}
                newContents.Edit(remaining, entry)
            default:
                continue
            }
        }
    } else if list, ok := body.([]interface{}); ok {
        for i, val := range list {
            if path != nil && strconv.Itoa(i) != search {
                continue
            }
            found = true
            
            if remaining == nil {
                switch list[i].(type) {
                case float64:
                    if f, err := strconv.ParseFloat(entry, 64); err == nil {
                        list[i] = f
                    } else {
                        list[i] = entry
                    }
                case bool:
                    if entry == "true" {
                        list[i] = true
                    } else {
                        list[i] = false
                    }
                default:
                    list[i] = entry
                }
                fmt.Println(fmt.Sprintf("Editing %d", i))
                continue
            }
            
            switch value := val.(type) {
            case map[string]interface{}:
                newContents := &Contents{value}
                newContents.Edit(remaining, entry)
            case []interface{}:
                newContents := &Contents{value}
                newContents.Edit(remaining, entry)
            default:
                continue
            }
        }
    }
    
    if !found {
        panic(fmt.Sprintf("[0x12]: '%s' was not found", search))
    }
}
