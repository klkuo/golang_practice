package main

import "github.com/go-martini/martini"

func main() {
    m := martini.Classic()
    m.Get("/", func() string {
        return "Hello 世界!"
    })

    m.RunOnAddr(":8080")
}
