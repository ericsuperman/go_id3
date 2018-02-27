package main

import (
    "os"
	"fmt"
	"log"

	"github.com/bogem/id3v2"
)

func main() {

    // argsWithProg := os.Args
    // argsWithoutProg := os.Args[1:]

	if(len(os.Args) != 2) {
 		log.Fatal("Error - no input file.")
	}
	
    arg := os.Args[1]

	// Open file and parse tag in it.
	tag, err := id3v2.Open(arg, id3v2.Options{Parse: true})
	if err != nil {
 		log.Fatal("Error while opening mp3 file: ", err)
 	}
	defer tag.Close()

	// Read frames.
	fmt.Println(tag.Artist())
	fmt.Println(tag.Title())

	// Set simple text frames.
	tag.SetArtist("New artist")
	tag.SetTitle("New title")

	// Set comment frame.
	comment := id3v2.CommentFrame{
		Encoding:    id3v2.EncodingUTF8,
		Language:    "eng",
		Description: "My opinion",
		Text:        "Very good song",
	}
	tag.AddCommentFrame(comment)

	// Write it to file.
	if err = tag.Save(); err != nil {
		log.Fatal("Error while saving a tag: ", err)
	}
}