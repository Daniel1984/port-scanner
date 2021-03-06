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

type errResp struct {
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

func (rv ReqValidator) ValidDecimalString(field, value string) {
	_, err := strconv.Atoi(value)
	if err != nil {
		rv.Errors.Add(field, "invalid decimal string")
	}
}

func (rv ReqValidator) Valid() bool {
	return len(rv.Errors) == 0
}

func (rv ReqValidator) GetErrResp() []byte {
	er := errResp{
		Result:        "ERROR",
		Cause:         "INVALID_REQUEST",
		InvalidFields: rv.Errors,
	}

	b, _ := json.Marshal(er)
	return b
}
