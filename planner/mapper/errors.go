// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package mapper

import "github.com/sourcenetwork/defradb/errors"

var (
	ErrUnableToIdAggregateChild = errors.New("unable to identify aggregate child")
	ErrAggregateTargetMissing   = errors.New("aggregate must be provided with a property to aggregate")
	ErrFailedToFindHostField    = errors.New("failed to find host field")
)
