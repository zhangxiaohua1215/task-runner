package job

import (
	"bytes"
	"os/exec"
	"task-runner/service"
)

var TaskQueue = make(chan Task, 1000)

type Task struct {
	ID        int64
	ScriptID  int64
	Arguments []string
	FilePath  string
	Ext       string
	StdIn     []byte
	// 输出参数
	// StartedAt   time.Time
	// CompletedAt time.Time
	// StdOut     []byte
	// ExitedCode int
}

type ScriptExecutor interface {
	Execute(scriptPath string, args []string, stdin []byte) (stdout []byte, exitCode int)
}

type exeExecutor struct{}

func (e *exeExecutor) Execute(scriptPath string, args []string, stdin []byte) (stdout []byte, exitCode int) {
	// 执行脚本
	cmd := exec.Command(scriptPath, args...)
	cmd.Stdin = bytes.NewBuffer(stdin)

	stdout, _ = cmd.Output()

	// 执行结果
	exitCode = cmd.ProcessState.ExitCode()

	return stdout, exitCode

}

func worker() {
	for t := range TaskQueue {
		// 更新任务状态为正在执行
		service.TaskServiceInstance.Start(t.ID)

		var executor ScriptExecutor
		switch t.Ext {
		case ".py":
			// 执行 Python 脚本
		case ".sh":
			// 执行 Shell 脚本
		case ".exe":
			executor = &exeExecutor{}
		case ".ps1":
			// 执行 PowerShell 脚本
		}
		stdout, exeCode := executor.Execute(t.FilePath, t.Arguments, t.StdIn)

		status := service.TaskStatusCompleted
		if exeCode != 0 {
			status = service.TaskStatusFailed
		}

		// 更新任务状态为已完成
		service.TaskServiceInstance.Complete(t.ID, status, stdout, exeCode)
	}
}

func RunWorker(workerCnt int) {
	for i := 0; i < workerCnt; i++ {
		go worker()
	}
}
