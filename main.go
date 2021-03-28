package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "sync"
    "time"
)

var httpClient = &http.Client{
   Timeout: time.Second * 10,
}

/**
{
	"userId": 1,
	"id": 1,
	"title": "delectus aut autem",
	"completed": false
}
 */

type User struct {
    UserId int `json:"userId"`
    Id int `json:"id"`
    Title string `json:"title"`
    Completed bool `json:"completed"`
}

func main() {
    var wg sync.WaitGroup

    fmt.Println("Making concurrent api calls")
    for i := 0; i < 2; i++ {
        wg.Add(1)
        go apiCall(i, &wg)
    }

    wg.Wait()
}

func apiCall(i int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Calling api ", i)

    r, err := httpClient.Get("https://jsonplaceholder.typicode.com/todos/1")
    if err != nil {
        log.Fatal(err)
    }
    defer r.Body.Close()

    body, err2 := ioutil.ReadAll(r.Body)
    if err2 != nil {
        panic(err.Error())
    }
    log.Printf("body = %v", string(body))

    foo := new(User)
    err3 := json.Unmarshal(body, &foo)
    if err3 != nil {
        fmt.Println("whoops:", err3)
    }

    log.Printf("s = %v", foo)
    log.Println(foo.Title)
}
