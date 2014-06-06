package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var dozertime = 1 // seconds
var currenthash []byte
var urly = "http://localhost:8080/"

func main() {
	for {
		res, err := http.Get(urly)
		if err != nil {
			log.Fatal(err)
		}
		page, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		hashy := sha1.New()
		hashy.Write(page)
		if !bytes.Equal(currenthash, hashy.Sum(nil)) {
			fmt.Println("\n\n**PANTS ON FIRE!**\n")
		}
		fmt.Printf("%s", string(page))
		currenthash = hashy.Sum(nil)
		time.Sleep(time.Duration(dozertime) * time.Second)
	}

}
