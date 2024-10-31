package job

import (
	"io"
	"log"
	"os"
	"os/exec"
	"task-runner/service"
	"task-runner/utils"
)

var TaskQueue = make(chan Task, 1000)

type Task struct {
	ID            int64
	ScriptID      int64
	Arguments     []string
	ScriptPath    string
	InputFileName string
	Ext           string
	Input         io.Reader
	// 输出参数
	// StartedAt   time.Time
	// CompletedAt time.Time
	// StdOut     []byte
	// ExitedCode int
}

type Executor interface {
	Execute(scriptPath string, args []string, stdin io.Reader, stdout io.Writer) (exitCode int)
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

func (e *exeExecutor) Execute(scriptPath string, args []string, stdin io.Reader, stdout io.Writer) (exitCode int) {
	// 执行脚本
	cmd := exec.Command(scriptPath, args...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout

	// 执行脚本
	_ = cmd.Run()
	// 执行结果
	exitCode = cmd.ProcessState.ExitCode()

	return exitCode

}

func worker() {
	for t := range TaskQueue {
		executor := NewExecutor(t.Ext)
		// 更新任务状态为正在执行
		service.AppServiceGroup.Start(t.ID)

		// 创建输出文件
		dst := utils.GenResultFilePath(t.ID)
		f, err := os.Create(dst)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		// 执行脚本
		exeCode := executor.Execute(t.ScriptPath, t.Arguments, t.Input, f)

		status := service.TaskStatusCompleted
		if exeCode != 0 {
			status = service.TaskStatusFailed
		}

		// 更新任务状态为已完成
		service.AppServiceGroup.Complete(t.ID, status, exeCode)
	}
}

func RunWorker(workerCnt int) {
	for i := 0; i < workerCnt; i++ {
		go worker()
	}
}
