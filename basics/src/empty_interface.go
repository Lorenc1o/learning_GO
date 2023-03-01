package main

import (
	"bytes"
	"fmt"
	"io"
)

// The empty interface is the interface with zero methods.
// It is implemented by all types.
// It is used to represent a value of unknown type.
func main() {
	var myObj interface{} = NewBufferedWriterCloser() // myObj is an interface type with a concrete value of type BufferedWriterCloser
	if wc, ok := myObj.(WriterCloser); ok {           // Type assertion: myObj is an interface type with a concrete value of type WriterCloser
		wc.Write([]byte("Hello Go!"))
		wc.Close()
	}
	r, ok := myObj.(io.Reader) // Type assertion: myObj is an interface type with a concrete value of type io.Reader
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed.")
	}

	// We can do a type switch on an interface type
	switch v := myObj.(type) {
	case io.Reader:
		fmt.Println("Reader", v)
	case io.Writer:
		fmt.Println("Writer", v)
	case WriterCloser:
		fmt.Println("WriterCloser", v)
	default:
		fmt.Println("Unknown type")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
