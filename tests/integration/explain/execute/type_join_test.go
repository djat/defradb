// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package test_explain_execute

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestExecuteExplainRequestWithAOneToOneJoin(t *testing.T) {
	test := testUtils.TestCase{

		Description: "Explain a one-to-one join relation query, with alias.",

		Actions: []any{
			gqlSchemaExecuteExplain(),

			// Authors
			create2AuthorDocuments(),

			// Contacts
			create2AuthorContactDocuments(),

			testUtils.Request{
				Request: `query @explain(type: execute) {
					Author {
						OnlyEmail: contact {
							email
						}
					}
				}`,

				Results: []dataMap{
					{
						"explain": dataMap{
							"executionSuccess": true,
							"sizeOfResult":     2,
							"planExecutions":   uint64(3),
							"selectTopNode": dataMap{
								"selectNode": dataMap{
									"iterations":    uint64(3),
									"filterMatches": uint64(2),
									"typeIndexJoin": dataMap{
										"iterations": uint64(3),
										"scanNode": dataMap{
											"iterations":    uint64(3),
											"docFetches":    uint64(3),
											"filterMatches": uint64(2),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	executeTestCase(t, test)
}

func TestExecuteExplainWithMultipleOneToOneJoins(t *testing.T) {
	test := testUtils.TestCase{

		Description: "Explain (execute) with two one-to-one join relation.",

		Actions: []any{
			gqlSchemaExecuteExplain(),

			// Authors
			create2AuthorDocuments(),

			// Contacts
			create2AuthorContactDocuments(),

			testUtils.Request{
				Request: `query @explain(type: execute) {
					Author {
						OnlyEmail: contact {
							email
						}
						contact {
							cell
							email
						}
					}
				}`,

				Results: []dataMap{
					{
						"explain": dataMap{
							"executionSuccess": true,
							"sizeOfResult":     2,
							"planExecutions":   uint64(3),
							"selectTopNode": dataMap{
								"selectNode": dataMap{
									"iterations":    uint64(3),
									"filterMatches": uint64(2),
									"parallelNode": []dataMap{
										{
											"typeIndexJoin": dataMap{
												"iterations": uint64(3),
												"scanNode": dataMap{
													"iterations":    uint64(3),
													"docFetches":    uint64(3),
													"filterMatches": uint64(2),
												},
											},
										},
										{
											"typeIndexJoin": dataMap{
												"iterations": uint64(3),
												"scanNode": dataMap{
													"iterations":    uint64(3),
													"docFetches":    uint64(3),
													"filterMatches": uint64(2),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	executeTestCase(t, test)
}

func TestExecuteExplainWithTwoLevelDeepNestedJoins(t *testing.T) {
	test := testUtils.TestCase{

		Description: "Explain (execute) with two nested level deep one to one join.",

		Actions: []any{
			gqlSchemaExecuteExplain(),

			// Authors
			create2AuthorDocuments(),

			// Contacts
			create2AuthorContactDocuments(),

			// Addresses
			create2AddressDocuments(),

			testUtils.Request{
				Request: `query @explain(type: execute) {
					Author {
						name
						contact {
							email
							address {
								city
							}
						}
					}
				}`,

				Results: []dataMap{
					{
						"explain": dataMap{
							"executionSuccess": true,
							"sizeOfResult":     2,
							"planExecutions":   uint64(3),
							"selectTopNode": dataMap{
								"selectNode": dataMap{
									"iterations":    uint64(3),
									"filterMatches": uint64(2),
									"typeIndexJoin": dataMap{
										"iterations": uint64(3),
										"scanNode": dataMap{
											"iterations":    uint64(3),
											"docFetches":    uint64(3),
											"filterMatches": uint64(2),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	executeTestCase(t, test)
}
