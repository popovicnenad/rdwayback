package main

import (
    "bufio"
    "encoding/json"
    "flag"
    "fmt"
    "net/http"
    "net/url"
    "os"
)

func extractSubdomain(s string) string {
    u, err := url.Parse(s)
    if err != nil {
        return ""
    }
    return u.Host
}

func readWayback(domain string, filterSubdomains bool, outputFile string) {
    urlTemplate := "http://web.archive.org/cdx/search/cdx?url=*.%s/*&output=json&fl=original&collapse=urlkey"
    waybackUrl := fmt.Sprintf(urlTemplate, domain)

    client := &http.Client{}
    request, reqErr := http.NewRequest("GET", waybackUrl, nil)
    if reqErr != nil {
        panic(reqErr)
    }
    response, err := client.Do(request)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()

     if response.StatusCode != 200 {
        fmt.Print("Received HTTP", response.StatusCode)
        os.Exit(1)
    }

    var urlList [][]string
    decodeErr := json.NewDecoder(response.Body).Decode(&urlList)
    if decodeErr != nil {
        panic(decodeErr)
    }
    if len(urlList) == 0 {
        os.Exit(0)
    }

    foundUrls := make(map[string]bool)
    for _, ul := range urlList[1:] {
        for _, u := range ul {
            if u == "" {
                continue
            }
            if filterSubdomains {
                subdomain := extractSubdomain(u)
                if subdomain != "" {
                    foundUrls[subdomain] = true
                }
            } else {
                foundUrls[u] = true
            }
        }
    }

    if outputFile != "" {
        file, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
        if err != nil {
            panic(err)
        }
        defer file.Close()
        writer := bufio.NewWriter(file)
        for foundUrl, _ := range foundUrls {
            _, _ = fmt.Fprintln(writer, foundUrl)
            _ = writer.Flush()
            fmt.Println(foundUrl)
        }
    } else {
        for foundUrl, _ := range foundUrls {
            fmt.Println(foundUrl)
        }
    }
}


func main(){
    var domain = flag.String("d", "", "Input domain")
    var filterSubdomains = flag.Bool("sub", false, "Get list of subdomains")
    var outputFile = flag.String("o", "", "Output file")
    flag.Parse()

    if *domain == "" {
        fmt.Println("Error: Input domain is required")
        flag.Usage()
        os.Exit(1)
    }
    readWayback(*domain, *filterSubdomains, *outputFile)
}