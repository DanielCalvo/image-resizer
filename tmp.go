package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello world!")

	src := flag.String("src", "", "Directory where images to be resized are")
	dst := flag.String("name", "", "Directory where resized images go")
	flag.Parse()

	if *src == *dst { //I have to verify if you can have src and dst be the same or if you have an infinite loop
		log.Fatalln("src and dst cannot be the same")
	}

	//verify if src dir exists
	//verify if dst dir exists

	//have a function to resize one image, you're doing too much in those functions

}
