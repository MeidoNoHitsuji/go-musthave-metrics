package flags

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFlags(t *testing.T) {
	err := ParseFlags()
	assert.NoError(t, err)
	assert.Equal(t, ReportInterval, 10)
	assert.Equal(t, RollInterval, 2)
	assert.Equal(t, Addr, "localhost:8080")
}
