package tfjob

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

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
	// fmt, init은 로그 저장 없이 바로 실행
	if selectedTask == "init" || selectedTask == "fmt -recursive" {
		runSimpleTerraform(selectedTask)
	} else {
		runTerraformWithLog(selectedTask)
	}
}

func runSimpleTerraform(task string) {
	args := utils.ParseTerraformArgs(task)
	cmd := exec.Command("terraform", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Printf("\n🚀 실행 중: terraform %s\n\n", task)
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ 실행 실패:", err)
	} else {
		fmt.Println("✅ 완료:", task)
	}
}

func runTerraformWithLog(task string) error {
	timestamp := time.Now().Format("20060102_150405")

	// 현재 작업 디렉토리 기준 logs 디렉토리
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("작업 디렉토리 조회 실패: %w", err)
	}
	logDir := filepath.Join(cwd, "logs")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("로그 디렉토리 생성 실패: %w", err)
	}

	logFile := filepath.Join(logDir, fmt.Sprintf("%s_%s.log", task, timestamp))
	f, err := os.Create(logFile)
	if err != nil {
		return fmt.Errorf("로그 파일 생성 실패: %w", err)
	}
	defer f.Close()

	// 로그 헤더 작성
	header := fmt.Sprintf(`╭───────────────────────────── TFCLI 작업 로그 ─────────────────────────────╮
🕒 시간: %s
🔧 작업: terraform %s
📁 경로: %s
╰──────────────────────────────────────────────────────────────────────────╯

`, time.Now().Format(time.RFC3339), task, cwd)
	f.WriteString(header)
	fmt.Print(header)

	// 실행
	args := utils.ParseTerraformArgs(task)
	cmd := exec.Command("terraform", args...)

	// 표준 출력 로그 → 터미널 + 파일에 동시에 저장
	cmd.Stdout = io.MultiWriter(os.Stdout, f)
	cmd.Stderr = io.MultiWriter(os.Stderr, f)
	cmd.Stdin = os.Stdin

	fmt.Printf("🚀 실행 중: terraform %s → 로그 저장: %s\n\n", task, logFile)
	err = cmd.Run()

	// 작업 결과 출력
	if err != nil {
		msg := fmt.Sprintf("\n❌ 실행 실패: %v\n", err)
		f.WriteString(msg)
		fmt.Print(msg)
		return err
	} else {
		msg := "\n✅ 작업 성공\n"
		f.WriteString(msg)
		fmt.Print(msg)
		return nil
	}
}
