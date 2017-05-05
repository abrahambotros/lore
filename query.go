package lore

/*
Query provides a generic, chainable query instance for callers to configure a specific query.
*/
type Query struct {
	modelInterface ModelInterface
	sqlBuilder     SqlBuilderInterface
}

/*
NewQuery instantiates a new Query instance based on the given ModelInterface.
*/
func NewQuery(modelInterface ModelInterface) *Query {
	return &Query{
		modelInterface: modelInterface,
		sqlBuilder:     nil,
	}
}
