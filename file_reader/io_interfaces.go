package filereader

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	// "io/ioutil"
)

// Defines the model struct
type HashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

type Hash_Reader_Interface interface{
	io.Reader
	Hash() string
}

func NewHashReader(b []byte) *HashReader {
	return &HashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *HashReader) Hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func Hash_And_BroadCast(r Hash_Reader_Interface) error {

	// object, err := io.ReadAll(r)
	// if err != nil {
	// 	return err
	// }

	// hash := sha1.Sum(object)
	// fmt.Println(hex.EncodeToString(hash[:]))

	hash := r.Hash()
	fmt.Println(hash)

	return BroadCast(r)
}

func BroadCast(r io.Reader) error {

	object, err := ioutil.ReadAll(r) 
	if err != nil {
		return err
	}

	fmt.Println("Strings of the bytes: ", string(object))

	return nil
}
