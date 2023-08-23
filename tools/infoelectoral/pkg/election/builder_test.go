package election

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildFilenameGenerator(t *testing.T) {
	var filenameGenerator = buildFilenameGenerator(2, 3, 2023)

	assert.Equal(t, "03022303.DAT", filenameGenerator(true, "03"))
	assert.Equal(t, "", filenameGenerator(false, "03"))
}

func Test_buildCustomPrefixFilenameGenerator(t *testing.T) {
	var filenameGenerator = buildCustomPrefixFilenameGenerator(3, 2023)

	assert.Equal(t, "05012303.DAT", filenameGenerator(true, "0501"))
	assert.Equal(t, "", filenameGenerator(false, "0501"))
}
