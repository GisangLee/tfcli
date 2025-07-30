package project

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

func HandleCreateProject() {
	cspPrompt := promptui.Select{
		Label: "☁️ 초기화할 CSP를 선택하세요",
		Items: []string{"aws", "ncp", "gcp"},
	}
	_, selectedCSP, err := cspPrompt.Run()
	if err != nil {
		fmt.Println("❌ CSP 선택 실패:", err)
		return
	}
	createProjectStructure(selectedCSP)
}

func createProjectStructure(csp string) {
	root := filepath.Join(".", csp)
	modules := filepath.Join(root, "modules")
	envRoot := filepath.Join(root, "environment")
	envs := []string{"dev", "stage", "prod"}

	dirs := append([]string{modules}, func() []string {
		result := []string{}
		for _, e := range envs {
			result = append(result, filepath.Join(envRoot, e))
		}
		return result
	}()...)

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("❌ 디렉토리 생성 실패: %s\n", dir)
		} else {
			fmt.Printf("✅ 디렉토리 생성됨: %s\n", dir)
		}
	}
	fmt.Println("📁 프로젝트 기본 구조가 생성되었습니다.")
}
