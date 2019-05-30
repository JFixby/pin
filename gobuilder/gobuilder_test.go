
package gobuilder

import (
	"github.com/jfixby/pin"
	"os"
	"testing"

)

// TestGoBuider builds current project executable
func TestGoBuider(t *testing.T) {
	defer pin.VerifyNoAssetsLeaked()
	runExample()
}

func runExample() {
	testWorkingDir := pin.NewTempDir(os.TempDir(), "test-go-builder")

	testWorkingDir.MakeDir()
	defer testWorkingDir.Dispose()

	builder := &GoBuider{
		GoProjectPath:    DetermineProjectPackagePath("pin"),
		OutputFolderPath: testWorkingDir.Path(),
		BuildFileName:    "pin",
	}

	builder.Build()
	defer builder.Dispose()
}
