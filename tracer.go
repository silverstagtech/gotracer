package gotracer

import (
	"fmt"
	"testing"
)

// Tracer is able to collect messages and returns them back to you. Useful in testing and debugging.
type Tracer struct {
	collectedSlugs []interface{}
}

// New returns a pointer to a Tracer that is able to collect messages and display them in the order that it got them.
func New() *Tracer {
	return &Tracer{
		collectedSlugs: make([]interface{}, 0),
	}
}

// SendInterface will allow you to send any type into the Tracer.
func (t *Tracer) SendInterface(m interface{}) {
	t.collectedSlugs = append(t.collectedSlugs, m)
}

// Send collects the messages and stores them in a slice to preseve order.
func (t *Tracer) Send(m string) {
	t.SendInterface(m)
}

// SendBytes will add a []byte to the tracer
func (t *Tracer) SendBytes(m []byte) {
	t.SendInterface(m)
}

// SendByte will add a byte to the tracer
func (t *Tracer) SendByte(m byte) {
	t.SendInterface(m)
}

// Len tells you how many messages it has collected.
func (t *Tracer) Len() int {
	return len(t.collectedSlugs)
}

// Show will return a copy of the messages that the tracer has collected.
// It is able to convert []string, []bytes and byte to strings and return a []strings
// with all the results. Its up to the user to make sure that the messages
// will convert into something readable.
func (t *Tracer) Show() []string {
	out := []string{}
	for _, v := range t.collectedSlugs {
		switch v.(type) {
		case string:
			out = append(out, v.(string))
		case []string:
			out = append(out, fmt.Sprintf("%s", v.([]string)))
		case byte:
			out = append(out, string(v.(byte)))
		case []byte:
			out = append(out, string(v.([]byte)))
		}
	}
	return out
}

// ShowBytes will return a copy of the messages that the tracer has collected.
// It is able to convert string, []string and byte to []byte and return a []byte
// with all the results. Its up to the user to make sure that the messages
// will convert into something readable.
func (t *Tracer) ShowBytes() [][]byte {
	out := [][]byte{}
	for _, v := range t.collectedSlugs {
		switch v.(type) {
		case string:
			out = append(out, []byte(v.(string)))
		case byte:
			out = append(out, []byte{v.(byte)})
		case []byte:
			out = append(out, v.([]byte))
		}
	}
	return out
}

// ShowRaw will return a copy of the collected tracer slugs.
// It is up to the user to type reference the data collected.
func (t *Tracer) ShowRaw() []interface{} {
	output := make([]interface{}, len(t.collectedSlugs))
	copy(output, t.collectedSlugs)
	return output
}

// Reset will clear out any collected messages effectively resetting the collection.
func (t *Tracer) Reset() {
	t.collectedSlugs = make([]interface{}, 0)
}

// Println will print out everything using fmt.Println
// Given that the types stored are unknown its up to the user to make sure that they
// are safe for printing. Generally intended for []string, string, []byte, byte and numbers,
// but can do others.
// When printing you will get the postion that the message was collected in, then the message.
func (t *Tracer) Println() {
	for index, data := range t.collectedSlugs {
		fmt.Println(index+1, data)
	}
}

// PrintlnT is like Println but will make use of the testing.T Logf function.
func (t *Tracer) PrintlnT(testingT *testing.T) {
	for index, data := range t.collectedSlugs {
		testingT.Logf("%d: %v", index+1, data)
	}
}
