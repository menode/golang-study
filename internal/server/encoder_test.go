package server

import (
	"encoding/json"
	"fmt"
	"kratos-test/internal/es"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPStruct(t *testing.T) {
	a := &es.HTTPError{
		Errors: make(map[string][]string),
	}
	a.Errors["body"] = []string{"can't be blank"}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	fmt.Printf("%s", string(b))
	panic("test failed")
}
