package main

import "net/http"

func main() {
        panic(http.ListenAndServe(":8448", http.FileServer(http.Dir("."))))
}

