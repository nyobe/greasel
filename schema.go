package greasel

// Table represents a db table, either a physical table or a table derived from a query
type Table struct {
	// Table name
	source string

	// Reflection
	fields map[string]*Field
	intFields map[string]*IntField
}

func NewTable(name string) *Table {
	if name == "" {
		panic("table name required")
	}
	return &Table{
		source: name,
		fields: make(map[string]*Field),
		intFields: make(map[string]*IntField),
	}
}

// Fields returns all fields on this table
func (t *Table) Fields() []*Field {
	lst := make([]*Field, 0, len(t.fields))
	for _, f := range t.fields {
		lst = append(lst, f)
	}
	return lst
}

// Field returns field by name
func (t *Table) Field(name string) *Field {
	return t.fields[name]
}

// Field represents a table field of any type
type Field struct {
	parent *Table
	name   string
}

// Create field and register it with parent table
func NewField(parent *Table, name string) *Field {
	if _, exists := parent.fields[name]; exists {
		panic("field with this name already exists in table")
	}
	field := &Field{parent: parent, name: name}
	parent.fields[name] = field
	return field
}
