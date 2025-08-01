package cmd

import (
	"fmt"

	"github.com/GisangLee/tfcli/internal/template"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "🧩 모듈 템플릿 생성",
	Long: `🧩 Terraform 모듈 템플릿 생성

모듈 이름과 리소스 유형을 입력받아 표준 템플릿(main.tf, variables.tf 등)을 생성합니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		template.HandleCreateTemplate()
	},
}

func init() {
	templateCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgHiGreen).SprintFunc()
		bold := color.New(color.Bold).SprintFunc()

		fmt.Println(green("╭──────────────────────────────────────────────╮"))
		fmt.Println(green("│ 🧩 template - Terraform 모듈 템플릿 생성      │"))
		fmt.Println(green("╰──────────────────────────────────────────────╯"))
		fmt.Println()
		fmt.Println(bold("🧭 설명:"))
		fmt.Println("  원하는 이름과 리소스 종류에 맞춰 템플릿을 생성합니다.")
		fmt.Println()
		fmt.Println(bold("🛠 사용법:"))
		fmt.Println("  tfcli template")
		fmt.Println()
		fmt.Println(bold("📁 생성되는 파일:"))
		fmt.Println("  main.tf, variables.tf, outputs.tf")
		fmt.Println()
		fmt.Println(bold("🔧 옵션:"))
		fmt.Println("  -h, --help   도움말 출력")
	})
	rootCmd.AddCommand(templateCmd)
}
