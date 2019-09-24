package commandline

import (
	"fmt"
	"github.com/jfixby/pin"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ExternalProcess is a helper class wrapping command line execution
type ExternalProcess struct {
	// CommandName stores console command name
	CommandName string

	// Arguments stores console command arguments
	Arguments []string

	// WaitForExit set to true tells the Launch()
	// to hold until the process finish it's job
	WaitForExit bool

	// isRunning indicates if given process is already running or not
	// to avoid double launch or double stop
	isRunning bool

	runningCommand *exec.Cmd
}

// Dispose implements integration.LeakyAsset behaviour
func (process *ExternalProcess) Dispose() {
	killProcess(process, os.Stdout)
}

// FullConsoleCommand returns full console command string
func (process *ExternalProcess) FullConsoleCommand() string {
	cmd := process.runningCommand
	args := strings.Join(cmd.Args[1:], " ")
	return cmd.Path + " " + args
}

// ClearArguments clears the Arguments list
func (process *ExternalProcess) ClearArguments() {
	process.Arguments = []string{}
}

// Launch launches external process
// set the debugOutput true to redirect external process output
// to your os.Stdout and os.Stderr
func (process *ExternalProcess) Launch(debugOutput bool) {
	if process.isRunning {
		pin.ReportTestSetupMalfunction(
			fmt.Errorf("process is already running: %v",
				process.runningCommand,
			),
		)
	}
	process.isRunning = true

	process.runningCommand = exec.Command(
		process.CommandName, process.Arguments...)
	cmd := process.runningCommand
	fmt.Println("run command # " + cmd.Path)
	fmt.Println(strings.Join(cmd.Args[0:], "\n    "))
	fmt.Println()
	if debugOutput {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	err := cmd.Start()
	pin.CheckTestSetupMalfunction(err)

	pin.RegisterDisposableAsset(process)

	if process.WaitForExit {
		process.waitForExit()
		return
	}
}

// Stop interrupts the running process, and waits until it exits properly.
// It is important that the process be stopped via Stop(), otherwise,
// it will persist unless explicitly killed.
func (process *ExternalProcess) Stop() error {
	if !process.isRunning {
		pin.ReportTestSetupMalfunction(
			fmt.Errorf("process is not running: %v",
				process.runningCommand,
			),
		)
	}
	process.isRunning = false

	defer pin.DeRegisterDisposableAsset(process)

	return killProcess(process, os.Stdout)
}

// waitForExit waits for ext process to exit
func (process *ExternalProcess) waitForExit() {
	err := process.runningCommand.Wait()
	pin.CheckTestSetupMalfunction(err)
	process.isRunning = false
	pin.DeRegisterDisposableAsset(process)
}

// IsRunning indicates when ext process is working
func (process *ExternalProcess) IsRunning() bool {
	return process.isRunning
}

func (process *ExternalProcess) Wait() error {
	cmd := process.runningCommand
	return cmd.Wait()
}

// On windows, interrupt is not supported, so a kill signal is used instead.
func killProcess(process *ExternalProcess, logStream *os.File) error {
	cmd := process.runningCommand
	defer cmd.Wait()

	fmt.Fprintln(
		logStream,
		fmt.Sprintf(
			"Killing process: %v",
			process.FullConsoleCommand(),
		))

	osProcess := cmd.Process

	if runtime.GOOS == "windows" {
		err := osProcess.Signal(os.Kill)
		return err
	}

	err := osProcess.Signal(os.Interrupt)
	return err
}
