package gotracer

// New returns a pointer to a Tracer that is able to collect messages and display them in the order that it got them.
func New() *Tracer {
	return &Tracer{
		collectedSlugs: make([]string, 0),
	}
}

// Tracer is able to collect messages and returns them back to you. Useful in testing and debugging.
type Tracer struct {
	collectedSlugs []string
}

// Send satisfies the overrides.Overrider interface.
// It collects the messages and stores them in a slice to preseve order.
func (t *Tracer) Send(m string) {
	t.collectedSlugs = append(t.collectedSlugs, m)
}

// Len tells you how many messages it has collected.
func (t *Tracer) Len() int {
	return len(t.collectedSlugs)
}

// Show will return a copy of the messages that the tracer has collected.
func (t *Tracer) Show() []string {
	return t.collectedSlugs
}

// Reset will clear out any collected messages effectively resetting the collection.
func (t *Tracer) Reset() {
	t.collectedSlugs = make([]string, 0)
}
