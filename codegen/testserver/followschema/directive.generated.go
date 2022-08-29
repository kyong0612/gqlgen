// Code generated by github.com/kyong0612/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/kyong0612/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) dir_length_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["min"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("min"))
		arg0, err = ec.unmarshalNInt2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["min"] = arg0
	var arg1 *int
	if tmp, ok := rawArgs["max"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("max"))
		arg1, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["max"] = arg1
	var arg2 *string
	if tmp, ok := rawArgs["message"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("message"))
		arg2, err = ec.unmarshalOString2ᚖstring(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["message"] = arg2
	return args, nil
}

func (ec *executionContext) dir_logged_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNUUID2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) dir_order1_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["location"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("location"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["location"] = arg0
	return args, nil
}

func (ec *executionContext) dir_order2_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["location"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("location"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["location"] = arg0
	return args, nil
}

func (ec *executionContext) dir_range_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *int
	if tmp, ok := rawArgs["min"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("min"))
		arg0, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["min"] = arg0
	var arg1 *int
	if tmp, ok := rawArgs["max"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("max"))
		arg1, err = ec.unmarshalOInt2ᚖint(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["max"] = arg1
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

func (ec *executionContext) _fieldMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver) interface{} {
	fc := graphql.GetFieldContext(ctx)
	for _, d := range fc.Field.Directives {
		switch d.Name {
		case "logged":
			rawArgs := d.ArgumentMap(ec.Variables)
			args, err := ec.dir_logged_args(ctx, rawArgs)
			if err != nil {
				ec.Error(ctx, err)
				return nil
			}
			n := next
			next = func(ctx context.Context) (interface{}, error) {
				if ec.directives.Logged == nil {
					return nil, errors.New("directive logged is not implemented")
				}
				return ec.directives.Logged(ctx, obj, n, args["id"].(string))
			}
		}
	}
	res, err := ec.ResolverMiddleware(ctx, next)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	return res
}

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _ObjectDirectives_text(ctx context.Context, field graphql.CollectedField, obj *ObjectDirectives) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ObjectDirectives_text(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return obj.Text, nil
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			min, err := ec.unmarshalNInt2int(ctx, 0)
			if err != nil {
				return nil, err
			}
			max, err := ec.unmarshalOInt2ᚖint(ctx, 7)
			if err != nil {
				return nil, err
			}
			message, err := ec.unmarshalOString2ᚖstring(ctx, "not valid")
			if err != nil {
				return nil, err
			}
			if ec.directives.Length == nil {
				return nil, errors.New("directive length is not implemented")
			}
			return ec.directives.Length(ctx, obj, directive0, min, max, message)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(string); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be string`, tmp)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ObjectDirectives_text(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ObjectDirectives",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ObjectDirectives_nullableText(ctx context.Context, field graphql.CollectedField, obj *ObjectDirectives) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ObjectDirectives_nullableText(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return obj.NullableText, nil
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.ToNull == nil {
				return nil, errors.New("directive toNull is not implemented")
			}
			return ec.directives.ToNull(ctx, obj, directive0)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*string); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *string`, tmp)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ObjectDirectives_nullableText(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ObjectDirectives",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ObjectDirectives_order(ctx context.Context, field graphql.CollectedField, obj *ObjectDirectives) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ObjectDirectives_order(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Order, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	fc.Result = res
	return ec.marshalNString2ᚕstringᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ObjectDirectives_order(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ObjectDirectives",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ObjectDirectivesWithCustomGoModel_nullableText(ctx context.Context, field graphql.CollectedField, obj *ObjectDirectivesWithCustomGoModel) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ObjectDirectivesWithCustomGoModel_nullableText(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return obj.NullableText, nil
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			if ec.directives.ToNull == nil {
				return nil, errors.New("directive toNull is not implemented")
			}
			return ec.directives.ToNull(ctx, obj, directive0)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(string); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be string`, tmp)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ObjectDirectivesWithCustomGoModel_nullableText(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ObjectDirectivesWithCustomGoModel",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputInnerDirectives(ctx context.Context, obj interface{}) (InnerDirectives, error) {
	var it InnerDirectives
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"message"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "message":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("message"))
			directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalNString2string(ctx, v) }
			directive1 := func(ctx context.Context) (interface{}, error) {
				min, err := ec.unmarshalNInt2int(ctx, 1)
				if err != nil {
					return nil, err
				}
				message, err := ec.unmarshalOString2ᚖstring(ctx, "not valid")
				if err != nil {
					return nil, err
				}
				if ec.directives.Length == nil {
					return nil, errors.New("directive length is not implemented")
				}
				return ec.directives.Length(ctx, obj, directive0, min, nil, message)
			}

			tmp, err := directive1(ctx)
			if err != nil {
				return it, graphql.ErrorOnPath(ctx, err)
			}
			if data, ok := tmp.(string); ok {
				it.Message = data
			} else {
				err := fmt.Errorf(`unexpected type %T from directive, should be string`, tmp)
				return it, graphql.ErrorOnPath(ctx, err)
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputInputDirectives(ctx context.Context, obj interface{}) (InputDirectives, error) {
	var it InputDirectives
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"text", "nullableText", "inner", "innerNullable", "thirdParty"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "text":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("text"))
			directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalNString2string(ctx, v) }
			directive1 := func(ctx context.Context) (interface{}, error) {
				if ec.directives.Directive3 == nil {
					return nil, errors.New("directive directive3 is not implemented")
				}
				return ec.directives.Directive3(ctx, obj, directive0)
			}
			directive2 := func(ctx context.Context) (interface{}, error) {
				min, err := ec.unmarshalNInt2int(ctx, 0)
				if err != nil {
					return nil, err
				}
				max, err := ec.unmarshalOInt2ᚖint(ctx, 7)
				if err != nil {
					return nil, err
				}
				message, err := ec.unmarshalOString2ᚖstring(ctx, "not valid")
				if err != nil {
					return nil, err
				}
				if ec.directives.Length == nil {
					return nil, errors.New("directive length is not implemented")
				}
				return ec.directives.Length(ctx, obj, directive1, min, max, message)
			}

			tmp, err := directive2(ctx)
			if err != nil {
				return it, graphql.ErrorOnPath(ctx, err)
			}
			if data, ok := tmp.(string); ok {
				it.Text = data
			} else {
				err := fmt.Errorf(`unexpected type %T from directive, should be string`, tmp)
				return it, graphql.ErrorOnPath(ctx, err)
			}
		case "nullableText":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("nullableText"))
			directive0 := func(ctx context.Context) (interface{}, error) { return ec.unmarshalOString2ᚖstring(ctx, v) }
			directive1 := func(ctx context.Context) (interface{}, error) {
				if ec.directives.Directive3 == nil {
					return nil, errors.New("directive directive3 is not implemented")
				}
				return ec.directives.Directive3(ctx, obj, directive0)
			}
			directive2 := func(ctx context.Context) (interface{}, error) {
				if ec.directives.ToNull == nil {
					return nil, errors.New("directive toNull is not implemented")
				}
				return ec.directives.ToNull(ctx, obj, directive1)
			}

			tmp, err := directive2(ctx)
			if err != nil {
				return it, graphql.ErrorOnPath(ctx, err)
			}
			if data, ok := tmp.(*string); ok {
				it.NullableText = data
			} else if tmp == nil {
				it.NullableText = nil
			} else {
				err := fmt.Errorf(`unexpected type %T from directive, should be *string`, tmp)
				return it, graphql.ErrorOnPath(ctx, err)
			}
		case "inner":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("inner"))
			directive0 := func(ctx context.Context) (interface{}, error) {
				return ec.unmarshalNInnerDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐInnerDirectives(ctx, v)
			}
			directive1 := func(ctx context.Context) (interface{}, error) {
				if ec.directives.Directive3 == nil {
					return nil, errors.New("directive directive3 is not implemented")
				}
				return ec.directives.Directive3(ctx, obj, directive0)
			}

			tmp, err := directive1(ctx)
			if err != nil {
				return it, graphql.ErrorOnPath(ctx, err)
			}
			if data, ok := tmp.(*InnerDirectives); ok {
				it.Inner = data
			} else if tmp == nil {
				it.Inner = nil
			} else {
				err := fmt.Errorf(`unexpected type %T from directive, should be *github.com/kyong0612/gqlgen/codegen/testserver/followschema.InnerDirectives`, tmp)
				return it, graphql.ErrorOnPath(ctx, err)
			}
		case "innerNullable":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("innerNullable"))
			directive0 := func(ctx context.Context) (interface{}, error) {
				return ec.unmarshalOInnerDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐInnerDirectives(ctx, v)
			}
			directive1 := func(ctx context.Context) (interface{}, error) {
				if ec.directives.Directive3 == nil {
					return nil, errors.New("directive directive3 is not implemented")
				}
				return ec.directives.Directive3(ctx, obj, directive0)
			}

			tmp, err := directive1(ctx)
			if err != nil {
				return it, graphql.ErrorOnPath(ctx, err)
			}
			if data, ok := tmp.(*InnerDirectives); ok {
				it.InnerNullable = data
			} else if tmp == nil {
				it.InnerNullable = nil
			} else {
				err := fmt.Errorf(`unexpected type %T from directive, should be *github.com/kyong0612/gqlgen/codegen/testserver/followschema.InnerDirectives`, tmp)
				return it, graphql.ErrorOnPath(ctx, err)
			}
		case "thirdParty":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("thirdParty"))
			directive0 := func(ctx context.Context) (interface{}, error) {
				return ec.unmarshalOThirdParty2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐThirdParty(ctx, v)
			}
			directive1 := func(ctx context.Context) (interface{}, error) {
				if ec.directives.Directive3 == nil {
					return nil, errors.New("directive directive3 is not implemented")
				}
				return ec.directives.Directive3(ctx, obj, directive0)
			}
			directive2 := func(ctx context.Context) (interface{}, error) {
				min, err := ec.unmarshalNInt2int(ctx, 0)
				if err != nil {
					return nil, err
				}
				max, err := ec.unmarshalOInt2ᚖint(ctx, 7)
				if err != nil {
					return nil, err
				}
				if ec.directives.Length == nil {
					return nil, errors.New("directive length is not implemented")
				}
				return ec.directives.Length(ctx, obj, directive1, min, max, nil)
			}

			tmp, err := directive2(ctx)
			if err != nil {
				return it, graphql.ErrorOnPath(ctx, err)
			}
			if data, ok := tmp.(*ThirdParty); ok {
				it.ThirdParty = data
			} else if tmp == nil {
				it.ThirdParty = nil
			} else {
				err := fmt.Errorf(`unexpected type %T from directive, should be *github.com/kyong0612/gqlgen/codegen/testserver/followschema.ThirdParty`, tmp)
				return it, graphql.ErrorOnPath(ctx, err)
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var objectDirectivesImplementors = []string{"ObjectDirectives"}

func (ec *executionContext) _ObjectDirectives(ctx context.Context, sel ast.SelectionSet, obj *ObjectDirectives) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, objectDirectivesImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ObjectDirectives")
		case "text":

			out.Values[i] = ec._ObjectDirectives_text(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "nullableText":

			out.Values[i] = ec._ObjectDirectives_nullableText(ctx, field, obj)

		case "order":

			out.Values[i] = ec._ObjectDirectives_order(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var objectDirectivesWithCustomGoModelImplementors = []string{"ObjectDirectivesWithCustomGoModel"}

func (ec *executionContext) _ObjectDirectivesWithCustomGoModel(ctx context.Context, sel ast.SelectionSet, obj *ObjectDirectivesWithCustomGoModel) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, objectDirectivesWithCustomGoModelImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ObjectDirectivesWithCustomGoModel")
		case "nullableText":

			out.Values[i] = ec._ObjectDirectivesWithCustomGoModel_nullableText(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNInnerDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐInnerDirectives(ctx context.Context, v interface{}) (*InnerDirectives, error) {
	res, err := ec.unmarshalInputInnerDirectives(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNInputDirectives2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐInputDirectives(ctx context.Context, v interface{}) (InputDirectives, error) {
	res, err := ec.unmarshalInputInputDirectives(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOInnerDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐInnerDirectives(ctx context.Context, v interface{}) (*InnerDirectives, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputInnerDirectives(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOInputDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐInputDirectives(ctx context.Context, v interface{}) (*InputDirectives, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputInputDirectives(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOObjectDirectives2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐObjectDirectives(ctx context.Context, sel ast.SelectionSet, v *ObjectDirectives) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._ObjectDirectives(ctx, sel, v)
}

func (ec *executionContext) marshalOObjectDirectivesWithCustomGoModel2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐObjectDirectivesWithCustomGoModel(ctx context.Context, sel ast.SelectionSet, v *ObjectDirectivesWithCustomGoModel) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._ObjectDirectivesWithCustomGoModel(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
