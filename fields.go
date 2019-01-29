package greasel

// Concrete field types

type IntField struct {
	Field
}

func NewIntField(parent *Table, name string) *IntField {
	field := NewField(parent, name)
	intField := &IntField{*field}
	parent.intFields[name] = intField
	return intField
}

func (f *IntField) Eq(val int) Filter {
	return ValueComparisonFilter{
		Lhs: &f.Field,
		Op: "=",
		Rhs: val,
	}
}

func (f *IntField) IsEq(other *IntField) Filter {
	return FieldComparisonFilter{
		Lhs: &f.Field,
		Op: "=",
		Rhs: &other.Field,
	}
}

// ...