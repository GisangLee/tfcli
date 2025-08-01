package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func PromptCSP() (string, error) {
	cspPrompt := promptui.Select{
		Label: "☁️ CSP를 선택하세요",
		Items: []string{"aws", "ncp", "gcp"},
	}
	_, selectedCSP, err := cspPrompt.Run()
	if err != nil {
		return "", fmt.Errorf("❌ CSP 선택 실패: %w", err)
	}
	return selectedCSP, nil
}

func PromptEnv() (string, error) {
	envPrompt := promptui.Select{
		Label: "🌎 환경을 선택하세요",
		Items: []string{"dev", "stage", "prod"},
	}
	_, selectedEnv, err := envPrompt.Run()
	if err != nil {
		return "", fmt.Errorf("❌ 환경 선택 실패:", err)
	}

	return selectedEnv, nil
}
