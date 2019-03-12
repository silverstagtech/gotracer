package gotracer

import "testing"

func TestLen(t *testing.T) {
	tracing := New()
	tracing.Send("test")

	if tracing.Len() != 1 {
		t.Log("Tracing should have only 1 message on it.")
		t.Fail()
	}
}

func TestReset(t *testing.T) {
	tracing := New()
	tracing.Send("Test")
	tracing.Reset()

	if tracing.Len() != 0 {
		t.Log("Tracing should have no values to return once reset but Len is not zero")
		t.Fail()
	}
}

func TestShow(t *testing.T) {
	tracing := New()
	tracing.Send("One")
	tracing.Send("Two")

	if tracing.Show()[0] != "One" {
		t.Log("Show did not give back the correct message that was sent first.")
		t.Fail()
	}
}
