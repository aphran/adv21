package main

import (
    "time"
    "net/http"
    "net/http/cookiejar"
    "fmt"
    "log"
    "os"
    "io"

    "github.com/aphran/adv21/common"
)

func getAdvInput(index int) error {
    var err error
    if index < 1 || index > 24 {
        log.Panic("The index must be between 1 and 25")
    }

    // Initialize cookies
    jar, err := cookiejar.New(nil)
    if err != nil {
        log.Panic(err)
    }
    client := http.Client{Jar: jar}

    cookie := &http.Cookie{
        Name: "session",
        Value: common.ConfigData.SessionCookie,
        MaxAge: 99999,
    }

    //Create HTTP request
    req, err := http.NewRequest("GET", fmt.Sprintf(common.InputURL, index), nil)
    if err != nil {
        log.Panic(err)
    }
    req.AddCookie(cookie)
    resp, err := client.Do(req)
    if err != nil {
        log.Panic(err)
    }
    defer resp.Body.Close()

    //Prepare output directory
    err = os.MkdirAll(common.InputPath, 0755)
    if err != nil {
        log.Fatal(err)
    }
    //Write to file
    inputFile := fmt.Sprintf("%s/%d", common.InputPath, index)
    out, err := os.Create(inputFile)
    if err != nil {
        log.Panic(err)
    }
    _, err = io.Copy(out, resp.Body)
    return err
}

func main() {
    common.LoadConfig()
    fmt.Println(fmt.Sprintf("Using input data path: ./%s/\n", common.InputPath))

    if fmt.Sprintf("%v", time.Now().Month()) == "December" {
        for k := 1; k <= time.Now().Day(); k++ {
            fmt.Println("Getting input data for day", k)
            //getAdvInput(k)
        }
    }
}
