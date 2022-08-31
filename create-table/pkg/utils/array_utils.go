/*
Package utils contain util functions
*/
package utils

import "create-table/configs"

// Contain check contain TargetCreateTable with kubun, hassin and data type equal input
func Contain(target []configs.TargetCreateTable, kubun, hassin, dataType string) bool {
	for index := range target {
		if target[index].QKbn == kubun && target[index].Sndc == hassin && target[index].DataType == dataType {
			return true
		}
	}

	return false
}
