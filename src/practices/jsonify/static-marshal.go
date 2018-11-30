/**
 * FileName:   static-marshal.go
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
    Name string
    Value int
}

type Map struct {
    Bar int
    Foo int
}

type Message struct {
    Type string
    Data Object
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
