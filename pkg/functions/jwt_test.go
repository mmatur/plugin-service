package functions

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_customValidation(t *testing.T) {
	handler := JWTHandler{
		claims: map[string]check{
			"test": {
				header: "foo",
				value:  "value",
			},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/foo", nil)

	mapClaims := jwt.MapClaims{
		"test": "value",
	}

	err := handler.customValidation(req, mapClaims)
	require.NoError(t, err)

	assert.Equal(t, "value", req.Header.Get("foo"))
}

func Test_customValidation_errors(t *testing.T) {
	testCases := []struct {
		desc      string
		claims    map[string]check
		mapClaims jwt.MapClaims
		expected  string
	}{
		{
			desc: "missing claim",
			claims: map[string]check{
				"missing": {
					header: "foo",
					value:  "value",
				},
			},
			mapClaims: jwt.MapClaims{
				"test": "value",
			},
			expected: "claims: invalid JWT",
		},
		{
			desc: "invalid claim value",
			claims: map[string]check{
				"test": {
					header: "foo",
					value:  "nope",
				},
			},
			mapClaims: jwt.MapClaims{
				"test": "value",
			},
			expected: "claims: invalid JWT",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			handler := JWTHandler{claims: test.claims}

			req := httptest.NewRequest(http.MethodGet, "/foo", nil)

			err := handler.customValidation(req, test.mapClaims)
			require.EqualError(t, err, test.expected)

			assert.Empty(t, req.Header.Get("foo"))
		})
	}
}
