package cmd

import (
	"fmt"

	"github.com/GisangLee/tfcli/internal/tfjob"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var tfCmd = &cobra.Command{
	Use:   "tf",
	Short: "🚀 Terraform 작업 실행",
	Long: `🚀 Terraform 작업 실행

Terraform의 주요 작업(init, fmt, plan, apply, destroy)을 선택적으로 실행할 수 있습니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		tfjob.HandleTfJob()
	},
}

func init() {
	tfCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgHiGreen).SprintFunc()
		bold := color.New(color.Bold).SprintFunc()

		fmt.Println(green("╭──────────────────────────────────────────────╮"))
		fmt.Println(green("│ 🚀 tf - Terraform 작업 실행                    │"))
		fmt.Println(green("╰──────────────────────────────────────────────╯"))
		fmt.Println()
		fmt.Println(bold("🧭 설명:"))
		fmt.Println("  Terraform의 주요 명령어를 선택적으로 실행합니다.")
		fmt.Println()
		fmt.Println(bold("🛠 사용법:"))
		fmt.Println("  tfcli tf")
		fmt.Println()
		fmt.Println(bold("⚙️ 실행 가능한 작업:"))
		fmt.Println("  init, fmt -recursive, plan, apply, destroy")
		fmt.Println()
		fmt.Println(bold("🔧 옵션:"))
		fmt.Println("  -h, --help   도움말 출력")
	})
	rootCmd.AddCommand(tfCmd)
}
