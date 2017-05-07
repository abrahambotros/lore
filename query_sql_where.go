package lore

import "github.com/Masterminds/squirrel"

/*
Where returns a new part as a Sqlizer interface, allowing this to be called in typical WHERE
clauses via Squirrel.
*/
func Where(pred interface{}, args ...interface{}) squirrel.Sqlizer {
	return &part{
		pred: pred,
		args: args,
	}
}
