package main

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// width, height := getImageDimension("rainy.jpg")
	// fmt.Println("Width:", width, "Height:", height)

	resp, err := http.Get("https://i.ytimg.com/vi/f7TcUHYQu4U/hqdefault.jpg")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(resp.Body))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(m)
	// body, err := ioutil.ReadAll(resp.Body)
	// log.Println(body)
}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}
