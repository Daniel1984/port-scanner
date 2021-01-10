package reqvalidator

import (
	"encoding/json"
	"strconv"
	"strings"
)

type validationerrors map[string][]string

func (ve validationerrors) Add(field, message string) {
	ve[field] = append(ve[field], message)
}

func (ve validationerrors) Get(field string) string {
	errStr, ok := ve[field]
	if !ok {
		return ""
	}

	return errStr[0]
}

type RespErr struct {
	Result        string
	Cause         string
	InvalidFields validationerrors
}

type ReqValidator struct {
	Errors validationerrors
}

func New() ReqValidator {
	return ReqValidator{
		validationerrors(map[string][]string{}),
	}
}

func (rv ReqValidator) Required(field, value string) {
	if strings.TrimSpace(value) == "" {
		rv.Errors.Add(field, "can't be blank")
	}
}

func (rv ReqValidator) ValidStringInt(field, value string) {
	_, err := strconv.Atoi(value)
	if err != nil {
		rv.Errors.Add(field, "invalid string representation of int")
	}
}

func (rv ReqValidator) Valid() bool {
	return len(rv.Errors) == 0
}

func (rv ReqValidator) GetErrResp() []byte {
	re := RespErr{
		Result:        "ERROR",
		Cause:         "INVALID_REQUEST",
		InvalidFields: rv.Errors,
	}

	b, _ := json.Marshal(re)
	return b
}
