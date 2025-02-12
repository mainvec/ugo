package validate_test

import (
	"fmt"
	"testing"

	"github.com/bytestandard/bsugo/validate"
	"golang.org/x/exp/constraints"
)

func TestNewBucket(t *testing.T) {
	bucket := validate.NewBucket()
	if bucket == nil {
		t.Fatal("expoecting validator")
	}
	_, err := bucket.Validate("valid blank", "", validate.Blank)
	if err != nil {
		t.Fatal(err)
	}

	if !bucket.IsValid() {
		t.Fatalf("wanted valid, got not valid")
	}
	bucket.Validate("invalid blank", "not blank", validate.Blank)
	bucket.Validate("invalid not blank", "", validate.NotBlank)
	if bucket.IsValid() {
		t.Fatalf("wanted invalid, got valid")
	}
	if bucket.ErrorCount() != 2 {
		t.Fatalf("wanted 2 errors, got %v", bucket.ErrorCount())
	}
	msg := fmt.Sprintf("errors:%v", bucket.Result().ValidationErrors())
	fmt.Println(msg)

}

func TestOneOfRule(t *testing.T) {
	bucket := validate.NewBucket()
	bucket.Validate("valid oneof", "go", validate.OneOfRule("go", "js"))
	if !bucket.IsValid() {
		t.Fatalf("wanted valid, got not valid")
	}
	bucket.Validate("invalid oneof", "java", validate.OneOfRule("go", "js"))
	if bucket.IsValid() {
		t.Fatalf("wanted invalid, got valid")
	}
	if bucket.ErrorCount() != 1 {
		t.Fatalf("wanted 1 errors, got %v", bucket.ErrorCount())
	}
}

func TestRegExRule(t *testing.T) {
	bucket := validate.NewBucket()
	bucket.Validate("valid regex", "go", validate.RegExRule("go|js"))
	if !bucket.IsValid() {
		t.Fatalf("wanted valid, got not valid")
	}
	bucket.Validate("invalid regex", "java", validate.RegExRule("go"))
	if bucket.IsValid() {
		t.Fatalf("wanted invalid, got valid")
	}
	if bucket.ErrorCount() != 1 {
		t.Fatalf("wanted 1 errors, got %v", bucket.ErrorCount())
	}
}

func TestRange(t *testing.T) {
	type args[K constraints.Ordered] struct {
		from  K
		to    K
		value K
	}
	tests := []struct {
		name string
		args args[int]
		want bool
	}{

		{
			name: "valid range value",
			args: args[int]{from: 1, to: 10, value: 5},
			want: true,
		},

		{
			name: "invalid range value",
			args: args[int]{from: 1, to: 10, value: 11},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := validate.Range(tt.args.from, tt.args.to).RuleFunc()
			if got, _ := rule(tt.args.value); got != tt.want {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		})
	}
}
