package model

import (
	"angrymiao-ai/app/ent"
	pkgString "angrymiao-ai/pkg/string"
)

const (
	signDesc = "-"
	fieldID  = "id"
)

func GenerateOrderingFuncs(ordering []string) []ent.OrderFunc {
	orderingFunc := parseOrderingFuncsFromStrings(ordering)

	if len(orderingFunc) == 0 {
		orderingFunc = []ent.OrderFunc{ent.Desc(fieldID)}
	}

	return orderingFunc
}

func GenerateOrderingFuncFromString(field string) []ent.OrderFunc {
	orderFuncs := make([]ent.OrderFunc, 0)
	if field == "" {
		orderFuncs = append(orderFuncs, ent.Desc(fieldID))
	} else {
		orderFuncs = append(orderFuncs, parseOrderingFunc(field))
	}

	return orderFuncs
}

func parseOrderingFuncsFromStrings(fields []string) []ent.OrderFunc {
	orderFuncs := make([]ent.OrderFunc, 0)
	for _, v := range fields {
		if v != "" {
			orderFuncs = append(orderFuncs, parseOrderingFunc(v))
		}
	}

	return orderFuncs
}

func parseOrderingFunc(field string) ent.OrderFunc {
	var orderFunc ent.OrderFunc

	if pkgString.StartsWith(field, signDesc) {
		orderFunc = ent.Desc(field[1:])
	} else {
		orderFunc = ent.Asc(field)
	}

	return orderFunc
}
