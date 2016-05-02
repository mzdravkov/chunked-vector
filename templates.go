package chunkedvec

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	chunkedVec,
}

var chunkedVec = &typewriter.Template{
	Name: "ChunkedVec",
	Text: `
import (
	"bytes"
	"container/list"
	"fmt"
)

type {{.Name}}ChunkedVec struct {
	list      *list.List
	ChunkSize uint
	Empty     {{.Pointer}}{{.Name}}
}

// Creates a new {{.Name}}ChunkedVec with chunkSize as provided
func New{{.Name}}ChunkedVec(chunkSize uint) *{{.Name}}ChunkedVec {
	if chunkSize == 0 {
		chunkSize = 1024
	}

	return &{{.Name}}ChunkedVec{
		list:      list.New(),
		ChunkSize: chunkSize,
	}
}

// Adds the element to the ChunkedVec and returns the position it was added to
func (cv *{{.Name}}ChunkedVec) Add(element {{.Pointer}}{{.Name}}) (uint, uint) {
	listIndex := 0
	for e := cv.list.Front(); e != nil; e = e.Next() {
		for index, value := range e.Value.([]{{.Pointer}}{{.Name}}) {
			if value == cv.Empty {
				e.Value.([]{{.Pointer}}{{.Name}})[index] = element
				return uint(listIndex), uint(index)
			}
		}

		listIndex++
	}

	slice := make([]{{.Pointer}}{{.Name}}, cv.ChunkSize)
	slice[0] = element
	cv.list.PushBack(slice)

	return uint(listIndex), uint(0)
}

// Overwrites the given position to hold the given value
func (cv *{{.Name}}ChunkedVec) PutAt(element {{.Pointer}}{{.Name}}, listIndex, sliceIndex uint) {
	var i uint = 0
	e := cv.list.Front()
	for ; i < listIndex; e = e.Next() {
		i++
	}

	e.Value.([]{{.Pointer}}{{.Name}})[sliceIndex] = element
}

// Puts the cv.Empty value at the given position
func (cv *{{.Name}}ChunkedVec) DeleteAt(listIndex, sliceIndex uint) {
	cv.PutAt(cv.Empty, listIndex, sliceIndex)
}

// Returns the value that is on the given position
func (cv *{{.Name}}ChunkedVec) Get(listIndex, sliceIndex uint) {{.Pointer}}{{.Name}} {
	var i uint = 0
	e := cv.list.Front()
	for ; i < listIndex; e = e.Next() {
		i++
	}

	return e.Value.([]{{.Pointer}}{{.Name}})[sliceIndex]
}

// Remove list nodes that has arrays that are with the Empty element only
func (cv *{{.Name}}ChunkedVec) Shrink() {
	for e := cv.list.Front(); e != nil; e = e.Next() {
		allEmpty := true
		for _, value := range e.Value.([]{{.Pointer}}{{.Name}}) {
			if value != cv.Empty {
				allEmpty = false
				break
			}
		}

		if allEmpty {
			cv.list.Remove(e)
		}
	}
}

// Returns the number of non-empty valued elements
func (cv *{{.Name}}ChunkedVec) Len() int {
	number := 0

	for e := cv.list.Front(); e != nil; e = e.Next() {
		for _, value := range e.Value.([]{{.Pointer}}{{.Name}}) {
			if value != cv.Empty {
				number++
			}
		}
	}

	return number
}

// Returns the current capacity of the {{.Name}}ChunkedVec
// i.e. the number of elements it can currently hold without growing
func (cv *{{.Name}}ChunkedVec) Cap() int {
	return cv.list.Len() * int(cv.ChunkSize)
}

// Iter returns a channel of type {{.Pointer}}{{.Name}} that you can range over.
func (cv *{{.Name}}ChunkedVec) Iter() <-chan {{.Pointer}}{{.Name}} {
	ch := make(chan {{.Pointer}}{{.Name}})

	go func() {
		for e := cv.list.Front(); e != nil; e = e.Next() {
			for _, value := range e.Value.([]{{.Pointer}}{{.Name}}) {
				ch <- value
			}
			close(ch)
		}
	}()

	return ch
}

// Checks if the {{.Name}}ChunkedVec contains the given element
func (cv *{{.Name}}ChunkedVec) Contains(element {{.Pointer}}{{.Name}}) bool {
	for e := cv.list.Front(); e != nil; e = e.Next() {
		for _, value := range e.Value.([]{{.Pointer}}{{.Name}}) {
			if value == element {
				return true
			}
		}
	}

	return false
}

// Checks if the {{.Name}}ChunkedVec contains all of the given element
func (cv *{{.Name}}ChunkedVec) ContainsAll(searchingFor ...{{.Pointer}}{{.Name}}) bool {
	for _, s := range searchingFor {
		if !cv.Contains(s) {
			return false
		}
	}

	return true
}

// Checks if this {{.Name}}ChunkedVec is equal to another one
// Two {{.Name}}ChunkedVecs are equal if they have the same number of lists
// with slices that have the same values
func (cv *{{.Name}}ChunkedVec) Equal(other *{{.Name}}ChunkedVec) bool {
	// no worries, the complexity of this is O(1)
	if cv.list.Len() != other.list.Len() {
		return false
	}

	e2 := other.list.Front()
	for e1 := cv.list.Front(); e1 != nil; e1 = e1.Next() {
		len1 := len(e1.Value.([]{{.Pointer}}{{.Name}}))
		len2 := len(e2.Value.([]{{.Pointer}}{{.Name}}))
		if len1 != len2 {
			return false
		}

		for i := 0; i < len1; i++ {
			if e1.Value.([]{{.Pointer}}{{.Name}})[i] != e2.Value.([]string)[i] {
				return false
			}
		}

		e2 = e2.Next()
	}

	return true
}

// Clone returns a clone of the {{.Name}}ChunkedVec.
// Does NOT clone the underlying elements.
func (cv *{{.Name}}ChunkedVec) Clone() *{{.Name}}ChunkedVec {
	cloned{{.Name}}ChunkedVec := New{{.Name}}ChunkedVec(cv.ChunkSize)

	var listIndex uint = 0
	for e := cv.list.Front(); e != nil; e = e.Next() {
		for index, value := range e.Value.([]{{.Pointer}}{{.Name}}) {
			cloned{{.Name}}ChunkedVec.PutAt(value, listIndex, uint(index))
		}

		listIndex++
	}

	return cloned{{.Name}}ChunkedVec
}

// Clears all the data in the {{.Name}}ChunkedVec
func (cv *{{.Name}}ChunkedVec) Clear() {
	for e := cv.list.Front(); e != nil; e = e.Next() {
		cv.list.Remove(e)
	}
}

func (cv *{{.Name}}ChunkedVec) String() string {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "{\n")
	for e := cv.list.Front(); e != nil; e = e.Next() {
		slice := e.Value.([]{{.Pointer}}{{.Name}})
		if _, err := fmt.Fprintf(&buff, fmt.Sprintf("\t%s\n", slice)); err != nil {
			panic("Can't write to buffer")
		}
	}
	fmt.Fprintf(&buff, "\n}")

	return buff.{{.Name}}()
}
`,
	TypeConstraint: typewriter.Constraint{Comparable: true},
}
