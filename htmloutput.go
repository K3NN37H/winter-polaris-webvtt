package main

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"os"
)

const HtmlOutputName = "winter-polaris.html"

//go:embed template.html
var html string

// Generate HTML file with embedded vtt file
func Generate(trackData [][]byte) {
	t, _ := template.New("web").Parse(html)

	out, _ := os.Create(HtmlOutputName)
	defer out.Close()

	encoded := make([]string, len(trackData))

	for i, track := range trackData {
		encoded[i] = base64.StdEncoding.EncodeToString(track)
		if DEBUG {
			fmt.Println(encoded[i])
		}
	}

	data := struct {
		TrackSrc []string
	}{
		TrackSrc: encoded,
	}

	t.Execute(out, data)
}
