package mypkg

import "fmt"

func Query1() string {
	return fmt.Sprintf(`SELECT '%s' AS type_helper, id, name FROM users`, Query1Helper)
}

func Query2() string {
	return fmt.Sprintf(`SELECT '%s' AS type_helper, id, city, state FROM addresses`, Query2Helper)
}
