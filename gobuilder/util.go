
package gobuilder

import (
	"fmt"
	"github.com/jfixby/pin"
	"runtime"
	"strings"

)

// DetermineProjectPackagePath is used to determine target project path
// for the following execution of a Go builder.
// Starts from a current working directory, climbs up to a parent folder
// with the target name and returns its path as a result.
func DetermineProjectPackagePath(projectName string) string {
	// Determine import path of this package.
	_, launchDir, _, ok := runtime.Caller(1)
	if !ok {
		pin.CheckTestSetupMalfunction(
			fmt.Errorf("cannot get project <%v> path, launch dir is: %v ",
				projectName,
				launchDir,
			),
		)
	}
	sep := "/"
	steps := strings.Split(launchDir, sep)
	for i, s := range steps {
		if s == projectName {
			pkgPath := strings.Join(steps[:i+1], "/")
			return pkgPath
		}
	}
	pin.CheckTestSetupMalfunction(
		fmt.Errorf("cannot get project <%v> path, launch dir is: %v ",
			projectName,
			launchDir,
		),
	)
	return ""
}
