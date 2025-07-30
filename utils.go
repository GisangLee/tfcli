package main

import (
	"github.com/manifoldco/promptui"
)

func parseTerraformArgs(task string) []string {
	switch task {
	case "init":
		return []string{"init"}
	case "fmt -recursive":
		return []string{"fmt", "-recursive"}
	case "plan":
		return []string{"plan", "-var-file=terraform.tfvars"}
	case "apply":
		return []string{"apply", "-auto-approve"}
	case "destroy":
		return []string{"destroy", "-auto-approve"}
	default:
		return []string{}
	}
}

func selectRootMenu() string {
	menu := promptui.Select{
		Label: "✨ 실행할 작업을 선택하세요",
		Items: []string{"create-project", "create-template", "tf-job"},
	}
	_, selection, err := menu.Run()
	if err != nil {
		return ""
	}
	return selection
}
