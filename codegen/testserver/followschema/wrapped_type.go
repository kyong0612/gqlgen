package followschema

import "github.com/kyong0612/gqlgen/codegen/testserver/followschema/otherpkg"

type (
	WrappedScalar = otherpkg.Scalar
	WrappedStruct otherpkg.Struct
	WrappedMap    otherpkg.Map
	WrappedSlice  otherpkg.Slice
)
