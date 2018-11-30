/**
 * FileName:   free-marshal.go
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

type Object interface {}
type Array []Object

type JsonObject map[string]Object
type JsonArray Array

func main() {
    content, err := json.Marshal(JsonObject{
        "type": "map",
        "data": JsonObject{
            "bar": 1,
            "foo": 2,
        },
    })
    if err == nil {
        fmt.Printf("%s\n", content)
    } else {
        log.Fatal(err)
    }

    content, err = json.Marshal(JsonObject{
        "type": "list",
        "data": Array{
            JsonObject{
                "name": "bar",
                "value": 1,
            },
            JsonObject{
                "name": "foo",
                "value": 2,
            },
        },
    })
    if err == nil {
        fmt.Printf("%s\n", content)
    } else {
        log.Fatal(err)
    }
}
