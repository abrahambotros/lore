package lore

/*
Where returns a new part as a Sqlizer interface, allowing this to be called in typical WHERE
clauses via Squirrel.
*/
func Where(pred interface{}, args ...interface{}) *sqlPart {
	return newSqlPart(pred, args...)
}
