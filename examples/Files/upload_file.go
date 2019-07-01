package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/moltin/gomo"
	"github.com/moltin/gomo/core"
	"github.com/moltin/gomo/form"
)

func main() {
	filename := "example-image.png"

	// Instantiate a new client
	client := gomo.NewClient()

	// Authenticate against the Moltin API
	err := client.Authenticate()
	if err != nil {
		log.Fatal(err)
	}

	// Open the image to upload
	img, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()

	// Define the image to be uploaded, including the file
	// to be read
	file := core.File{
		FileName: filename,
		Public:   true,
		MimeType: "image/png",
		FileSize: 555,
		File: &form.File{
			Name:    filename,
			Content: img,
		},
	}

	// Upload the image
	err = client.Post(
		"files",
		gomo.Form(file),
		gomo.Data(&file),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print the resulting file object
	spew.Dump(file)
}
