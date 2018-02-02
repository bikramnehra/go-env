// -*- go -*-
package main

import (
	"encoding/hex"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", filePostHasher)
	addr := ":" + os.Getenv("PORT")
	fmt.Printf("Listening on %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func filePostHasher(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "GET":
			fmt.Printf("\n%+v\n\n", req)
			fmt.Fprintf(w,"\n%+v\n\n", req)
			fmt.Fprintln(w, strings.Join(os.Environ(), "\n"))
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
