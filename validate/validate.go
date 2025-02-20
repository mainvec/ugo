package validate

import (
	"cmp"
	"errors"
	"fmt"
	"regexp"
)

var _ ValidationRule = &DefaultValidationRule{}
var (
	Blank = &DefaultValidationRule{ruleName: "Blank",
		ruleErrorMsg: "should be blank",
		ruleFunc: func(value any) (bool, error) {
			r := IsBlank(value)
			return r, nil
		}}

	NotBlank = &DefaultValidationRule{ruleName: "NotBlank",
		ruleErrorMsg: "should not be blank",
		ruleFunc: func(value any) (bool, error) {
			r := IsBlank(value)
			return !r, nil
		}}
)

type Validator struct {
}

type ValidationBucket struct {
	result *DefaultValidationResult
}

func NewValidator() *Validator {
	return &Validator{}
}

func NewBucket() *ValidationBucket {
	return &ValidationBucket{
		result: &DefaultValidationResult{},
	}
}

func (b *ValidationBucket) Validate(propName string, value any, vRule ...ValidationRule) (bool, error) {
	v := true
	for _, r := range vRule {
		valid, err := r.RuleFunc()(value)
		if err != nil {
			return false, err
		}
		if !valid {
			v = false
			b.result.AddError(propName, r.RuleErrorMsg())
		}
	}
	return v, nil
}

func (b *ValidationBucket) Result() *DefaultValidationResult {
	return b.result
}
func (b *ValidationBucket) Error() error {
	if b.result.Valid() {
		return nil
	}
	errs := make([]error, b.ErrorCount())
	for i, e := range b.Result().ValidationErrors() {
		errs[i] = e
	}
	return errors.Join(errs...)
}

func (b *ValidationBucket) ErrorCount() int {
	return len(b.result.verrors)
}

func (b *ValidationBucket) IsValid() bool {
	return b.result.Valid()
}

type ValidationFunc func(value any) (bool, error)

type ValidationRule interface {
	RuleName() string
	RuleFunc() ValidationFunc
	RuleErrorMsg() string
}

type DefaultValidationRule struct {
	ruleName     string
	ruleErrorMsg string
	ruleFunc     ValidationFunc
}

func (r *DefaultValidationRule) RuleName() string {
	return r.ruleName
}

func (r *DefaultValidationRule) RuleFunc() ValidationFunc {
	return r.ruleFunc
}

func (r *DefaultValidationRule) RuleErrorMsg() string {
	return r.ruleErrorMsg
}

func NewValidationRule(ruleName, ruleErrorMsg string, ruleFunc ValidationFunc) ValidationRule {
	return &DefaultValidationRule{
		ruleName:     ruleName,
		ruleErrorMsg: ruleErrorMsg,
		ruleFunc:     ruleFunc,
	}
}

// Check is string len > 0. Accepts string or string pointer.
func IsBlank(v any) bool {
	if v == nil {
		return true
	}
	switch v := v.(type) {
	case string:
		return len(v) == 0
	case *string:
		return len(*v) == 0
	default:
		//no validation for other types. assume valid
		return true
	}
}

func Validate(value any, vRule ...ValidationRule) (bool, error) {

	for _, r := range vRule {
		valid, err := r.RuleFunc()(value)
		if err != nil {
			return false, err
		}
		if !valid {
			return false, nil
		}
	}
	return true, nil
}

type ValidationResult interface {
	Valid() bool
	ValidationErrors() []ValidationError
}

type ValidationError interface {
	String() string
	Error() string
}
type DefaultValidationResult struct {
	verrors []DefaultValidationError
}

type DefaultValidationError struct {
	fieldName       string
	validationError string
}

func (v *DefaultValidationResult) AddError(fieldName, errormsg string) *DefaultValidationResult {
	verr := DefaultValidationError{
		fieldName:       fieldName,
		validationError: errormsg,
	}
	v.verrors = append(v.verrors, verr)
	return v
}

func (v *DefaultValidationError) String() string {
	return fmt.Sprintf("invalid [%v], [%v]", v.fieldName, v.validationError)
}
func (v *DefaultValidationError) Error() string {
	return v.String()
}

func (v *DefaultValidationResult) Valid() bool {
	return len(v.verrors) == 0
}

func (v *DefaultValidationResult) ValidationErrors() []ValidationError {

	resultErrors := make([]ValidationError, len(v.verrors))
	for index := range v.verrors {
		resultErrors[index] = &v.verrors[index]
	}
	return resultErrors
}

func OneOfRule(values ...any) ValidationRule {
	return &DefaultValidationRule{
		ruleName:     "OneOf",
		ruleErrorMsg: "should be one of " + fmt.Sprintf("%v", values),
		ruleFunc: func(value any) (bool, error) {
			for _, v := range values {
				if value == v {
					return true, nil
				}
			}
			return false, nil
		},
	}
}

func RegExRule(pattern string) ValidationRule {
	regex := regexp.MustCompile(pattern)
	return &DefaultValidationRule{
		ruleName:     "RegEx",
		ruleErrorMsg: "should match pattern " + pattern,
		ruleFunc: func(value any) (bool, error) {
			if value == nil {
				return true, nil
			}
			switch value := value.(type) {
			case string:
				return regex.MatchString(value), nil
			case *string:
				return regex.MatchString(*value), nil
			case []byte:
				return regex.Match(value), nil
			case *[]byte:
				return regex.Match(*value), nil
			}
			//no validation for other types. assume valid
			return true, nil
		},
	}
}

// Range returns a validation rule that checks if the value is within the range.
func Range[K cmp.Ordered](from K, to K) ValidationRule {
	return &DefaultValidationRule{
		ruleName:     "Range",
		ruleErrorMsg: fmt.Sprintf("should be in range %v to %v", from, to),
		ruleFunc: func(value any) (bool, error) {
			if value == nil {
				return true, nil
			}
			switch value := value.(type) {
			case K:
				return value >= from && value <= to, nil
			case *K:
				return *value >= from && *value <= to, nil
			}
			//no validation for other types. assume valid
			return true, nil
		},
	}
}
