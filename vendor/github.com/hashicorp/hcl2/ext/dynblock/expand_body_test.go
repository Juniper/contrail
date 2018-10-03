package dynblock

import (
	"testing"

	"github.com/hashicorp/hcl2/hcldec"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcltest"
	"github.com/zclconf/go-cty/cty"
)

func TestExpand(t *testing.T) {
	srcBody := hcltest.MockBody(&hcl.BodyContent{
		Blocks: hcl.Blocks{
			{
				Type:        "a",
				Labels:      []string{"static0"},
				LabelRanges: []hcl.Range{hcl.Range{}},
				Body: hcltest.MockBody(&hcl.BodyContent{
					Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
						"val": hcltest.MockExprLiteral(cty.StringVal("static a 0")),
					}),
				}),
			},
			{
				Type: "b",
				Body: hcltest.MockBody(&hcl.BodyContent{
					Blocks: hcl.Blocks{
						{
							Type: "c",
							Body: hcltest.MockBody(&hcl.BodyContent{
								Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
									"val0": hcltest.MockExprLiteral(cty.StringVal("static c 0")),
								}),
							}),
						},
						{
							Type:        "dynamic",
							Labels:      []string{"c"},
							LabelRanges: []hcl.Range{hcl.Range{}},
							Body: hcltest.MockBody(&hcl.BodyContent{
								Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
									"for_each": hcltest.MockExprLiteral(cty.ListVal([]cty.Value{
										cty.StringVal("dynamic c 0"),
										cty.StringVal("dynamic c 1"),
									})),
									"iterator": hcltest.MockExprVariable("dyn_c"),
								}),
								Blocks: hcl.Blocks{
									{
										Type: "content",
										Body: hcltest.MockBody(&hcl.BodyContent{
											Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
												"val0": hcltest.MockExprTraversalSrc("dyn_c.value"),
											}),
										}),
									},
								},
							}),
						},
					},
				}),
			},
			{
				Type:        "dynamic",
				Labels:      []string{"a"},
				LabelRanges: []hcl.Range{hcl.Range{}},
				Body: hcltest.MockBody(&hcl.BodyContent{
					Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
						"for_each": hcltest.MockExprLiteral(cty.ListVal([]cty.Value{
							cty.StringVal("dynamic a 0"),
							cty.StringVal("dynamic a 1"),
							cty.StringVal("dynamic a 2"),
						})),
						"labels": hcltest.MockExprList([]hcl.Expression{
							hcltest.MockExprTraversalSrc("a.key"),
						}),
					}),
					Blocks: hcl.Blocks{
						{
							Type: "content",
							Body: hcltest.MockBody(&hcl.BodyContent{
								Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
									"val": hcltest.MockExprTraversalSrc("a.value"),
								}),
							}),
						},
					},
				}),
			},
			{
				Type:        "dynamic",
				Labels:      []string{"b"},
				LabelRanges: []hcl.Range{hcl.Range{}},
				Body: hcltest.MockBody(&hcl.BodyContent{
					Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
						"for_each": hcltest.MockExprLiteral(cty.ListVal([]cty.Value{
							cty.StringVal("dynamic b 0"),
							cty.StringVal("dynamic b 1"),
						})),
						"iterator": hcltest.MockExprVariable("dyn_b"),
					}),
					Blocks: hcl.Blocks{
						{
							Type: "content",
							Body: hcltest.MockBody(&hcl.BodyContent{
								Blocks: hcl.Blocks{
									{
										Type: "c",
										Body: hcltest.MockBody(&hcl.BodyContent{
											Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
												"val0": hcltest.MockExprLiteral(cty.StringVal("static c 1")),
												"val1": hcltest.MockExprTraversalSrc("dyn_b.value"),
											}),
										}),
									},
									{
										Type:        "dynamic",
										Labels:      []string{"c"},
										LabelRanges: []hcl.Range{hcl.Range{}},
										Body: hcltest.MockBody(&hcl.BodyContent{
											Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
												"for_each": hcltest.MockExprLiteral(cty.ListVal([]cty.Value{
													cty.StringVal("dynamic c 2"),
													cty.StringVal("dynamic c 3"),
												})),
											}),
											Blocks: hcl.Blocks{
												{
													Type: "content",
													Body: hcltest.MockBody(&hcl.BodyContent{
														Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
															"val0": hcltest.MockExprTraversalSrc("c.value"),
															"val1": hcltest.MockExprTraversalSrc("dyn_b.value"),
														}),
													}),
												},
											},
										}),
									},
								},
							}),
						},
					},
				}),
			},
			{
				Type:        "dynamic",
				Labels:      []string{"b"},
				LabelRanges: []hcl.Range{hcl.Range{}},
				Body: hcltest.MockBody(&hcl.BodyContent{
					Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
						"for_each": hcltest.MockExprLiteral(cty.MapVal(map[string]cty.Value{
							"foo": cty.ListVal([]cty.Value{
								cty.StringVal("dynamic c nested 0"),
								cty.StringVal("dynamic c nested 1"),
							}),
						})),
						"iterator": hcltest.MockExprVariable("dyn_b"),
					}),
					Blocks: hcl.Blocks{
						{
							Type: "content",
							Body: hcltest.MockBody(&hcl.BodyContent{
								Blocks: hcl.Blocks{
									{
										Type:        "dynamic",
										Labels:      []string{"c"},
										LabelRanges: []hcl.Range{hcl.Range{}},
										Body: hcltest.MockBody(&hcl.BodyContent{
											Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
												"for_each": hcltest.MockExprTraversalSrc("dyn_b.value"),
											}),
											Blocks: hcl.Blocks{
												{
													Type: "content",
													Body: hcltest.MockBody(&hcl.BodyContent{
														Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
															"val0": hcltest.MockExprTraversalSrc("c.value"),
															"val1": hcltest.MockExprTraversalSrc("dyn_b.key"),
														}),
													}),
												},
											},
										}),
									},
								},
							}),
						},
					},
				}),
			},
			{
				Type:        "a",
				Labels:      []string{"static1"},
				LabelRanges: []hcl.Range{hcl.Range{}},
				Body: hcltest.MockBody(&hcl.BodyContent{
					Attributes: hcltest.MockAttrs(map[string]hcl.Expression{
						"val": hcltest.MockExprLiteral(cty.StringVal("static a 1")),
					}),
				}),
			},
		},
	})

	dynBody := Expand(srcBody, nil)
	var remain hcl.Body

	t.Run("PartialDecode", func(t *testing.T) {
		decSpec := &hcldec.BlockMapSpec{
			TypeName:   "a",
			LabelNames: []string{"key"},
			Nested: &hcldec.AttrSpec{
				Name:     "val",
				Type:     cty.String,
				Required: true,
			},
		}

		var got cty.Value
		var diags hcl.Diagnostics
		got, remain, diags = hcldec.PartialDecode(dynBody, decSpec, nil)
		if len(diags) != 0 {
			t.Errorf("unexpected diagnostics")
			for _, diag := range diags {
				t.Logf("- %s", diag)
			}
			return
		}

		want := cty.MapVal(map[string]cty.Value{
			"static0": cty.StringVal("static a 0"),
			"static1": cty.StringVal("static a 1"),
			"0":       cty.StringVal("dynamic a 0"),
			"1":       cty.StringVal("dynamic a 1"),
			"2":       cty.StringVal("dynamic a 2"),
		})

		if !got.RawEquals(want) {
			t.Errorf("wrong result\ngot:  %#v\nwant: %#v", got, want)
		}
	})

	t.Run("Decode", func(t *testing.T) {
		decSpec := &hcldec.BlockListSpec{
			TypeName: "b",
			Nested: &hcldec.BlockListSpec{
				TypeName: "c",
				Nested: &hcldec.ObjectSpec{
					"val0": &hcldec.AttrSpec{
						Name: "val0",
						Type: cty.String,
					},
					"val1": &hcldec.AttrSpec{
						Name: "val1",
						Type: cty.String,
					},
				},
			},
		}

		var got cty.Value
		var diags hcl.Diagnostics
		got, diags = hcldec.Decode(remain, decSpec, nil)
		if len(diags) != 0 {
			t.Errorf("unexpected diagnostics")
			for _, diag := range diags {
				t.Logf("- %s", diag)
			}
			return
		}

		want := cty.ListVal([]cty.Value{
			cty.ListVal([]cty.Value{
				cty.ObjectVal(map[string]cty.Value{
					"val0": cty.StringVal("static c 0"),
					"val1": cty.NullVal(cty.String),
				}),
				cty.ObjectVal(map[string]cty.Value{
					"val0": cty.StringVal("dynamic c 0"),
					"val1": cty.NullVal(cty.String),
				}),
				cty.ObjectVal(map[string]cty.Value{
					"val0": cty.StringVal("dynamic c 1"),
					"val1": cty.NullVal(cty.String),
				}),
			}),
			cty.ListVal([]cty.Value{
				cty.ObjectVal(map[string]cty.Value{
					"val0": cty.StringVal("static c 1"),
					"val1": cty.StringVal("dynamic b 0"),
				}),
				cty.ObjectVal(map[string]cty.Value{
					"val0": cty.StringVal("dynamic c 2"),
					"val1": cty.StringVal("dynamic b 0"),
				}),
				cty.ObjectVal(map[string]cty.Value{
					"val0": cty.StringVal("dynamic c 3"),
					"val1": cty.StringVal("dynamic b 0"),
				}),
			}),
			cty.ListVal([]cty.Value{
				cty.ObjectVal(map[string]cty.Value{
					"val0": cty.StringVal("static c 1"),
					"val1": cty.StringVal("dynamic b 1"),
				}),
				cty.ObjectVal(map[string]cty.Value{
					"val0": cty.StringVal("dynamic c 2"),
					"val1": cty.StringVal("dynamic b 1"),
				}),
				cty.ObjectVal(map[string]cty.Value{
					"val0": cty.StringVal("dynamic c 3"),
					"val1": cty.StringVal("dynamic b 1"),
				}),
			}),
			cty.ListVal([]cty.Value{
				cty.ObjectVal(map[string]cty.Value{
					"val0": cty.StringVal("dynamic c nested 0"),
					"val1": cty.StringVal("foo"),
				}),
				cty.ObjectVal(map[string]cty.Value{
					"val0": cty.StringVal("dynamic c nested 1"),
					"val1": cty.StringVal("foo"),
				}),
			}),
		})

		if !got.RawEquals(want) {
			t.Errorf("wrong result\ngot:  %#v\nwant: %#v", got, want)
		}
	})

}
