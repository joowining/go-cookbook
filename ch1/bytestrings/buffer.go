package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

func Buffer(rawString string) *bytes.Buffer {
	// 원시 바이트
	rawBytes := []byte(rawString)
	// 버퍼 생성 
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	// 버퍼 생성 2
	b = bytes.NewBuffer(rawBytes)

	// 버퍼 생성 3 
	// new() and bytes.NewBuffer() and bytes.NewBufferString()
	// return same result
	b = bytes.NewBufferString(rawString)
	return b
}

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
