
package gobuilder

import (
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/commandline"
	"go/build"
	"path/filepath"
	"runtime"
	"sync"

)

// GoBuider helps to build a target Go project
type GoBuider struct {
	// GoProjectPath is a path to the target Go-project
	GoProjectPath string

	// BuildFileName stores executable file name
	BuildFileName string

	// OutputFolderPath points to output executable parent folder
	OutputFolderPath string

	compileMtx sync.Mutex
}

// Dispose required to implement integration.LeakyAsset
func (builder *GoBuider) Dispose() {
	deleteOutputExecutable(builder)
	pin.DeRegisterDisposableAsset(builder)
}

// Executable returns full path to an executable target file
func (builder *GoBuider) Executable() string {
	outputPath := filepath.Join(
		builder.OutputFolderPath, builder.BuildFileName)
	if runtime.GOOS == "windows" {
		outputPath += ".exe"
	}
	return outputPath
}

// Build compiles target project and writes output to the target output folder
func (builder *GoBuider) Build() {
	builder.compileMtx.Lock()
	defer builder.compileMtx.Unlock()

	goProjectPath := builder.GoProjectPath
	outputFolderPath := builder.OutputFolderPath
	pin.MakeDirs(outputFolderPath)

	target := builder.Executable()
	if pin.FileExists(target) {
		deleteOutputExecutable(builder)
		pin.DeRegisterDisposableAsset(builder)
	}

	// check project path
	pkg, err := build.ImportDir(goProjectPath, build.FindOnly)
	pin.CheckTestSetupMalfunction(err)
	goProjectPath = pkg.ImportPath

	runBuildCommand(builder, goProjectPath)
	pin.RegisterDisposableAsset(builder)
}

// runBuildCommand calls `go build`
func runBuildCommand(builder *GoBuider, goProjectPath string) {
	// Build and output an executable in a static temp path.
	proc := &commandline.ExternalProcess{
		CommandName: "go",
		WaitForExit: true,
	}
	proc.Arguments = append(proc.Arguments, "build")
	proc.Arguments = append(proc.Arguments, "-v")
	//proc.Arguments = append(proc.Arguments, "-x")
	proc.Arguments = append(proc.Arguments, "-o")
	proc.Arguments = append(proc.Arguments, builder.Executable())
	proc.Arguments = append(proc.Arguments, goProjectPath)

	proc.Launch(true)
}

func deleteOutputExecutable(builder *GoBuider) {
	pin.DeleteFile(builder.Executable())
}
