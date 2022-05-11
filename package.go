//go:generate go run generate.go

package postgresqlreservedwords

import "strings"

func IsReserved(word string) bool {
	_, ok := Words[strings.ToLower(strings.TrimSpace(word))]
	return ok
}
