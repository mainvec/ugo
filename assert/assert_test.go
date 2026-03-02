package assert

import "testing"

func TestAssert_WithTrueAssertion_DoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Assert panicked unexpectedly: %v", r)
		}
	}()

	Assert(true, "This should not panic")
}

func TestAssert_WithFalseAssertion_Panics(t *testing.T) {
	expectedMessage := "Assertion failed"

	defer func() {
		if r := recover(); r == nil {
			t.Error("Assert did not panic as expected")
		} else if r != expectedMessage {
			t.Errorf("Expected panic message '%s', got '%v'", expectedMessage, r)
		}
	}()

	Assert(false, expectedMessage)
}

func TestAssert_WithFalseAssertion_PanicsWithCorrectMessage(t *testing.T) {
	testCases := []struct {
		name            string
		assertion       bool
		expectedMessage string
	}{
		{
			name:            "custom error message 1",
			assertion:       false,
			expectedMessage: "Value cannot be nil",
		},
		{
			name:            "custom error message 2",
			assertion:       false,
			expectedMessage: "Index out of bounds",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("Assert did not panic as expected")
				} else if r != tc.expectedMessage {
					t.Errorf("Expected panic message '%s', got '%v'", tc.expectedMessage, r)
				}
			}()

			Assert(tc.assertion, tc.expectedMessage)
		})
	}
}
