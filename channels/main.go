package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// launch new child go routine, make separate channel to execute the function
		// When the function makes a blocking call, execute the next code at main routine
		// Different scenario when using more than one cpu core
		//
		// An important bug, when creating child routine, if the main routine has nothing else to do
		// The program quit early despite child routine is still running, this can be avoid by using channel
		go checkLink(link, c)
	}

	// Sleep the routine until the channel receive a new message
	// fmt.Println(<- c)

	// Infinite loop ( with better and clear syntax )
	for l := range c {

		//l := l

		// Create a go routine that executes A Function literal ( or lambda / anonymous function )
		// At a parameter, create a copy as a new variable into the function
		// that didnt changed during sleep when using channel.
		go func(link string) {
			time.Sleep(time.Second * 5)

			// That wait for new message in a channel
			// Then execute a checkLink function with channel message and channel as an argument
			checkLink(link, c)
		}(l)

	}


}
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
