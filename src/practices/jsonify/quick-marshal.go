/**
 * FileName:   quick-marshal.go
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

type Pair struct {
    Name string
    Value int
}

func main() {
    pair := Pair{
        Name: "bar",
        Value: 1,
    }

    content, err := json.Marshal(pair)
    if err == nil {
        fmt.Printf("%s\n", content)
    } else {
        log.Fatal(pair)
    }
}
