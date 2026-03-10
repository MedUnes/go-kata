package myerror

import (
	"errors"
	"fmt"
	"testing"
)

func TestDoThingMatrix(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(bool) error
		fail     bool
		wantNil  bool
		wantType string
	}{
		{
			name:     "BuggyDoThing returns typed nil (should NOT be nil)",
			fn:       BuggyDoThing,
			fail:     false,
			wantNil:  false,
			wantType: "*myerror.MyError",
		},
		{
			name:     "BuggyDoThing returns error (fail=true)",
			fn:       BuggyDoThing,
			fail:     true,
			wantNil:  false,
			wantType: "*myerror.MyError",
		},
		{
			name:     "FixedDoThing returns nil (should be nil)",
			fn:       FixedDoThing,
			fail:     false,
			wantNil:  true,
			wantType: "<nil>",
		},
		{
			name:     "FixedDoThing returns error (fail=true)",
			fn:       FixedDoThing,
			fail:     true,
			wantNil:  false,
			wantType: "*myerror.MyError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn(tt.fail)
			gotNil := err == nil
			gotType := "<nil>"
			if err != nil {
				gotType = fmt.Sprintf("%T", err)
			}
			if gotNil != tt.wantNil {
				t.Errorf("expected nil: %v, got: %v (type: %s, value: %#v)", tt.wantNil, gotNil, gotType, err)
			}
			if gotType != tt.wantType {
				t.Errorf("expected type: %s, got: %s", tt.wantType, gotType)
			}
		})
	}
}

func TestErrorsAsExtractionMatrix(t *testing.T) {
	tests := []struct {
		name     string
		wrap     bool
		expectOp string
	}{
		{
			name:     "Direct MyError",
			wrap:     false,
			expectOp: "extract",
		},
		{
			name:     "Wrapped MyError",
			wrap:     true,
			expectOp: "extract",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orig := &MyError{Op: "extract"}
			var err error = orig
			if tt.wrap {
				err = WrapError(err)
			}
			var me *MyError
			if !errors.As(err, &me) {
				t.Fatalf("errors.As failed to extract *MyError")
			}
			if me == nil || me.Op != tt.expectOp {
				t.Fatalf("extracted wrong value: %#v", me)
			}
		})
	}
}
