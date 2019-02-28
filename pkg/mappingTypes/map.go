package mappingTypes

import "fmt"

func MapType(grpcType string, repeated bool) string {
	typescriptType := mapType(grpcType)
	if repeated {
		return fmt.Sprintf("Array<%s>", typescriptType)
	}
	return typescriptType
}

func mapType(grpcType string) string {
	m := map[string]string{
		"string": "string",
		"bool":   "boolean",
		"int64":  "string",
		"int32":  "number",
	}
	out, ok := m[grpcType]
	if !ok {
		return grpcType
	}
	return out
}
