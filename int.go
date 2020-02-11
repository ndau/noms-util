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
	"fmt"
	"reflect"
	"strconv"

	"github.com/attic-labs/noms/go/marshal"
	nt "github.com/attic-labs/noms/go/types"
)

// Int is an integer which can serialize itself to noms.
//
// Noms doesn't provide a native integer type, so it's on us
// to handle this case. We manage it by providing this type wrapper which
// knows how to serialize itself to and from a base-36 string, which
// turns out to be enormously faster in noms than an array of bytes.
type Int int64

// NomsValue creates a nt.Value from this Int
//
// As this operation cannot fail, we can expose a simpler interface than
// MarshalNoms.
func (i Int) NomsValue() nt.Value {
	return nt.String(strconv.FormatInt(int64(i), 36))
}

// MarshalNoms satisfies the marshal.Marshaler interface
func (i Int) MarshalNoms(vrw nt.ValueReadWriter) (val nt.Value, err error) {
	return i.NomsValue(), nil
}

// static assert that Int satisfies marshal.Marshaler
var _ marshal.Marshaler = (*Int)(nil)

// IntFrom is a bit of sugar to simplify creating an Int from a nt.Value
func IntFrom(v nt.Value) (Int, error) {
	n := Int(0)
	err := n.UnmarshalNoms(v)
	return n, err
}

// UnmarshalNoms satisfies the marshal.Unmarshaler interface
func (i *Int) UnmarshalNoms(v nt.Value) (err error) {
	if i == nil {
		return fmt.Errorf("Int.UnmarshalNoms called on nil int")
	}
	ns, ok := v.(nt.String)
	if !ok {
		return fmt.Errorf(
			"Int.UnmarshalNoms value parameter must be nt.String; was %s",
			reflect.ValueOf(v).Type(),
		)
	}
	i64, err := strconv.ParseInt(string(ns), 36, 64)
	if err != nil {
		return err
	}
	*i = Int(i64)
	return nil
}

// static assert that Int satisfies marshal.Marshaler
var _ marshal.Unmarshaler = (*Int)(nil)
