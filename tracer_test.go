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

func TestShowBytes(t *testing.T) {
	tracing := New()
	tracing.SendBytes([]byte("One"))
	tracing.SendByte('a')

	if tracing.ShowBytes()[0][0] != []byte("One")[0] {
		t.Log("Show did not give back the correct message.")
		t.Fail()
	}

	if tracing.ShowBytes()[1][0] != 'a' {
		t.Log("Show did not give back the correct message.")
		t.Fail()
	}
}

func TestShowInterface(t *testing.T) {
	tracing := New()
	tracing.SendInterface("String")
	tracing.SendInterface([]byte("ByteSlice"))
	tracing.SendInterface(3)
	tracing.SendInterface(map[string]string{"map": "string"})

	if tracing.ShowRaw()[2] != 3 {
		t.Log("Show did not give back the correct message.")
		t.Fail()
	}
}

func TestPrinters(t *testing.T) {
	// We can only look for panics here.
	tracing := New()
	tracing.SendInterface("String")
	tracing.SendInterface([]byte("ByteSlice"))
	tracing.SendInterface(3)
	tracing.SendInterface(map[string]string{"mapKey": "stringValue"})

	tracing.Println()
	tracing.PrintlnT(t)
}
