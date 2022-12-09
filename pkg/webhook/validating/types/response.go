package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type ValidatingResponse struct {
	Allowed  bool     `json:"allowed"`
	Message  string   `json:"message,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	Patch    []byte   `json:"patch,omitempty"`
}

func ValidatingResponseFromFile(filePath string) (*ValidatingResponse, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read %s: %s", filePath, err)
	}

	if len(data) == 0 {
		return nil, nil
	}
	return ValidatingResponseFromBytes(data)
}

func ValidatingResponseFromBytes(data []byte) (*ValidatingResponse, error) {
	return ValidatingResponseFromReader(bytes.NewReader(data))
}

func ValidatingResponseFromReader(r io.Reader) (*ValidatingResponse, error) {
	response := new(ValidatingResponse)

	dec := json.NewDecoder(r)

	err := dec.Decode(response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *ValidatingResponse) Dump() string {
	b := new(strings.Builder)
	b.WriteString("ValidatingResponse(allowed=")
	b.WriteString(strconv.FormatBool(r.Allowed))
	if r.Message != "" {
		b.WriteString(",")
		b.WriteString(r.Message)
	}
	for _, warning := range r.Warnings {
		b.WriteString(",")
		b.WriteString(warning)
	}
	b.WriteString(")")
	return b.String()
}
