package mock

import (
	"context"
	"testing"
)

func Test_mock_Value(t *testing.T) {
	s := New(context.Background())

	got, err := s.Value()

	if err != nil {
		t.Errorf("Value() failed: %v", err)
	}

	if got == nil {
		t.Errorf("Value() failed: %v", err)
	}

	f, ok := got.(float64)

	if !ok {
		t.Errorf("Value() failed to cast")
	}

	if f < 0 {
		t.Errorf("Value() failed: zeroed value")
	}

	t.Logf("value=%f", f)
}
