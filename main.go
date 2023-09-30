package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func main() {
	// Read contents from file
	bytes, err := os.ReadFile("windows-1252.txt")
	if err != nil {
		log.Fatalln(err)
	}
	content := string(bytes)
	fmt.Println(string(content))

	// Encode content to utf-8
	utf8Content := toUtf8(content)
	fmt.Println(utf8Content)
}

func toUtf8(str string) string {
	r := transform.NewReader(bytes.NewReader([]byte(str)), charmap.ISO8859_1.NewDecoder())
	decoded, _ := io.ReadAll(r)
	return string(decoded)
}
