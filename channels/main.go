package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//This calls checkLink in a serial/sequential manner
	/*for _, link := range links {
		checkLink(link)
	} */

	//added go keyword. When running this, the main function ended before finishing each of the child go routines. Need to use channels
	/*for _, link := range links {
		go checkLink(link)
	}*/

	//need to pass channel to checklink
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c) //receive a value from this channel and then log it
}

//this function will be used as a goroutine
/* func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
} */

//this function uses a channel
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep its up"
}
