// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package simple

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestQuerySimpleWithLikeStringContainsFilterBlockContainsString(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with basic like-string filter contains string",
		Request: `query {
					Users(filter: {Name: {_like: "%Stormborn%"}}) {
						Name
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
					"HeightM": 1.65
				}`,
				`{
					"Name": "Viserys I Targaryen, King of the Andals",
					"HeightM": 1.82
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
			},
		},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithLikeStringContainsFilterBlockAsPrefixString(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with basic like-string filter with string as prefix",
		Request: `query {
					Users(filter: {Name: {_like: "Viserys%"}}) {
						Name
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
					"HeightM": 1.65
				}`,
				`{
					"Name": "Viserys I Targaryen, King of the Andals",
					"HeightM": 1.82
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Viserys I Targaryen, King of the Andals",
			},
		},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithLikeStringContainsFilterBlockAsSuffixString(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with basic like-string filter with string as suffix",
		Request: `query {
					Users(filter: {Name: {_like: "%Andals"}}) {
						Name
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
					"HeightM": 1.65
				}`,
				`{
					"Name": "Viserys I Targaryen, King of the Andals",
					"HeightM": 1.82
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Viserys I Targaryen, King of the Andals",
			},
		},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithLikeStringContainsFilterBlockExactString(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with basic like-string filter with string as suffix",
		Request: `query {
					Users(filter: {Name: {_like: "Daenerys Stormborn of House Targaryen, the First of Her Name"}}) {
						Name
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
					"HeightM": 1.65
				}`,
				`{
					"Name": "Viserys I Targaryen, King of the Andals",
					"HeightM": 1.82
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
			},
		},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithLikeStringContainsFilterBlockContainsStringMuplitpleResults(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with basic like-string filter with contains string mulitple results",
		Request: `query {
					Users(filter: {Name: {_like: "%Targaryen%"}}) {
						Name
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
					"HeightM": 1.65
				}`,
				`{
					"Name": "Viserys I Targaryen, King of the Andals",
					"HeightM": 1.82
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
			},
			{
				"Name": "Viserys I Targaryen, King of the Andals",
			},
		},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithLikeStringContainsFilterBlockHasStartAndEnd(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with basic like-string filter with string as start and end",
		Request: `query {
					Users(filter: {Name: {_like: "Daenerys%Name"}}) {
						Name
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
					"HeightM": 1.65
				}`,
				`{
					"Name": "Viserys I Targaryen, King of the Andals",
					"HeightM": 1.82
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
			},
		},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithLikeStringContainsFilterBlockHasBoth(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with basic like-string filter with none of the strings",
		Request: `query {
					Users(filter: {_and: [{Name: {_like: "%Baratheon%"}}, {Name: {_like: "%Stormborn%"}}]}) {
						Name
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
					"HeightM": 1.65
				}`,
				`{
					"Name": "Viserys I Targaryen, King of the Andals",
					"HeightM": 1.82
				}`,
			},
		},
		Results: []map[string]any{},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithLikeStringContainsFilterBlockHasEither(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with basic like-string filter with either strings",
		Request: `query {
					Users(filter: {_or: [{Name: {_like: "%Baratheon%"}}, {Name: {_like: "%Stormborn%"}}]}) {
						Name
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
					"HeightM": 1.65
				}`,
				`{
					"Name": "Viserys I Targaryen, King of the Andals",
					"HeightM": 1.82
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
			},
		},
	}

	executeTestCase(t, test)
}

func TestQuerySimpleWithLikeStringContainsFilterBlockPropNotSet(t *testing.T) {
	test := testUtils.RequestTestCase{
		Description: "Simple query with basic like-string filter with either strings",
		Request: `query {
					Users(filter: {Name: {_like: "%King%"}}) {
						Name
					}
				}`,
		Docs: map[int][]string{
			0: {
				`{
					"Name": "Daenerys Stormborn of House Targaryen, the First of Her Name",
					"HeightM": 1.65
				}`,
				`{
					"Name": "Viserys I Targaryen, King of the Andals",
					"HeightM": 1.82
				}`,
				`{
					"HeightM": 1.92
				}`,
			},
		},
		Results: []map[string]any{
			{
				"Name": "Viserys I Targaryen, King of the Andals",
			},
		},
	}

	executeTestCase(t, test)
}
