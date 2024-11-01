package job

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

	cmd := exec.Command(scriptPath, args...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout

	// 执行脚本
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	// 执行结果
	return cmd.ProcessState.ExitCode()
}

// 根据id创建输出文件id, host => f, 

func worker() {
	for t := range TaskQueue {
		executor := NewExecutor(t.Ext)
		// 更新任务状态为正在执行
		service.AppServiceGroup.Start(t.ID)

		// 创建输出文件
		dst := utils.GenResultFilePath(t.ID)
		err := os.MkdirAll(filepath.Dir(dst), os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		f, err := os.Create(dst)
		if err != nil {
			log.Fatalln(err)
		}
		fileUrl := fmt.Sprintf("http://localhost:8080/%s", dst)
		// 执行脚本
		exeCode := executor.Execute(t.ScriptPath, t.Arguments, t.Input, f)
		f.Close()

		status := service.TaskStatusCompleted
		if exeCode != 0 {
			status = service.TaskStatusFailed
		}

		// 更新任务状态为已完成
		service.AppServiceGroup.Complete(t.ID, status, exeCode, fileUrl)
	}
}

func RunWorker(workerCnt int) {
	for i := 0; i < workerCnt; i++ {
		go worker()
	}
}
