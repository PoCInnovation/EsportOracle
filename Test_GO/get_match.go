package main

import(
    "net/http"
    "fmt"
    "io/ioutil"
    "github.com/joho/godotenv"
    "os"
    "log"
)

var TOKEN string

func getURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
    // now you can use os.Getenv ...  
    TOKEN = os.Getenv("TOKEN")
    return "https://api.pandascore.co/videogames"
}

func main() {
    url := getURI()
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    // Set the request headers
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", TOKEN))

    // Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Response Status:", resp.Status)
    fmt.Println("Response Headers:", resp.Header)
    fmt.Println("Response Body:", string(body))
}