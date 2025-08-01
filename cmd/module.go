package cmd

import (
	"github.com/GisangLee/tfcli/internal/module"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "📦 모듈 자동 참조",
	Long: `📦 모듈 자동 참조

지정한 모듈을 Terraform 환경 디렉토리(main.tf)에 자동으로 추가합니다.

예시:
  tfcli module add vpc --source=./modules/vpc
  → environment/{env}/main.tf에 module "vpc" 블록 삽입
`,
	Run: func(cmd *cobra.Command, args []string) {
		module.HandleModuleAdd()
	},
}

func init() {
	rootCmd.AddCommand(moduleCmd)

	moduleCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgHiGreen).SprintFunc()
		bold := color.New(color.Bold).SprintFunc()

		println(green("╭────────────────────────────────────────────────────────╮"))
		println(green("│ 📦 tfcli module - 모듈 자동 참조 도우미                          │"))
		println(green("╰────────────────────────────────────────────────────────╯"))

		println()
		println(bold("🛠 사용법:"))
		println("  tfcli module add [모듈명] --source=경로")

		println()
		println(bold("📚 예시:"))
		println("  tfcli module add vpc --source=./modules/vpc")

		println()
		println(bold("🔧 옵션:"))
		println("  -h, --help   도움말 출력")
	})
}
