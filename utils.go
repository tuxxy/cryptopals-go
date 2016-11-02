package main

import (
    "net/http"
    "io/ioutil"
)

func GetBody(url string) []byte {
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body_data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    return body_data
}
