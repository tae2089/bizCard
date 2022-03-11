// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bizCard/ent/bizcard"
	"bizCard/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bizcardFields := schema.BizCard{}.Fields()
	_ = bizcardFields
	// bizcardDescAge is the schema descriptor for age field.
	bizcardDescAge := bizcardFields[3].Descriptor()
	// bizcard.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	bizcard.AgeValidator = bizcardDescAge.Validators[0].(func(int) error)
}
