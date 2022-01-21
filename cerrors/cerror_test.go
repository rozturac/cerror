package cerrors

import (
	"errors"
	"github.com/rozturac/cerror"
	"testing"
)

func TestIs_WhenSourceIsCustomErrorOnly(t *testing.T) {
	sourceError := cerror.NullReferenceError("test")
	targetError := errors.New("test")
	want := false

	got := Is(sourceError, targetError)

	if want != got {
		t.Errorf("Is(sourceError, targetError)=%v, expected: %v", got, want)
	}
}

func TestIs_WhenTargetIsCustomErrorOnly(t *testing.T) {
	sourceError := errors.New("test")
	targetError := cerror.NullReferenceError("test")
	want := false

	got := Is(sourceError, targetError)

	if want != got {
		t.Errorf("Is(sourceError, targetError)=%v, expected: %v", got, want)
	}
}

func TestIs_WhenTargetAndSourceIsCustomError(t *testing.T) {
	sourceError := cerror.NullReferenceError("test")
	targetError := cerror.NullReferenceError("test")
	want := true

	got := Is(sourceError, targetError)

	if want != got {
		t.Errorf("Is(sourceError, targetError)=%v, expected: %v", got, want)
	}
}

func TestIs_WhenTargetAndSourceIsNotCustomError(t *testing.T) {
	err := errors.New("test error")
	sourceError := err
	targetError := err
	want := true

	got := Is(sourceError, targetError)

	if want != got {
		t.Errorf("Is(sourceError, targetError)=%v, expected: %v", got, want)
	}
}

func TestIs_WhenTargetAndSourceIsNil(t *testing.T) {
	var (
		sourceError error = nil
		targetError error = nil
		want        bool  = true
	)

	got := Is(sourceError, targetError)

	if want != got {
		t.Errorf("Is(sourceError, targetError)=%v, expected: %v", got, want)
	}
}

func TestIs_WhenSourceIsNilOnly(t *testing.T) {
	var (
		sourceError error = nil
		targetError error = errors.New("test")
		want        bool  = false
	)

	got := Is(sourceError, targetError)

	if want != got {
		t.Errorf("Is(sourceError, targetError)=%v, expected: %v", got, want)
	}
}

func TestIs_WhenTargetIsNilOnly(t *testing.T) {
	var (
		sourceError error = errors.New("test")
		targetError error = nil
		want        bool  = false
	)

	got := Is(sourceError, targetError)

	if want != got {
		t.Errorf("Is(sourceError, targetError)=%v, expected: %v", got, want)
	}
}
