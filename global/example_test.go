package global_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

	"go.nickng.io/sesstype"
	"go.nickng.io/sesstype/global"
)

func ExampleParse() {
	rd := strings.NewReader("A->B: l().end")
	g, err := global.Parse(rd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g.String())
	// Output:
	// A → B: l().end
}

func ExampleParse_diagnosis() {
	// This example shows how to use caret diagnosis to visualise syntax errors.
	//
	var buf bytes.Buffer // Buffer is created for re-displaying for diag
	rd := strings.NewReader("A->B: l(.end")
	tee := io.TeeReader(rd, &buf) // Tee reader setup to copy parsed code
	g, err := global.Parse(tee)
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
	// 1:9 → A->B: l(.end
	//               ↑
}
