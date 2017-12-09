// Command parse-sesstype parses a sesstype file.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"go.nickng.io/sesstype"
	"go.nickng.io/sesstype/global"
	"go.nickng.io/sesstype/local"
)

var (
	outfile   = flag.String("o", "", "Output file (default: stdout)")
	localtype = flag.Bool("local", false, "Parse as local type")

	reader = os.Stdin
	writer = os.Stdout
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: parse-sesstype [options] file.mpst\n\nOptions:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()

	if len(flag.Args()) > 0 {
		var infile string
		if len(flag.Args()) > 0 {
			infile = flag.Arg(0)
		}
		rdFile, err := os.Open(infile)
		if err != nil {
			log.Fatal(err)
		}
		defer rdFile.Close()
		reader = rdFile
	} else {
		fmt.Fprintf(os.Stderr, "Reading from stdin\n")
	}

	if *outfile != "" {
		wrFile, err := os.OpenFile(*outfile, os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer wrFile.Close()
		writer = wrFile
	}

	if *localtype {
		var buf bytes.Buffer
		tee := io.TeeReader(reader, &buf)
		l, err := local.Parse(tee)
		if err != nil {
			if err, ok := err.(*sesstype.ErrParse); ok {
				diag := err.Pos.CaretDiag(buf.Bytes())
				fmt.Fprintf(os.Stderr, "%v:\n%s", err, diag)
			}
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(writer, "%s\n", l.String())
	} else {
		var buf bytes.Buffer
		tee := io.TeeReader(reader, &buf)
		g, err := global.Parse(tee)
		if err != nil {
			if err, ok := err.(*sesstype.ErrParse); ok {
				diag := err.Pos.CaretDiag(buf.Bytes())
				fmt.Fprintf(os.Stderr, "%v:\n%s", err, diag)
			}
			os.Exit(1)
		}
		fmt.Fprintf(writer, "%s\n", g.String())
	}
}
