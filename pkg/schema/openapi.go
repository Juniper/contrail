package schema

import (
	"github.com/go-openapi/spec"
)

//ToOpenAPI generates OpenAPI commands.
// nolint: gocyclo
func (api *API) ToOpenAPI() (*spec.Swagger, error) {
	definitions := spec.Definitions{}
	paths := &spec.Paths{
		Paths: map[string]spec.PathItem{},
	}
	openAPI := &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Swagger:  "2.0",
			BasePath: "/",
			Schemes:  []string{"https"},
			Consumes: []string{"application/json"},
			Produces: []string{"application/json"},
			Info: &spec.Info{
				InfoProps: spec.InfoProps{
					Version: "4.0",
					Title:   "Contrail API OpenAPI2.0 Definitions",
					License: &spec.License{
						Name: "Apache2.0",
					},
				},
			},
			Paths:       paths,
			Definitions: definitions,
		},
	}
	for _, apiSchema := range api.Schemas {
		if apiSchema.Type == AbstractType {
			continue
		}
		d, err := apiSchema.JSONSchema.ToOpenAPI()
		if err != nil {
			return nil, err
		}
		// add reference and back ref

		for _, reference := range apiSchema.References {
			referenceSchema := spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: reference.Description,
					Properties: map[string]spec.Schema{
						"uuid": {
							SchemaProps: spec.SchemaProps{
								Type: spec.StringOrArray([]string{"string"}),
							},
						},
						"to": {
							SchemaProps: spec.SchemaProps{
								Type: spec.StringOrArray([]string{"array"}),
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type: spec.StringOrArray([]string{"string"}),
										},
									},
								},
							},
						},
					},
				},
			}
			var ref spec.Ref
			ref, err = spec.NewRef("#/definitions/" + reference.RefType)
			if err != nil {
				return nil, err
			}
			if reference.RefType != "" {
				referenceSchema.Properties["attr"] = spec.Schema{
					SchemaProps: spec.SchemaProps{
						Ref: ref,
					},
				}
			}
			d.Properties[reference.LinkTo.ID+"_ref"] = referenceSchema
		}

		for _, backref := range apiSchema.Children {
			var ref spec.Ref
			ref, err = spec.NewRef("#/definitions/" + backref.LinkTo.JSONSchema.GoName + "APIType")
			if err != nil {
				return nil, err
			}
			d.Properties[backref.LinkTo.ID+"s"] = spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: backref.Description,
					Type:        spec.StringOrArray([]string{"array"}),
					Items: &spec.SchemaOrArray{
						Schema: &spec.Schema{
							SchemaProps: spec.SchemaProps{
								Ref: ref,
							},
						},
					},
				},
			}
		}
		definitions[apiSchema.JSONSchema.GoName+"APIType"] = *d

		ref, err := spec.NewRef("#/definitions/" + apiSchema.JSONSchema.GoName + "APIType")
		if err != nil {
			return nil, err
		}
		listAPIRef, err := spec.NewRef("#/definitions/" + apiSchema.JSONSchema.GoName + "APIListType")
		if err != nil {
			return nil, err
		}
		singleAPIRef, err := spec.NewRef("#/definitions/" + apiSchema.JSONSchema.GoName + "APISingleType")
		if err != nil {
			return nil, err
		}
		//TODO add path for this resource.

		pathItem := spec.PathItem{
			PathItemProps: spec.PathItemProps{
				Get: &spec.Operation{
					OperationProps: spec.OperationProps{
						//TODO Parameters:
						Responses: &spec.Responses{
							ResponsesProps: spec.ResponsesProps{
								StatusCodeResponses: map[int]spec.Response{
									200: {
										ResponseProps: spec.ResponseProps{
											Description: "Show resource",
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Ref: singleAPIRef,
												},
											},
										},
									},
									404: {
										ResponseProps: spec.ResponseProps{
											Description: "Resource not found",
										},
									},
									401: {
										ResponseProps: spec.ResponseProps{
											Description: "Unauthorized",
										},
									},
									500: {
										ResponseProps: spec.ResponseProps{
											Description: "Server Side Error",
										},
									},
								},
							},
						},
					},
				},
				//TODO
				Delete: &spec.Operation{
					OperationProps: spec.OperationProps{
						//TODO Parameters:
						Responses: &spec.Responses{
							ResponsesProps: spec.ResponsesProps{
								StatusCodeResponses: map[int]spec.Response{
									200: {
										ResponseProps: spec.ResponseProps{
											Description: "Delete a resource",
										},
									},
									401: {
										ResponseProps: spec.ResponseProps{
											Description: "Unauthorized",
										},
									},
									404: {
										ResponseProps: spec.ResponseProps{
											Description: "Resource not found",
										},
									},
									409: {
										ResponseProps: spec.ResponseProps{
											Description: "Data conflict",
										},
									},
									500: {
										ResponseProps: spec.ResponseProps{
											Description: "Server Side Error",
										},
									},
								},
							},
						},
					},
				},
				//TODO
				Put: &spec.Operation{
					OperationProps: spec.OperationProps{
						Parameters: []spec.Parameter{
							{
								ParamProps: spec.ParamProps{
									Required: true,
									In:       "body",
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Ref: singleAPIRef,
										},
									},
								},
							},
						},
						Responses: &spec.Responses{
							ResponsesProps: spec.ResponsesProps{
								StatusCodeResponses: map[int]spec.Response{
									200: {
										ResponseProps: spec.ResponseProps{
											Description: "Update a resource",
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Ref: singleAPIRef,
												},
											},
										},
									},
									400: {
										ResponseProps: spec.ResponseProps{
											Description: "Bad request",
										},
									},
									401: {
										ResponseProps: spec.ResponseProps{
											Description: "Unauthorized",
										},
									},
									404: {
										ResponseProps: spec.ResponseProps{
											Description: "Resource not found",
										},
									},
									409: {
										ResponseProps: spec.ResponseProps{
											Description: "Data conflict",
										},
									},
									500: {
										ResponseProps: spec.ResponseProps{
											Description: "Server Side Error",
										},
									},
								},
							},
						},
					},
				},
			},
		}
		pluralPathItem := spec.PathItem{
			PathItemProps: spec.PathItemProps{
				Post: &spec.Operation{
					OperationProps: spec.OperationProps{
						Parameters: []spec.Parameter{
							{
								ParamProps: spec.ParamProps{
									In:       "body",
									Required: true,
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Ref: singleAPIRef,
										},
									},
								},
							},
						},
						Responses: &spec.Responses{
							ResponsesProps: spec.ResponsesProps{
								StatusCodeResponses: map[int]spec.Response{
									200: {
										ResponseProps: spec.ResponseProps{
											Description: "Create a resource",
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Ref: singleAPIRef,
												},
											},
										},
									},
									400: {
										ResponseProps: spec.ResponseProps{
											Description: "Bad request",
										},
									},
									401: {
										ResponseProps: spec.ResponseProps{
											Description: "Unauthorized",
										},
									},
									404: {
										ResponseProps: spec.ResponseProps{
											Description: "Resource not found",
										},
									},
									409: {
										ResponseProps: spec.ResponseProps{
											Description: "Data conflict",
										},
									},
									500: {
										ResponseProps: spec.ResponseProps{
											Description: "Server Side Error",
										},
									},
								},
							},
						},
					},
				},
				Get: &spec.Operation{
					OperationProps: spec.OperationProps{
						Parameters: []spec.Parameter{
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "parent_id",
									Description: "parent_uuid",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "parent_fq_name_str",
									Description: "parent’s fully-qualified name delimited by ‘:’",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "pobj_uuids",
									Description: "Commna separated object uuids <example1_uuid>,<example2_uuid>",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "detail",
									Description: "True if you need detailed data",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "back_ref_id",
									Description: "back_ref_uuid",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "page_marker",
									Description: "Pagenation start marker",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "page_limit",
									Description: "Pagenation limit",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "count",
									Description: "Return only resource counts",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "fields",
									Description: " Comma separated object field list you are interested in",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "shared",
									Description: "Included shared object in response.",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "filters",
									Description: "Comma separated fileter list. Example check==a,check==b,name==Bob",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "exclude_hrefs",
									Description: "",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "exclude_clildren",
									Description: "",
									Required:    false,
								},
							},
							{
								ParamProps: spec.ParamProps{
									In:          "query",
									Name:        "exclude_back_refs",
									Description: "",
									Required:    false,
								},
							},
						},
						Responses: &spec.Responses{
							ResponsesProps: spec.ResponsesProps{
								StatusCodeResponses: map[int]spec.Response{
									200: {
										ResponseProps: spec.ResponseProps{
											Description: "list a resource",
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Ref: listAPIRef,
												},
											},
										},
									},
									400: {
										ResponseProps: spec.ResponseProps{
											Description: "Bad request",
										},
									},
									404: {
										ResponseProps: spec.ResponseProps{
											Description: "Resource not found",
										},
									},
									401: {
										ResponseProps: spec.ResponseProps{
											Description: "Unauthorized",
										},
									},
									500: {
										ResponseProps: spec.ResponseProps{
											Description: "Server Side Error",
										},
									},
								},
							},
						},
					},
				},
			},
		}

		paths.Paths["/"+apiSchema.Path+"/{id}"] = pathItem
		paths.Paths["/"+apiSchema.PluralPath] = pluralPathItem

		definitions[apiSchema.JSONSchema.GoName+"APIListType"] = spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					apiSchema.PluralPath: {
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray([]string{"array"}),
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref,
									},
								},
							},
						},
					},
				},
			},
		}
		definitions[apiSchema.JSONSchema.GoName+"APISingleType"] = spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					apiSchema.Path: {
						SchemaProps: spec.SchemaProps{
							Ref: ref,
						},
					},
				},
			},
		}
	}
	for _, definitionSchema := range api.Definitions {
		for path, definition := range definitionSchema.Definitions {
			d, err := definition.ToOpenAPI()
			if err != nil {
				return nil, err
			}
			definitions[path] = *d
		}
	}
	return openAPI, nil
}

