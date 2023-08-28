package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	link := "https://open.spotify.com/track/52CBUrIdyf8tbZaUY9iawE"

	u, err := url.Parse(link)	

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("url\t:", link)
	fmt.Println("host\t:", u.Host)
	fmt.Println("scheme\t:", u.Scheme)
	fmt.Println("path\t:", u.Path)
}