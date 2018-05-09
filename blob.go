package backing

import (
	"bytes"
	"io/ioutil"

	"github.com/attic-labs/noms/go/datas"
	nt "github.com/attic-labs/noms/go/types"
	"github.com/pkg/errors"
)

// Blob creates a noms blob from a byte slice
func Blob(db datas.Database, bs []byte) nt.Blob {
	return nt.NewBlob(db, bytes.NewReader(bs))
}

// Unblob creates a byte slice from a noms blob
func Unblob(blob nt.Blob) ([]byte, error) {
	bytes, err := ioutil.ReadAll(blob.Reader())
	return bytes, errors.Wrap(err, "types.Unblob failed")
}
