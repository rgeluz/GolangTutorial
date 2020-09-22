package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//fmt.Println(resp)

	/*
		//bs := []byte{}
		bs := make([]byte, 99999) //takes a slice and adds number of elements indicated exp 99999 elements. We choose 99999 arbitrarily, assuming this is large enough capacity

		resp.Body.Read(bs) //read is adding data into the bs "byteslice"
		fmt.Println(string(bs))
	*/

	//shorter version
	io.Copy(os.Stdout, resp.Body)

}
