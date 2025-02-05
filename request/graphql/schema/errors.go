// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package schema

import "github.com/sourcenetwork/defradb/errors"

const (
	errDuplicateField             string = "duplicate field"
	errFieldMissingRelation       string = "field missing associated relation"
	errRelationMissingField       string = "relation missing field"
	errAggregateTargetNotFound    string = "aggregate target not found"
	errSchemaTypeAlreadyExist     string = "schema type already exists"
	errObjectNotFoundDuringThunk  string = "object not found whilst executing fields thunk"
	errTypeNotFound               string = "no type found for given name"
	errRelationNotFound           string = "no relation found"
	errNonNullForTypeNotSupported string = "NonNull variants for type are not supported"
)

var (
	ErrDuplicateField             = errors.New(errDuplicateField)
	ErrFieldMissingRelation       = errors.New(errFieldMissingRelation)
	ErrRelationMissingField       = errors.New(errRelationMissingField)
	ErrAggregateTargetNotFound    = errors.New(errAggregateTargetNotFound)
	ErrSchemaTypeAlreadyExist     = errors.New(errSchemaTypeAlreadyExist)
	ErrObjectNotFoundDuringThunk  = errors.New(errObjectNotFoundDuringThunk)
	ErrTypeNotFound               = errors.New(errTypeNotFound)
	ErrRelationNotFound           = errors.New(errRelationNotFound)
	ErrNonNullForTypeNotSupported = errors.New(errNonNullForTypeNotSupported)
	ErrRelationMutlipleTypes      = errors.New("relation type can only be either One or Many, not both")
	ErrRelationMissingTypes       = errors.New("relation is missing its defined types and fields")
	ErrRelationInvalidType        = errors.New("relation has an invalid type to be finalize")
	ErrMultipleRelationPrimaries  = errors.New("relation can only have a single field set as primary")
	// NonNull is the literal name of the GQL type, so we have to disable the linter
	//nolint:revive
	ErrNonNullNotSupported = errors.New("NonNull fields are not currently supported")
)

func NewErrDuplicateField(objectName, fieldName string) error {
	return errors.New(
		errDuplicateField,
		errors.NewKV("Object", objectName),
		errors.NewKV("Field", fieldName),
	)
}

func NewErrFieldMissingRelation(objectName, fieldName string, objectType string) error {
	return errors.New(
		errFieldMissingRelation,
		errors.NewKV("Object", objectName),
		errors.NewKV("Field", fieldName),
		errors.NewKV("ObjectType", objectType),
	)
}

func NewErrRelationMissingField(objectName, fieldName string) error {
	return errors.New(
		errRelationMissingField,
		errors.NewKV("Object", objectName),
		errors.NewKV("Field", fieldName),
	)
}

func NewErrAggregateTargetNotFound(objectName, target string) error {
	return errors.New(
		errAggregateTargetNotFound,
		errors.NewKV("Object", objectName),
		errors.NewKV("Target", target),
	)
}

func NewErrSchemaTypeAlreadyExist(name string) error {
	return errors.New(
		errSchemaTypeAlreadyExist,
		errors.NewKV("Name", name),
	)
}

func NewErrObjectNotFoundDuringThunk(object string) error {
	return errors.New(
		errObjectNotFoundDuringThunk,
		errors.NewKV("Object", object),
	)
}

func NewErrTypeNotFound(typeName string) error {
	return errors.New(
		errTypeNotFound,
		errors.NewKV("Type", typeName),
	)
}

func NewErrNonNullForTypeNotSupported(typeName string) error {
	return errors.New(
		errNonNullForTypeNotSupported,
		errors.NewKV("Type", typeName),
	)
}

func NewErrRelationNotFound(relationName string) error {
	return errors.New(
		errRelationNotFound,
		errors.NewKV("RelationName", relationName),
	)
}
