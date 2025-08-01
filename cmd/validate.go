package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var validateAllCmd = &cobra.Command{
	Use:   "validate-all",
	Short: "🛠️ 모든 Terraform 구성에 대해 validate 실행",
	Run: func(cmd *cobra.Command, args []string) {
		cspPrompt := promptui.Select{
			Label: "☁️ CSP를 선택하세요",
			Items: []string{"aws", "ncp", "gcp"},
		}
		_, selectedCSP, err := cspPrompt.Run()
		if err != nil {
			fmt.Println("❌ CSP 선택 실패:", err)
			return
		}
		validateAll(selectedCSP)
	},
}

func validateAll(csp string) {
	rootPaths := []string{
		filepath.Join(csp, "modules"),
		filepath.Join(csp, "environment", "dev"),
		filepath.Join(csp, "environment", "stage"),
		filepath.Join(csp, "environment", "prod"),
	}

	success := 0
	fail := 0
	fmt.Printf("✔ %s\n", csp)

	for _, path := range rootPaths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		fmt.Printf("📍 검사 대상: %-40s", path)
		cmd := exec.Command("terraform", "validate")
		cmd.Dir = path
		if err := cmd.Run(); err != nil {
			fmt.Printf("❌ 실패\n")
			fail++
		} else {
			fmt.Printf("✅ 성공\n")
			success++
		}
	}

	fmt.Printf("\n✅ 유효성 검사 완료: %d 성공 / %d 실패\n", success, fail)
}

func init() {
	validateAllCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgHiGreen).SprintFunc()
		bold := color.New(color.Bold).SprintFunc()

		fmt.Println(green("╭────────────────────────────────────────────────────╮"))
		fmt.Println(green("│ 🛠️ validate-all - 모든 Terraform 디렉토리 검사     │"))
		fmt.Println(green("╰────────────────────────────────────────────────────╯"))

		fmt.Println()
		fmt.Println(bold("🧭 설명:"))
		fmt.Println("  선택한 CSP의 modules/와 environment/* 디렉토리를 순회하며 terraform validate를 수행합니다.")

		fmt.Println()
		fmt.Println(bold("🛠 사용법:"))
		fmt.Println("  tfcli validate-all")

		fmt.Println()
		fmt.Println(bold("🔧 옵션:"))
		fmt.Println("  -h, --help   도움말 출력")
	})
	rootCmd.AddCommand(validateAllCmd)
}
