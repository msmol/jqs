package main

import (
    "os"
    "fmt"
    "encoding/json"
    "net/url"
)

func main() {

    if len(os.Args) != 2 {
        fmt.Println("Must provide exactly one argument")
        os.Exit(1)
    }
    unParsedJson := os.Args[1]

    var objmap map[string]interface{}

    err := json.Unmarshal([]byte(unParsedJson), &objmap)

    if err != nil {
        fmt.Println(fmt.Sprintf("Invalid JSON provided: %s", unParsedJson))
        os.Exit(1)
    }

    qs := ""

    for k, v := range objmap {
        val := fmt.Sprintf("%v", v)
        qs += fmt.Sprintf("%s=%s", url.QueryEscape(k), url.QueryEscape(val))
    }

    fmt.Println(qs)
}

