package backing

// ----- ---- --- -- -
// Copyright 2020 The Axiom Foundation. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----

import (
	"bytes"
	"io/ioutil"

	nt "github.com/ndau/noms/go/types"
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
