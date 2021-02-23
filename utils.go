package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
)

type Resp struct {
    Encoded []byte `json:"encoded"`
}

// getJSON fetches the contents of the given URL
// and decodes it as JSON into the given result,
// which should be a pointer to the expected data.
func GetJSON(url string) ([]byte, error) {
    resp, err := http.Get(url)
    var sszResp Resp
    if err != nil {
        return sszResp.Encoded, fmt.Errorf("cannot fetch URL %q: %v", url, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return sszResp.Encoded, fmt.Errorf("unexpected http GET status: %s", resp.Status)
    }
    bodybytes, err := ioutil.ReadAll(resp.Body)
    fmt.Println("bodybytes:", bodybytes)
    if err != nil {
        return sszResp.Encoded, fmt.Errorf("ERROR reading response:", err)
    }
    err = json.Unmarshal(bodybytes, &sszResp)
    if err != nil {
        return sszResp.Encoded, fmt.Errorf("ERROR Unmarshalling reply from prysm:", err)
    }
    fmt.Println("SSZ Encoded:", sszResp.Encoded)
    return sszResp.Encoded, nil
}

