// Make a request to youtube API for tricking videos
package main

import (
    "fmt"
    "log"
    "encoding/json"
    "net/http"
    "io/ioutil"
)
var base_url = "https://www.googleapis.com/youtube/v3"
var query = "/search?part=snippet&q=Martial+Arts+Tricking"
var api_key = "&key="

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
    // read in API key
    key, error_ := ioutil.ReadFile("user_data.txt")
    if error_ != nil {
        log.Fatal(error_)
    }
    api_key += string(key)

    // buil url
    url := base_url + query + api_key

    // make request
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    body, error_g := ioutil.ReadAll(resp.Body)
    if error_g != nil {
        fmt.Println(error_g)
    }
    fmt.Println(resp.Status)
    fmt.Println(resp.Request)
    fmt.Println(string(body))
    defer resp.Body.Close()
    // parse response
    var vs Videos
    err = json.NewDecoder(resp.Body).Decode(&vs)
    if err != nil {
        fmt.Printf("%T\n", err)
        // fmt.Println("Error:", err)
        log.Fatal(err)
    }
    for _, v := range vs.Items {
        fmt.Println(v.Id.VideoId, v.Snippet.Title)
    }
}
