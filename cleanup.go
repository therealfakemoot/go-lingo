package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"regexp"
)

var (
	HEADER_END   = regexp.MustCompile(`^\*\*\* START [\w ]+ \*\*\*$`)
	FOOTER_START = regexp.MustCompile(`End of the Project Gutenberg`)
)

// StripHeaders attempts to remove the Project Gutenberg pre and post amble text.
//
// uses readers to simplify its usage in pipelines, so you aren't constrained to []string or whatever
func StripHeaders(r io.Reader) (io.Reader, error) {
	// the write bool starts off false, because we know we're in the header. when the end of the header is reached, we turn it to true, until we find the start of the footer, and then we finish writing and return
	var write bool
	var b bytes.Buffer
	s := bufio.NewScanner(r)
	for s.Scan() {
		t := s.Text()
		if FOOTER_START.MatchString(t) {
			log.Printf("found footer start")
			return &b, nil
		}
		if write {
			b.WriteString(t + "\n")
			log.Printf("writing content")
		}
		if HEADER_END.MatchString(t) {
			write = true
			log.Printf("found header end")
		}
	}

	return &b, fmt.Errorf("did not find footer")
}
