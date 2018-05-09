package backing

import (
	"encoding/binary"

	"github.com/attic-labs/noms/go/datas"
	nt "github.com/attic-labs/noms/go/types"
)

// Int is an integer which can serialize itself to a noms blob.
//
// Noms doesn't provide a native integer type, so it's on us
// to handle this case. We manage it by providing this type wrapper which
// knows how to serialize itself to and from `[]byte`.
//
// Simple big-endian serialization is used.
type Int int64

// ToBlob converts this Int into a Blob datatype
func (b Int) ToBlob(db datas.Database) nt.Blob {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(b))
	return Blob(db, bytes)
}

// IntFromBlob creates a Int from a blob, if possible
func IntFromBlob(blob nt.Blob) (Int, error) {
	bytes, err := Unblob(blob)
	if err != nil {
		return Int(0), err
	}

	return Int(int64(binary.BigEndian.Uint64(bytes))), nil
}
