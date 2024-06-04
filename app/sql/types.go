package sql

type SelectStatement struct {
	Fields []string
	Table  string
	Where  *Expression
}

type Expression struct {
	Attributes []Attribute
	Operator   string
}

type Attribute struct {
	Type  string // "COLUMN", "VALUE", "EXPR"
	Value string
	Expr  *Expression
}
