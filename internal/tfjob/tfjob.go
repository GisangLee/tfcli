package tfjob

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/GisangLee/tfcli/internal/utils"
	"github.com/manifoldco/promptui"
)

func HandleTfJob() {
	cspPrompt := promptui.Select{
		Label: "☁️ CSP를 선택하세요",
		Items: []string{"aws", "ncp", "gcp"},
	}
	_, selectedCSP, err := cspPrompt.Run()
	if err != nil {
		fmt.Println("❌ CSP 선택 실패:", err)
		return
	}

	envPrompt := promptui.Select{
		Label: "🌎 환경을 선택하세요",
		Items: []string{"dev", "stage", "prod"},
	}
	_, selectedEnv, err := envPrompt.Run()
	if err != nil {
		fmt.Println("❌ 환경 선택 실패:", err)
		return
	}

	taskPrompt := promptui.Select{
		Label: "⚙️ 실행할 Terraform 작업을 선택하세요",
		Items: []string{"init", "fmt -recursive", "plan", "apply", "destroy"},
	}
	_, selectedTask, err := taskPrompt.Run()
	if err != nil {
		fmt.Println("❌ 작업 선택 실패:", err)
		return
	}

	projectRoot := filepath.Join(selectedCSP)

	if selectedTask == "fmt -recursive" {
		if err := os.Chdir(projectRoot); err != nil {
			fmt.Println("❌ 프로젝트 루트 이동 실패:", err)
			return
		}
		fmt.Println("📁 위치 이동 (fmt 전체 대상):", projectRoot)
	} else {
		envPath := filepath.Join(projectRoot, "environment", selectedEnv)
		if err := os.Chdir(envPath); err != nil {
			fmt.Println("❌ 환경 디렉토리 이동 실패:", err)
			return
		}
		fmt.Println("📁 위치 이동:", envPath)
	}

	args := utils.ParseTerraformArgs(selectedTask)
	cmd := exec.Command("terraform", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Printf("\n🚀 실행 중: terraform %s\n\n", selectedTask)
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ 실행 중 오류:", err)
		return
	}
	fmt.Println("✅ 완료:", selectedTask)
}
