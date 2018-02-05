// -*- go -*-
package main

import (
	"encoding/json"
	"time"
	"encoding/hex"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"math/rand"
)

func getCurrentDifficulty() float64 {
	data := make(map[string]interface{})
        var getDifficulty = &http.Client{Timeout: 10 * time.Second}
        res, err := getDifficulty.Get("https://blockexplorer.com/api/status?q=getDifficulty")
        if err != nil {
                panic(err)
        }
        defer res.Body.Close()
	difficulty, err := ioutil.ReadAll(res.Body)
        json.Unmarshal(difficulty, &data)
	return data["difficulty"].(float64)
}

var currentDifficulty = getCurrentDifficulty()
var dummyBlockData = make([]byte, 176)

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", filePostHasher)
	addr := ":" + os.Getenv("PORT")
	fmt.Printf("Listening on %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func filePostHasher(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "GET":
			rand.Read(dummyBlockData)
			hash := sha256.New()
			hash.Write(dummyBlockData)
			//fmt.Fprintln(w, float64(hash.Sum(nil)))
			fmt.Fprintln(w, hex.EncodeToString(hash.Sum(nil)))
		case "POST":
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				fmt.Println("error reading body")
				panic(err)
			}
			hash := sha256.New()
			hash.Write(body)
			fmt.Fprintln(w, hex.EncodeToString(hash.Sum(nil)))
	}
}
