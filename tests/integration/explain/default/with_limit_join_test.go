// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package test_explain_default

import (
	"testing"

	explainUtils "github.com/sourcenetwork/defradb/tests/integration/explain"
)

var limitTypeJoinPattern = dataMap{
	"root": dataMap{
		"scanNode": dataMap{},
	},
	"subType": dataMap{
		"selectTopNode": dataMap{
			"limitNode": dataMap{
				"selectNode": dataMap{
					"scanNode": dataMap{},
				},
			},
		},
	},
}

func TestDefaultExplainRequestWithOnlyLimitOnRelatedChild(t *testing.T) {
	test := explainUtils.ExplainRequestTestCase{

		Description: "Explain (default) request with only limit on related child.",

		Request: `query @explain {
			Author {
				name
				articles(limit: 1) {
					name
				}
			}
		}`,

		Docs: map[int][]string{
			// articles
			0: {
				`{
					"name": "After Guantánamo, Another Injustice",
					"author_id": "bae-41598f0c-19bc-5da6-813b-e80f14a10df3"
				}`,

				`{
					"name": "To my dear readers",
					"author_id": "bae-b769708d-f552-5c3d-a402-ccfd7ac7fb04"
				}`,

				`{
					"name": "Twinklestar's Favourite Xmas Cookie",
					"author_id": "bae-b769708d-f552-5c3d-a402-ccfd7ac7fb04"
				}`,

				`{
					"name": "C++ 100",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "C++ 101",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "C++ 200",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "C++ 202",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "Rust 100",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,

				`{
					"name": "Rust 101",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,

				`{
					"name": "Rust 200",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,

				`{
					"name": "Rust 202",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,
			},

			// authors
			2: {
				// _key: bae-41598f0c-19bc-5da6-813b-e80f14a10df3
				`{
					"name": "John Grisham",
					"age": 65,
					"verified": true
				}`,

				// _key: bae-aa839756-588e-5b57-887d-33689a06e375
				`{
					"name": "Shahzad Sisley",
					"age": 26,
					"verified": true
				}`,

				// _key: bae-b769708d-f552-5c3d-a402-ccfd7ac7fb04
				`{
					"name": "Cornelia Funke",
					"age": 62,
					"verified": false
				}`,

				// _key: bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69
				`{
					"name": "Andrew Lone",
					"age": 28,
					"verified": true
				}`,
			},
		},

		ExpectedPatterns: []dataMap{
			{
				"explain": dataMap{
					"selectTopNode": dataMap{
						"selectNode": dataMap{
							"typeIndexJoin": limitTypeJoinPattern,
						},
					},
				},
			},
		},

		ExpectedTargets: []explainUtils.PlanNodeTargetCase{
			{
				TargetNodeName:    "limitNode",
				IncludeChildNodes: false,
				ExpectedAttributes: dataMap{
					"limit":  uint64(1),
					"offset": uint64(0),
				},
			},
		},
	}

	runExplainTest(t, test)
}

func TestDefaultExplainRequestWithOnlyOffsetOnRelatedChild(t *testing.T) {
	test := explainUtils.ExplainRequestTestCase{

		Description: "Explain (default) request with only offset on related child.",

		Request: `query @explain {
			Author {
				name
				articles(offset: 2) {
					name
				}
			}
		}`,

		Docs: map[int][]string{
			// articles
			0: {
				`{
					"name": "After Guantánamo, Another Injustice",
					"author_id": "bae-41598f0c-19bc-5da6-813b-e80f14a10df3"
				}`,

				`{
					"name": "To my dear readers",
					"author_id": "bae-b769708d-f552-5c3d-a402-ccfd7ac7fb04"
				}`,

				`{
					"name": "Twinklestar's Favourite Xmas Cookie",
					"author_id": "bae-b769708d-f552-5c3d-a402-ccfd7ac7fb04"
				}`,

				`{
					"name": "C++ 100",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "C++ 101",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "C++ 200",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "C++ 202",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "Rust 100",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,

				`{
					"name": "Rust 101",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,

				`{
					"name": "Rust 200",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,

				`{
					"name": "Rust 202",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,
			},

			// authors
			2: {
				// _key: bae-41598f0c-19bc-5da6-813b-e80f14a10df3
				`{
					"name": "John Grisham",
					"age": 65,
					"verified": true
				}`,

				// _key: bae-aa839756-588e-5b57-887d-33689a06e375
				`{
					"name": "Shahzad Sisley",
					"age": 26,
					"verified": true
				}`,

				// _key: bae-b769708d-f552-5c3d-a402-ccfd7ac7fb04
				`{
					"name": "Cornelia Funke",
					"age": 62,
					"verified": false
				}`,

				// _key: bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69
				`{
					"name": "Andrew Lone",
					"age": 28,
					"verified": true
				}`,
			},
		},

		ExpectedPatterns: []dataMap{
			{
				"explain": dataMap{
					"selectTopNode": dataMap{
						"selectNode": dataMap{
							"typeIndexJoin": limitTypeJoinPattern,
						},
					},
				},
			},
		},

		ExpectedTargets: []explainUtils.PlanNodeTargetCase{
			{
				TargetNodeName:    "limitNode",
				IncludeChildNodes: false,
				ExpectedAttributes: dataMap{
					"limit":  nil,
					"offset": uint64(2),
				},
			},
		},
	}

	runExplainTest(t, test)
}

func TestDefaultExplainRequestWithBothLimitAndOffsetOnRelatedChild(t *testing.T) {
	test := explainUtils.ExplainRequestTestCase{

		Description: "Explain (default) request with both limit and offset on related child.",

		Request: `query @explain {
			Author {
				name
				articles(limit: 2, offset: 2) {
					name
				}
			}
		}`,

		Docs: map[int][]string{
			// articles
			0: {
				`{
					"name": "After Guantánamo, Another Injustice",
					"author_id": "bae-41598f0c-19bc-5da6-813b-e80f14a10df3"
				}`,

				`{
					"name": "To my dear readers",
					"author_id": "bae-b769708d-f552-5c3d-a402-ccfd7ac7fb04"
				}`,

				`{
					"name": "Twinklestar's Favourite Xmas Cookie",
					"author_id": "bae-b769708d-f552-5c3d-a402-ccfd7ac7fb04"
				}`,

				`{
					"name": "C++ 100",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "C++ 101",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "C++ 200",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "C++ 202",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "Rust 100",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,

				`{
					"name": "Rust 101",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,

				`{
					"name": "Rust 200",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,

				`{
					"name": "Rust 202",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,
			},

			// authors
			2: {
				// _key: bae-41598f0c-19bc-5da6-813b-e80f14a10df3
				`{
					"name": "John Grisham",
					"age": 65,
					"verified": true
				}`,

				// _key: bae-aa839756-588e-5b57-887d-33689a06e375
				`{
					"name": "Shahzad Sisley",
					"age": 26,
					"verified": true
				}`,

				// _key: bae-b769708d-f552-5c3d-a402-ccfd7ac7fb04
				`{
					"name": "Cornelia Funke",
					"age": 62,
					"verified": false
				}`,

				// _key: bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69
				`{
					"name": "Andrew Lone",
					"age": 28,
					"verified": true
				}`,
			},
		},

		ExpectedPatterns: []dataMap{
			{
				"explain": dataMap{
					"selectTopNode": dataMap{
						"selectNode": dataMap{
							"typeIndexJoin": limitTypeJoinPattern,
						},
					},
				},
			},
		},

		ExpectedTargets: []explainUtils.PlanNodeTargetCase{
			{
				TargetNodeName:    "limitNode",
				IncludeChildNodes: false,
				ExpectedAttributes: dataMap{
					"limit":  uint64(2),
					"offset": uint64(2),
				},
			},
		},
	}

	runExplainTest(t, test)
}

func TestDefaultExplainRequestWithLimitOnRelatedChildAndBothLimitAndOffsetOnParent(t *testing.T) {
	test := explainUtils.ExplainRequestTestCase{

		Description: "Explain (default) request with limit on related child & both limit + offset on parent.",

		Request: `query @explain {
			Author(limit: 3, offset: 1) {
				name
				articles(limit: 2) {
					name
				}
			}
		}`,

		Docs: map[int][]string{
			// articles
			0: {
				`{
					"name": "After Guantánamo, Another Injustice",
					"author_id": "bae-41598f0c-19bc-5da6-813b-e80f14a10df3"
				}`,

				`{
					"name": "To my dear readers",
					"author_id": "bae-b769708d-f552-5c3d-a402-ccfd7ac7fb04"
				}`,

				`{
					"name": "Twinklestar's Favourite Xmas Cookie",
					"author_id": "bae-b769708d-f552-5c3d-a402-ccfd7ac7fb04"
				}`,

				`{
					"name": "C++ 100",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "C++ 101",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "C++ 200",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "C++ 202",
					"author_id": "bae-aa839756-588e-5b57-887d-33689a06e375"
				}`,

				`{
					"name": "Rust 100",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,

				`{
					"name": "Rust 101",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,

				`{
					"name": "Rust 200",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,

				`{
					"name": "Rust 202",
					"author_id": "bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69"
				}`,
			},

			// authors
			2: {
				// _key: bae-41598f0c-19bc-5da6-813b-e80f14a10df3
				`{
					"name": "John Grisham",
					"age": 65,
					"verified": true
				}`,

				// _key: bae-aa839756-588e-5b57-887d-33689a06e375
				`{
					"name": "Shahzad Sisley",
					"age": 26,
					"verified": true
				}`,

				// _key: bae-b769708d-f552-5c3d-a402-ccfd7ac7fb04
				`{
					"name": "Cornelia Funke",
					"age": 62,
					"verified": false
				}`,

				// _key: bae-e7e87bbb-1079-59db-b4b9-0e14b24d5b69
				`{
					"name": "Andrew Lone",
					"age": 28,
					"verified": true
				}`,
			},
		},

		ExpectedPatterns: []dataMap{
			{
				"explain": dataMap{
					"selectTopNode": dataMap{
						"limitNode": dataMap{
							"selectNode": dataMap{
								"typeIndexJoin": limitTypeJoinPattern,
							},
						},
					},
				},
			},
		},

		ExpectedTargets: []explainUtils.PlanNodeTargetCase{
			{
				TargetNodeName:    "limitNode",
				OccurancesToSkip:  0,
				IncludeChildNodes: false,
				ExpectedAttributes: dataMap{
					"limit":  uint64(3),
					"offset": uint64(1),
				},
			},
			{
				TargetNodeName:    "limitNode",
				OccurancesToSkip:  1,
				IncludeChildNodes: false,
				ExpectedAttributes: dataMap{
					"limit":  uint64(2),
					"offset": uint64(0),
				},
			},
		},
	}

	runExplainTest(t, test)
}
