
package pin

import (
	"path/filepath"
	"testing"
)

// TestLeakyAssets creates example leaky assets,
// checks they were properly disposed before exit.
func TestLeakyAssets(t *testing.T) {
	a := NewTempDir("", "a")
	a.MakeDir()

	b := NewTempDir("", filepath.Join("a", "b"))
	b.MakeDir()

	c := NewTempDir("", filepath.Join("a", "b", "c"))
	c.MakeDir()

	c.Dispose()
	b.Dispose()
	a.Dispose()

	VerifyNoAssetsLeaked()

	b.MakeDir()
	a.MakeDir()
	c.MakeDir()

	b.Dispose()
	c.Dispose()
	a.Dispose()

	VerifyNoAssetsLeaked()
}
