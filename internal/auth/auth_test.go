package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := http.Header{
		"Authorization": []string{"ApiKey 123"},
	}
	want := "123"
	got, err := GetAPIKey(h)
	if !reflect.DeepEqual(want, got) || err != nil {
		t.Fatalf("expected: %v, got: %v, error: %v", want, got, err)
	}
}
