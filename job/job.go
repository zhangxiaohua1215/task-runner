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
	Input     string
	// 输出参数
	// StartedAt   time.Time
	// CompletedAt time.Time
	// StdOut     []byte
	// ExitedCode int
}

type Executor interface {
	Execute(scriptPath string, args []string, stdin []byte) (stdout []byte, exitCode int)
}

func NewExecutor(ext string) Executor {
	switch ext {
	case ".py":
		return nil
	case ".sh":
		return nil
	case ".exe":
		return &exeExecutor{}
	case ".ps1":
		return nil
	}
	return nil
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
		executor := NewExecutor(t.Ext)
		// 更新任务状态为正在执行
		service.AppServiceGroup.Start(t.ID)

		stdout, exeCode := executor.Execute(t.FilePath, t.Arguments, []byte(t.Input))

		status := service.TaskStatusCompleted
		if exeCode != 0 {
			status = service.TaskStatusFailed
		}

		// 更新任务状态为已完成
		service.AppServiceGroup.Complete(t.ID, status, stdout, exeCode)
	}
}

func RunWorker(workerCnt int) {
	for i := 0; i < workerCnt; i++ {
		go worker()
	}
}
