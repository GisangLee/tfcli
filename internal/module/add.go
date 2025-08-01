package module

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GisangLee/tfcli/internal/utils"
	"github.com/manifoldco/promptui"
)

func HandleModuleAdd() {
	selectedCSP, _ := utils.PromptCSP()
	selectedEnv, _ := utils.PromptEnv()

	namePrompt := promptui.Prompt{Label: "📦 모듈 이름 입력 (예: vpc)"}
	moduleName, err := namePrompt.Run()
	if err != nil || moduleName == "" {
		fmt.Println("❌ 모듈 이름 입력 실패")
		return
	}
	sourcePath := filepath.Join("../../modules", moduleName)
	if err := insertModuleBlock(selectedCSP, selectedEnv, moduleName, sourcePath); err != nil {
		fmt.Println("❌ 모듈 삽입 실패:", err)
	}
}

// InsertModuleBlock 삽입 함수
func insertModuleBlock(csp, env, name, source string) error {
	mainTfPath := filepath.Join(csp, "environment", env, "main.tf")

	// 파일이 없으면 생성
	if _, err := os.Stat(mainTfPath); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(mainTfPath), os.ModePerm); err != nil {
			return fmt.Errorf("디렉토리 생성 실패: %w", err)
		}
		if _, err := os.Create(mainTfPath); err != nil {
			return fmt.Errorf("main.tf 생성 실패: %w", err)
		}
	}

	// 중복 확인
	exists, err := moduleExists(mainTfPath, name)
	if err != nil {
		return err
	}
	if exists {
		fmt.Printf("⚠️  module \"%s\" 블록이 이미 존재합니다 (건너뜀)\n", name)
		return nil
	}

	// 모듈 블록 삽입
	block := fmt.Sprintf(`

module "%s" {
  source = "%s"
}
`, name, source)

	f, err := os.OpenFile(mainTfPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("main.tf 열기 실패: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(block); err != nil {
		return fmt.Errorf("main.tf 쓰기 실패: %w", err)
	}

	fmt.Printf("✅ module \"%s\" 블록이 삽입되었습니다: %s\n", name, mainTfPath)
	return nil
}

func moduleExists(path, name string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, fmt.Errorf("파일 열기 실패: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, fmt.Sprintf("module \"%s\"", name)) {
			return true, nil
		}
	}
	return false, scanner.Err()
}
