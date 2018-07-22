// Common / Shared types

package common

// TimeRange's fields are nil-able.
type TimeRange struct {
	start *string
	end *string
}

func (tr TimeRange) Start() *string {
	return tr.start
}

func (tr TimeRange) End() *string {
	return tr.end
}
