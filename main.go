package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var enc = japanese.ShiftJIS

var DEBUG = false

const ScriptFilename = "C97_drama_CD_Translated_Script.txt"

func main() {
	flag.BoolVar(&DEBUG, "debug", false, "Turn on debug output")
	flag.Parse()

	// open and read entire script file
	script, err := readLines(ScriptFilename)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Generate buffer with track 1 data
	track1 := generateTrackVtt(Track1, script)
	track2 := generateTrackVtt(Track2, script)

	Generate([][]byte{track1.Bytes(), track2.Bytes()})
}

// Generate WebVTT for the given track in a Buffer
func generateTrackVtt(lineInfo []LineTimestamp, script []string) bytes.Buffer {
	var output bytes.Buffer

	fmt.Fprintln(&output, "WEBVTT")
	fmt.Fprintln(&output, "")

	for _, line := range lineInfo {
		lineNum := line.line
		timestamp := line.timestamp
		count := 0
		fmt.Fprintln(&output, timestamp, " line:0")
		for len(script[lineNum-1+count]) > 0 {
			if DEBUG {
				fmt.Println(lineNum, script[lineNum-1+count], timestamp)
			}
			fmt.Fprintln(&output, script[lineNum-1+count])
			count++
		}
		fmt.Fprintln(&output, "")
	}

	return output
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := transform.NewReader(file, enc.NewDecoder())

	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
