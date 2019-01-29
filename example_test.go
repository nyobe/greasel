package greasel

import "fmt"

// Schemas

type userTable struct {
	*Table
	Id *IntField
}

func NewUserAlias() *userTable {
	table := NewTable("users")
	return &userTable{
		Table: table,
		Id: NewIntField(table, "id"),
	}
}

// Generate relation condition from bind
func (u *userTable) PhotoAssoc(p *photoTable) (*Table,  Filter) {
	return p.Table, p.UserId.IsEq(u.Id)
}

type photoTable struct {
	*Table
	Id *IntField
	UserId *IntField
}

func NewPhotoAlias() *photoTable {
	table := NewTable("posts")
	return &photoTable{
		Table: table,
		Id: NewIntField(table, "id"),
		UserId: NewIntField(table, "user_id"),
	}
}

// Examples

func Example() {
	u, p := NewUserAlias(), NewPhotoAlias()

	q := From(u.Table). // fixme -> interface
		Where(u.Id.Eq(1)).
		InnerJoin(u.PhotoAssoc(p))

	fmt.Println(q.SQL())

	// Output:
	// FROM users u0 INNER JOIN posts AS p1 ON (p1.user_id = u0.id )  WHERE (u0.id = ? )
}