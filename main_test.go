package main

import (
	_ "embed"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_urlParam(t *testing.T) {

	r := httptest.NewRequest("GET", "/?a=123", nil)
	check := urlParam(r, "a", "fb")
	assert.Equal(t, check, "123")
	check2 := urlParam(r, "a", 123)
	assert.Equal(t, check2, 123)

}
