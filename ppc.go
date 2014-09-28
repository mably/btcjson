// Copyright (c) 2014-2014 PPCD developers.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcjson

import (
	"fmt"
)

// FloatAmount.
type FloatAmount float64

// MarshalJSON provides a custom Marshal method for FloatAmount.
func (v *FloatAmount) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.6f", *v)), nil
}