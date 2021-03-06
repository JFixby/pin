package commandline

// ExecutablePathProvider wraps class responsible for executable
// path resolution
type ExecutablePathProvider interface {
	// Executable returns full path to an executable target file
	Executable() string
}

type ExplicitExecutablePathString struct {
	PathString string
}

func (pth *ExplicitExecutablePathString) Executable() string {
	return pth.PathString
}
