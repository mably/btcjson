// Copyright (c) 2014-2014 PPCD developers.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcjson

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// FloatAmount specific type with custom marshalling
type FloatAmount float64

// MarshalJSON provides a custom Marshal method for FloatAmount.
func (v *FloatAmount) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.6f", *v)), nil
}

// GetDifficultyResult models the data of getdifficulty command.
type GetDifficultyResult struct {
	ProofOfWork    float64 `json:"proof-of-work"`
	ProofOfStake   float64 `json:"proof-of-stake"`
	SearchInterval int32   `json:"search-interval"`
}

// GetKernelStakeModifierCmd is a type handling custom marshaling and
// unmarshaling of getkernelstakemodifier JSON RPC commands.
type GetKernelStakeModifierCmd struct {
	id      interface{}
	Hash    string
	Verbose bool
}

// Enforce that GetKernelStakeModifierCmd satisifies the Cmd interface.
var _ Cmd = &GetKernelStakeModifierCmd{}

// NewGetKernelStakeModifierCmd creates a new GetKernelStakeModifierCmd.
func NewGetKernelStakeModifierCmd(id interface{}, hash string, optArgs ...bool) (*GetKernelStakeModifierCmd, error) {
	// default verbose is set to true to match old behavior
	verbose, verboseTx := true, false

	optArgsLen := len(optArgs)
	if optArgsLen > 0 {
		if optArgsLen > 2 {
			return nil, ErrTooManyOptArgs
		}
		verbose = optArgs[0]
		if optArgsLen > 1 {
			verboseTx = optArgs[1]

			if !verbose && verboseTx {
				return nil, ErrInvalidParams
			}
		}
	}

	return &GetKernelStakeModifierCmd{
		id:      id,
		Hash:    hash,
		Verbose: verbose,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (cmd *GetKernelStakeModifierCmd) Id() interface{} {
	return cmd.id
}

// Method satisfies the Cmd interface by returning the json method.
func (cmd *GetKernelStakeModifierCmd) Method() string {
	return "getkernelstakemodifier"
}

// MarshalJSON returns the JSON encoding of cmd.  Part of the Cmd interface.
func (cmd *GetKernelStakeModifierCmd) MarshalJSON() ([]byte, error) {
	params := make([]interface{}, 1, 3)
	params[0] = cmd.Hash
	if !cmd.Verbose {
		// set optional verbose argument to false
		params = append(params, false)
	}
	// Fill and marshal a RawCmd.
	raw, err := NewRawCmd(cmd.id, cmd.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd.  Part of
// the Cmd interface.
func (cmd *GetKernelStakeModifierCmd) UnmarshalJSON(b []byte) error {
	// Unmashal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) > 2 || len(r.Params) < 1 {
		return ErrWrongNumberOfParams
	}

	var hash string
	if err := json.Unmarshal(r.Params[0], &hash); err != nil {
		return fmt.Errorf("parameter 'hash' must be a string: %v", err)
	}

	optArgs := make([]bool, 0, 2)
	if len(r.Params) > 1 {
		var verbose bool
		if err := json.Unmarshal(r.Params[1], &verbose); err != nil {
			return fmt.Errorf("second optional parameter 'verbose' must be a bool: %v", err)
		}
		optArgs = append(optArgs, verbose)
	}

	newCmd, err := NewGetKernelStakeModifierCmd(r.Id, hash, optArgs...)
	if err != nil {
		return err
	}

	*cmd = *newCmd
	return nil
}

// StakeModifier specific type with custom marshalling
type StakeModifier uint64

// MarshalJSON provides a custom Marshal method for StakeModifier.
func (v StakeModifier) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%d\"", uint64(v))), nil
}

// UnmarshalJSON provides a custom Unmarshal method for StakeModifier.
func (v *StakeModifier) UnmarshalJSON(b []byte) (err error) {
	var s string
	json.Unmarshal(b, &s)
	var u uint64
	u, err = strconv.ParseUint(s, 0, 64)
	*v = StakeModifier(u)
	return
}

// KernelStakeModifierResult models the data from the getkernelstakemodifier
// command when the verbose flag is set.  When the verbose flag is not set,
// getkernelstakemodifier return a hex-encoded string.
type KernelStakeModifierResult struct {
	Hash                string        `json:"hash"`
	KernelStakeModifier StakeModifier `json:"kernelstakemodifier"`
}
