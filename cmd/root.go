package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tfcli",
	Short: "🌱 tfcli - Terraform 프로젝트를 손쉽게 관리하는 CLI 도구",
	Long:  `🌱 tfcli는 Terraform 기반 인프라 관리를 더욱 편리하게 해주는 CLI 도구입니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintBanner()
		customHelp(cmd, args)
	},
}

func Execute() {
	cobra.OnInitialize()
	rootCmd.SetHelpFunc(customHelp)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌ 오류:", err)
		os.Exit(1)
	}
}

func customHelp(cmd *cobra.Command, args []string) {
	green := color.New(color.FgHiGreen).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	fmt.Println(green("╭────────────────────────────────────────────────────────╮"))
	fmt.Println(green("│ 🌱 tfcli - Terraform 프로젝트를 손쉽게 관리하는 CLI 도구  │"))
	fmt.Println(green("╰────────────────────────────────────────────────────────╯"))

	fmt.Println()
	fmt.Println(bold("🛠 사용법:"))
	fmt.Println("  tfcli [명령어] [옵션]")

	fmt.Println()
	fmt.Println(bold("📚 사용 가능한 명령어:"))
	fmt.Println("  project     📁 TFCLI 프로젝트 구조 생성")
	fmt.Println("  template    🧩 모듈 템플릿 생성")
	fmt.Println("  module      📦 모듈 자동 참조")
	fmt.Println("  tf          🚀 Terraform 작업 실행 (init/fmt --recursive/plan/apply/destroy)")

	fmt.Println()
	fmt.Println(bold("🔧 옵션:"))
	fmt.Println("  -h, --help   도움말 출력")
}
