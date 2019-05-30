
package commandline

import (
	"github.com/jfixby/pin"
	"testing"

)

// TestGoExample launches example `go help` process
func TestGoExample(t *testing.T) {

	{
		proc := &ExternalProcess{
			CommandName: "go",
			WaitForExit: true,
		}
		proc.Arguments = append(proc.Arguments, "help")

		debugOutput := true
		proc.Launch(debugOutput)
	}

	// Verify proper disposal
	pin.VerifyNoAssetsLeaked()
}
