package main

import (
	"fmt"
	"net/http"
	"time"
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

	//fmt.Println(<-c) //receive a value from this channel and then log it. This is a blocking call, because it is waiting for a value

	/*fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println(<-c)*/ //This one will hang, because its waiting for antoher message, however there was only five go routines initiatied and five <-c received by the main go routine and printed

	/* for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	} */

	/*
		for { // always true
			go checkLink(<-c, c)
		} */

	//alternative loop syntax from above. l is link, and c is channel
	for l := range c {
		//go checkLink(l, c)
		go func(link string) { //functional literal is equivalent to a anonymous function or lambda function
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

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
	_, err := http.Get(link) //This is a blocking call, because it is waiting for a value
	if err != nil {
		fmt.Println(link, "might be down!")
		//c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	//c <- "Yep its up"

	c <- link
}
