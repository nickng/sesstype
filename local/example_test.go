package local_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

	"go.nickng.io/sesstype"
	"go.nickng.io/sesstype/local"
)

func ExampleParse() {
	rd := strings.NewReader("A ? l().end")
	g, err := local.Parse(rd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g.String())
	// Output:
	// A ?l().end
}

func ExampleParse_diagnosis() {
	// This example shows how to use caret diagnosis to visualise syntax errors.
	//
	var buf bytes.Buffer // Buffer is created for re-displaying for diag
	rd := strings.NewReader("A ? l(.end")
	tee := io.TeeReader(rd, &buf) // Tee reader setup to copy parsed code
	g, err := local.Parse(tee)
	if err != nil {
		if err, ok := err.(*sesstype.ErrParse); ok {
			diag := err.Pos.CaretDiag(buf.Bytes()) // This returns the diagnosis
			fmt.Println(string(diag))
			return
		}
		return
	}
	fmt.Println(g.String())
	// Output:
	// 1:7 → A ? l(.end
	//             ↑
}
