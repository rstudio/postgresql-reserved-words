package postgresqlreservedwords_test

import (
	"testing"

	postgresqlreservedwords "github.com/rstudio/postgresql-reserved-words"
	"github.com/stretchr/testify/assert"
)

func TestWords(t *testing.T) {
	assert.GreaterOrEqual(t, len(postgresqlreservedwords.Words), 700)
}

func TestIsReserved(t *testing.T) {
	assert.True(t, postgresqlreservedwords.IsReserved("SELECT"))
	assert.False(t, postgresqlreservedwords.IsReserved("Shellac"))
}
