package main

import (
    "os"
    "fmt"
    "encoding/json"
    "net/url"
)

func main() {

    unParsedJson := os.Args[1]

    var objmap map[string]interface{}

    err := json.Unmarshal([]byte(unParsedJson), &objmap)

    if err != nil {
        panic(err)
    }

    qs := ""

    for k, v := range objmap {
        val := fmt.Sprintf("%v", v)
        qs += fmt.Sprintf("%s=%s", url.QueryEscape(k), url.QueryEscape(val))
    }

    fmt.Println(qs)
}

