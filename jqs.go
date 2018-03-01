package main

import (
    "os"
    "fmt"
    "bufio"
    "encoding/json"
    "net/url"
)

func main() {

    var unparsedJson string

    fi, err := os.Stdin.Stat()
    if err != nil {
        fmt.Println("Error reading from stdin")
        os.Exit(1)
    }

    if fi.Mode() & os.ModeNamedPipe == 0 {
        // no piped data, read first arg

        if len(os.Args) != 2 {
            fmt.Println("Must provide exactly one argument")
            os.Exit(1)
        }

        unparsedJson = fmt.Sprintf("%v", os.Args[1])
    } else {
        // pipe, read from stdin

        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            unparsedJson = fmt.Sprintf("%v", scanner.Text())
        }
    }

    var objmap map[string]interface{}

    err = json.Unmarshal([]byte(unparsedJson), &objmap)

    if err != nil {
        fmt.Println(fmt.Sprintf("Invalid JSON provided: %s", unparsedJson))
        os.Exit(1)
    }

    qs := ""

    for k, v := range objmap {
        if qs == "" {
            qs += "?"
        } else {
            qs += "&"
        }
        var val string
        switch t := v.(type) {
        default:
            fmt.Println("unexpected type %T", t)
        case map[string]interface{}:
            j, _ := json.Marshal(v)
            val = fmt.Sprintf("%s", j)
        case string:
            val = fmt.Sprintf("%s", v)
        case []interface{}:
            j, _ := json.Marshal(v)
            val = fmt.Sprintf("%s", j)
        }

        qs += fmt.Sprintf("%s=%s", url.QueryEscape(k), url.QueryEscape(val))
    }

    fmt.Println(qs)
}
