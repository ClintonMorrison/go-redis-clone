package primitives

type Value struct {
	Value string
	Type string
}

func (v Value) Serialize() string {
	return v.Value
}
