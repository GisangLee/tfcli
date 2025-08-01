package template

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/GisangLee/tfcli/internal/utils"
	"github.com/manifoldco/promptui"
)

func HandleCreateTemplate() {

	selectedCSP, _ := utils.PromptCSP()
	prompt := promptui.Prompt{
		Label: "📦 생성할 리소스 이름 입력 (예: vpc, alb, eks 등)",
	}
	resourceName, err := prompt.Run()
	if err != nil || resourceName == "" {
		fmt.Println("❌ 리소스 이름 입력 실패")
		return
	}

	createTemplate(selectedCSP, resourceName)
}

func createTemplate(csp, resource string) {
	absTargetDir, err := filepath.Abs(filepath.Join(csp, "modules", resource))
	if err != nil {
		fmt.Println("❌ 절대경로 변환 실패:", err)
		return
	}
	targetDir := absTargetDir

	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		fmt.Println("❌ 디렉토리 생성 실패:", err)
		return
	}

	createFile := func(filename, content string) {
		fullPath := filepath.Join(targetDir, filename)
		if _, err := os.Stat(fullPath); err == nil {
			fmt.Printf("⚠️  이미 존재합니다: %s (건너뜀)\n", fullPath)
			return
		}
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			fmt.Printf("❌ %s 생성 실패: %v\n", filename, err)
		} else {
			fmt.Printf("✅ %s 생성 완료: %s\n", filename, fullPath)
		}
	}

	createFile("main.tf", "// main.tf 템플릿입니다.\n")
	createFile("variables.tf", "// variables.tf 템플릿입니다.\n")
	createFile("outputs.tf", "// outputs.tf 템플릿입니다.\n")

	if csp == "ncp" {
		createFile("providers.tf", `terraform {
  required_providers {
    ncloud = {
      source = "NaverCloudPlatform/ncloud"
    }
  }
  required_version = ">= 0.13"
}`)
	}

	fmt.Println("\n📦 템플릿 생성이 완료되었습니다.")
}
