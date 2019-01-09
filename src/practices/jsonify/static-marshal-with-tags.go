/**
 * FileName:   static-marshal-with-tags.go
 * Author:     Fasion Chan
 * @contact:   fasionchan@gmail.com
 * @version:   $Id$
 *
 * Description:
 *
 * Changelog:
 *
 **/

package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Object interface{}

type Pair struct {
    Name string `json:"name"`
    Value int `json:"value"`
}

type Map struct {
    Bar int `json:"bar"`
    Foo int `json:"value"`
}

type Message struct {
    Type string `json:"type"`
    Data Object `json:"data"`
}

func main() {

    content, err := json.Marshal(Message{
        Type: "map",
        Data: Map{
            Bar: 1,
            Foo: 2,
        },
    })
    if err == nil {
        fmt.Printf("%s\n", content)
    } else {
        log.Fatal(err)
    }

    content, err = json.Marshal(Message{
        Type: "list",
        Data: []Pair{
            Pair{
                Name: "Bar",
                Value: 1,
            },
            Pair{
                Name: "Foo",
                Value: 2,
            },
        },
    })
    if err == nil {
        fmt.Printf("%s\n", content)
    } else {
        log.Fatal(err)
    }
}