//ToOpenAPI translate json schema to OpenAPI format.
func (s *JSONSchema) ToOpenAPI() (*spec.Schema, error) {
	if s == nil {
		return nil, nil
	}
	refType := s.getRefType()
	if refType != "" {
		ref, err := spec.NewRef("#/definitions/" + refType)
		if err != nil {
			return nil, err
		}
		return &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Ref:         ref,
				Description: s.Description,
			},
		}, nil
	}
	//items
	items, err := s.Items.ToOpenAPI()
	if err != nil {
		return nil, err
	}
	//properties
	properties := map[string]spec.Schema{}
	for key, property := range s.Properties {
		var p *spec.Schema
		p, err = property.ToOpenAPI()
		if err != nil {
			return nil, err
		}
		properties[key] = *p
	}
	return &spec.Schema{
		SchemaProps: spec.SchemaProps{
			ID:          s.ID,
			Description: s.Description,
			Type:        spec.StringOrArray([]string{s.Type}),
			Title:       s.Title,
			//TODO(nati) support this.
			//Format: s.Format,
			//Maximum: s.Maximum,
			//Minimum: s.Minimum,
			//Pattern: s.Pattern,
			//Enum: s.Enum,
			Default:  s.Default,
			Required: s.Required,
			Items: &spec.SchemaOrArray{
				Schema: items,
			},
			Properties: properties,
		},
	}, nil
}
