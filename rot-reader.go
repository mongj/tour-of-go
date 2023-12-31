package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
    isCapital := x >= 'A' && x <= 'Z'
    if !isCapital && (x < 'a' || x > 'z') {
        return x
    }

    x += 13
    if isCapital && x > 'Z' || !isCapital && x > 'z' {
        x -= 26
    }
    return x
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
    for i := 0; i <= n; i++ {
        b[i] = rot13(b[i])
    }
    return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
