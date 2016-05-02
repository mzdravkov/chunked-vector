package chunkedvec

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewChunkedVecWriter())
	if err != nil {
		panic(err)
	}
}

type ChunkedVecWriter struct{}

func NewChunkedVecWriter() *ChunkedVecWriter {
	return &ChunkedVecWriter{}
}

func (sw *ChunkedVecWriter) Name() string {
	return "ChunkedVec"
}

func (sw *ChunkedVecWriter) Imports(t typewriter.Type) []typewriter.ImportSpec {
	// none
	return []typewriter.ImportSpec{
		typewriter.ImportSpec{Path: "bytes"},
		typewriter.ImportSpec{Path: "container/list"},
		typewriter.ImportSpec{Path: "fmt"},
	}
}

func (cvw *ChunkedVecWriter) Write(w io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(cvw)

	if !found {
		// nothing to be done
		return nil
	}

	license := `// This is an implementation of https://github.com/mzdravkov/chunked-vector
// The MIT License (MIT)
// Copyright (c) 2016 Mihail Zdravkov (mihail0zdravkov@gmail.com)
`

	if _, err := w.Write([]byte(license)); err != nil {
		return err
	}

	tmpl, err := templates.ByTag(t, tag)

	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, t); err != nil {
		return err
	}

	return nil
}
