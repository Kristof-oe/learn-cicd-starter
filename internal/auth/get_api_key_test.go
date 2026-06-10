package auth

import (
	"net/http"
	"testing"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		a, b        string
		expected    string
		expectError bool
	}{
		{a: "Authorization", b: "ApiKey seged", expected: "seged"},
		{a: "", b: "", expected: "", expectError: true},
		{a: "Authorization", b: "Bearer seged3", expected: "", expectError: true},
	}
	for _, i := range tests {
		header := http.Header{}

		if i.a != "" {
			header.Add(i.a, i.b)
		}

		result, err := GetAPIKey(header)

		if i.expectError && err != nil {
			continue
		} else if i.expectError && err == nil {
			t.Error(ErrNoAuthHeaderIncluded)
		} else if !i.expectError && err != nil {
			t.Errorf("Unexpted error")
		} else if !i.expectError && result != i.expected {
			t.Errorf("The %s and %s should give back %s, not %s", i.a, i.b, result, i.expected)
		}
	}
}
