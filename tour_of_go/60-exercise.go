// Exercise: Rot13 Reader
//
// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
//
// For example, the gzip.NewReader function takes an io.Reader (a stream of gzipped data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).
//
// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the ROT13 substitution cipher to all alphabetical characters.
//
// The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (int, error) {
	derot13 := func(c byte) byte {
		var v byte
		if (c >= 'A' && c < 'N') || (c >= 'a' && c < 'n') {
			v = c + 13
		} else if (c > 'M' && c <= 'Z') || (c > 'm' && c <= 'z') {
			v = c - 13
		}
		return v
	}

	_p := make([]byte, 100)

	if n, err := rot.r.Read(_p); err != nil {
		return n, err
	}

	for i, v := range _p {
		p[i] = derot13(byte(v))
	}

	return len(_p), nil

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
