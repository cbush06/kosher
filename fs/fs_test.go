package fs

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestNewFs(t *testing.T) {
	t.Run("Invalid-Working-Directory", func(t *testing.T) {
		_, err := newFs(afero.NewMemMapFs(), "blahblahblah")
		assert.NotNil(t, err, "expected newFs() to fail, but did not")
	})
}
