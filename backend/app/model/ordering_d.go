package model

import (
	"angrymiao-ai/app/ent"
	"angrymiao-ai/pkg"
)

const (
	descFlag = "-"
	idField  = "id"
)

func GenerateOrderingsFromStrings(fields []string) []ent.OrderFunc {
	orderFuncs := make([]ent.OrderFunc, 0)
	for _, v := range fields {
		orderFuncs = append(orderFuncs, ParseOrderingFunc(v))
	}
	return orderFuncs
}

func ParseOrderingFunc(field string) ent.OrderFunc {
	var orderFunc ent.OrderFunc

	switch {
	case field == "":
		orderFunc = ent.Desc(idField)
	case pkg.StartsWith(field, descFlag):
		orderFunc = ent.Desc(field[1:])
	default:
		orderFunc = ent.Asc(field)
	}

	return orderFunc
}
