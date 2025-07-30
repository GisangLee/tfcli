package cmd

import (
	"github.com/GisangLee/tfcli/internal/banner"
	"github.com/GisangLee/tfcli/internal/project"
	"github.com/GisangLee/tfcli/internal/template"
	"github.com/GisangLee/tfcli/internal/tfjob"

	"github.com/manifoldco/promptui"
)

func Execute() {
	banner.Show()

	rootPrompt := promptui.Select{
		Label: "✨ 실행할 작업을 선택하세요",
		Items: []string{"create-project", "create-template", "tf-job"},
	}
	_, selectedTask, err := rootPrompt.Run()
	if err != nil {
		println("❌ 작업 선택 실패:", err)
		return
	}

	switch selectedTask {
	case "create-project":
		project.HandleCreateProject()
	case "create-template":
		template.HandleCreateTemplate()
	case "tf-job":
		tfjob.HandleTfJob()
	default:
		println("❌ 알 수 없는 작업입니다.")
	}
}
