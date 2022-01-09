package main

import (
	"flag"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

func main() {
	src := flag.String("src", "/home/daniel/Offtopic/Images/mobile pix/", "Directory where images to be resized are")
	dst := flag.String("name", "/tmp/resized", "Directory where resized images will go")
	flag.Parse()

	if *src == *dst { //I don't want images to go to the same place as to avoid the risk of overwriting them
		log.Fatalln("Source and destination directories cannot be the same")
	}

	_, err := os.Stat(*src)
	if err != nil {
		log.Fatalln("Can't os.Stat() source directory:", err)
	}

	dstDir, err := os.Stat(*dst)
	if os.IsNotExist(err) {
		mkdirErr := os.Mkdir(*dst, 0755)
		if mkdirErr != nil {
			log.Fatalln("Unable to create destination directory:", mkdirErr)
		}
	}
	if err == nil {
		if dstDir.IsDir() {
			log.Println("Destination directory is a directory and already exists. Neat!")
		} else {
			log.Fatalln("Destination directory is a file, exiting")
		}
	}

	srcFiles, err := ioutil.ReadDir(*src)
	if err != nil {
		log.Fatalln("Unable to read contents of source directory:", err)
	}

	ch := make(chan string)

	go func() {
		for _, srcFile := range srcFiles {
			if !strings.HasSuffix(srcFile.Name(), ".jpg") {
				continue //if not an image, don't do anything
			}
			ch <- srcFile.Name()
		}
		close(ch)
	}()

	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for imgName := range ch {
				log.Println(imgName)
				img, err := imgio.Open(*src + string(os.PathSeparator) + imgName)
				if err != nil {
					log.Println("Couldn't open", imgName, "as image:", err)
					continue
				}
				resized := transform.Resize(img, img.Bounds().Dx()/2, img.Bounds().Dy()/2, transform.Linear)
				err = imgio.Save(*dst+string(os.PathSeparator)+imgName, resized, imgio.JPEGEncoder(82))
				if err != nil {
					log.Fatalln("Unable to save image:", err)
				}

			}
		}()
	}
	wg.Wait()
	log.Println("Program finished")
}
