package save

import (
	"encoding/binary"
	"io"
	"strings"
)

// Makes it clearer that a field is just padding.
type padding uint8

// Booleans are single byte values in saves, and since save files
//  are dumps from C++ structures, they are padded, and thus bools
//  are often followed by three bytes of padding to make the size
//  aligned to a four-byte boundary.
type boolPadding [3]uint8

type vector3 struct {
	X, Y, Z float32
}

func mustRead(file io.Reader, data interface{}) {
	err := binary.Read(file, binary.LittleEndian, data)

	if err != nil {
		panic(err)
	}
}

func mustWrite(file io.Writer, data interface{}) {
	err := binary.Write(file, binary.LittleEndian, data)

	if err != nil {
		panic(err)
	}
}

func nullTerminate(str *string) {
	index := strings.IndexRune(*str, '\x00')
	*str = (*str)[:index]
}