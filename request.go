// Make a request to youtube API for tricking videos
package main

import (
    "fmt"
    "log"
    "encoding/json"
    "net/http"
)

var base_url = "https://www.googleapis.com/youtube/v3"
var query = "/search?part=snippet&q=Martial+Arts+Tricking"
var api_key = "&key={YOUR API KEY}"

type Video struct {
    Id struct {
        Kind string `json:"kind"`
        VideoId string `json:"videoId"`
        } `json:"id"`
    Snippet struct {
        Title string `json:"title"`
        } `json:"snippet"`
}

type Videos struct {
    Items []Video `json:"items"`
}

func main() {
    url := base_url + query + api_key
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    var vs Videos
    err = json.NewDecoder(resp.Body).Decode(&vs)
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(vs.Items)
}
