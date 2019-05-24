package mappingTypes

import (
	"fmt"
)

func MapType(grpcType string, repeated bool) string {
	if grpcType == "" {
		return ""
	}
	typescriptType := mapType(grpcType)
	if repeated {
		return fmt.Sprintf("Array<%s>", typescriptType)
	}
	return typescriptType
}

func mapType(grpcType string) string {
	m := map[string]string{
		"string":  "string",
		"bool":    "boolean",
		"int64":   "string",
		"int32":   "number",
		"uint32":  "number",
		"uint64":  "string",
		"float32": "number",
	}
	out, ok := m[grpcType]
	if !ok {
		return grpcType
	}
	return out
}

// TODO Map type should be a thing in this
// 		Example
// 		export interface CompanyComposition {
//  	// companies_weight is a map of company id to their weight, i.e. their proportion of the portfolio
//  	companies_weight: CompanyWeightMap
//  	unallocated_weight: number
//		}
//		export type CompanyWeightMap = {
//  	[key: string]: number
//		}
