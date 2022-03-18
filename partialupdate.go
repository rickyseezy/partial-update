package partialupdate

import (
	"fmt"
	"strings"
)

// UpdateRequest holds the fields and their value
type UpdateRequest map[string]interface{}

// AllowFields a simple map of allow fields
type AllowFields map[string]string

type Query struct {
	FieldsIndex string
	Args        []interface{}
}

type PartialUpdate struct {
	debug       bool
	allowFields AllowFields
	request     UpdateRequest
}

func NewPartialUpdate(af AllowFields, req UpdateRequest, printlog bool) *PartialUpdate {
	return &PartialUpdate{
		allowFields: af,
		request:     req,
		debug:       printlog,
	}
}

func (pu *PartialUpdate) BuildQuery() *Query {
	var fields []string
	var args []interface{}

	for k, v := range pu.request {
		if f, ok := pu.allowFields[k]; ok {
			fields = append(fields, f)
			args = append(args, v)
		}
	}

	fieldsString := strings.Join(fields, ", ")

	argsString := ""
	for i := range args {
		fieldIndex := i + 1
		if fieldIndex == len(args) {
			argsString += fmt.Sprintf("%s=$%d", fields[i], fieldIndex)
			continue
		}
		argsString += fmt.Sprintf("%s=$%d, ", fields[i], fieldIndex)
	}

	if pu.debug {
		for i, a := range args {
			fmt.Println(fmt.Sprintf("%d = %v", i+1, a))
		}
		fmt.Println("ARGUMENTS:")
		fmt.Println(args)

		fmt.Println("FIELDS TO UPDATE: ", fieldsString)
	}

	return &Query{
		FieldsIndex: argsString,
		Args:        args,
	}
}
