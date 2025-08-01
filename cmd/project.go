package cmd

import (
	"fmt"

	"github.com/GisangLee/tfcli/internal/project"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "📁 Terraform 프로젝트 구조 생성",
	Long: `📁 Terraform 프로젝트 구조 생성

해당 명령어를 통해 AWS/NCP/GCP 환경의 Terraform 프로젝트 디렉토리를 쉽게 초기화할 수 있습니다.

예시:
  tfcli project
  → CSP 선택 → dev/stage/prod 환경 디렉토리 및 modules 디렉토리 생성
`,
	Run: func(cmd *cobra.Command, args []string) {
		project.HandleCreateProject()
	},
}

func init() {
	projectCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgHiGreen).SprintFunc()
		bold := color.New(color.Bold).SprintFunc()

		fmt.Println(green("╭──────────────────────────────────────────────╮"))
		fmt.Println(green("│ 📁 project - Terraform 프로젝트 구조 생성     │"))
		fmt.Println(green("╰──────────────────────────────────────────────╯"))

		fmt.Println()
		fmt.Println(bold("🧭 설명:"))
		fmt.Println("  AWS / NCP / GCP 환경의 Terraform 프로젝트 디렉토리를 빠르게 생성합니다.")

		fmt.Println()
		fmt.Println(bold("🛠 사용법:"))
		fmt.Println("  tfcli project")

		fmt.Println()
		fmt.Println(bold("📂 생성 구조:"))
		fmt.Println("  [csp]/")
		fmt.Println("    ├─ modules/")
		fmt.Println("    └─ environment/")
		fmt.Println("         ├─ dev/")
		fmt.Println("         ├─ stage/")
		fmt.Println("         └─ prod/")

		fmt.Println()
		fmt.Println(bold("🔧 옵션:"))
		fmt.Println("  -h, --help   도움말 출력")
	})
	rootCmd.AddCommand(projectCmd)
}
