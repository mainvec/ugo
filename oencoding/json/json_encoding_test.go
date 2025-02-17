package json_test

import (
	"testing"

	"github.com/mainvec/ugo/oencoding"
	_ "github.com/mainvec/ugo/oencoding/json"
)

type TestObject struct {
	Name string
	Age  int
}

func TestJSONEncoding(t *testing.T) {
	enc, ok := oencoding.LookupEncoding("json")
	if !ok {
		t.Errorf("Expected to find json encoding")
	}
	obj := TestObject{Name: "John", Age: 25}
	data, err := enc.Encode(obj)
	if err != nil {
		t.Errorf("Error encoding object: %v", err)
	}
	var obj2 TestObject
	err = enc.Decode(data, &obj2)
	if err != nil {
		t.Errorf("Error decoding object: %v", err)
	}
	if obj != obj2 {
		t.Errorf("Expected %v, got %v", obj, obj2)
	}
	if enc.MimeType() != "application/json" {
		t.Errorf("Expected application/json, got %s", enc.MimeType())
	}
}
