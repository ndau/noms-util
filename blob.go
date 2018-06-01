package backing

import (
	"bytes"
	"io/ioutil"

	nt "github.com/attic-labs/noms/go/types"
	"github.com/pkg/errors"
)

// Blob creates a noms blob from a byte slice
//
// Note: any datas.Database satisfies nt.ValueReadWriter
func Blob(vrw nt.ValueReadWriter, bs []byte) nt.Blob {
	return nt.NewBlob(vrw, bytes.NewReader(bs))
}

// Unblob creates a byte slice from a noms blob
func Unblob(blob nt.Blob) ([]byte, error) {
	bytes, err := ioutil.ReadAll(blob.Reader())
	return bytes, errors.Wrap(err, "Unblob failed")
}
