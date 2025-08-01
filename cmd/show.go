package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var format string

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "🎨 terraform show 결과 출력",
	Run: func(cmd *cobra.Command, args []string) {
		runShow()
	},
}

func init() {
	showCmd.Flags().StringVar(&format, "format", "markdown", "출력 형식: markdown, json, html 중 하나")
	showCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgHiGreen).SprintFunc()
		bold := color.New(color.Bold).SprintFunc()

		fmt.Println(green("╭────────────────────────────────────────────╮"))
		fmt.Println(green("│ 🎨 show - Terraform 상태 보기              │"))
		fmt.Println(green("╰────────────────────────────────────────────╯"))
		fmt.Println()
		fmt.Println(bold("🧭 설명:"))
		fmt.Println("  terraform show 명령을 실행하고 원하는 포맷으로 출력합니다.")
		fmt.Println()
		fmt.Println(bold("🛠 사용법:"))
		fmt.Println("  tfcli show --format=json")
		fmt.Println()
		fmt.Println(bold("🔧 옵션:"))
		fmt.Println("  --format=markdown|json|html")
	})
	rootCmd.AddCommand(showCmd)
}

func runShow() {
	// CSP 선택
	prompt := promptui.Select{
		Label: "☁️ CSP를 선택하세요",
		Items: []string{"aws", "ncp", "gcp"},
	}
	_, csp, err := prompt.Run()
	if err != nil {
		fmt.Println("❌ CSP 선택 실패:", err)
		return
	}

	promptEnv := promptui.Select{
		Label: "📦 환경을 선택하세요",
		Items: []string{"dev", "stage", "prod"},
	}
	_, env, err := promptEnv.Run()
	if err != nil {
		fmt.Println("❌ 환경 선택 실패:", err)
		return
	}

	tfstatePath := fmt.Sprintf("%s/environment/%s/.terraform/terraform.tfstate", csp, env)
	data, err := os.ReadFile(tfstatePath)
	if err != nil {
		fmt.Println("❌ tfstate 파일을 읽을 수 없습니다:", err)
		return
	}

	switch format {
	case "json":
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, data, "", "  ")
		if err != nil {
			fmt.Println("❌ JSON 포맷 처리 실패:", err)
			return
		}
		fmt.Println(prettyJSON.String())

	case "markdown":
		fmt.Println("📄 *Terraform State (Markdown)*")
		fmt.Println("```json")
		fmt.Println(string(data))
		fmt.Println("```")

	case "html":
		fmt.Println("<h2>Terraform State (HTML)</h2>")
		fmt.Println("<pre>")
		fmt.Println(strings.ReplaceAll(string(data), "<", "&lt;")) // escape
		fmt.Println("</pre>")

	default:
		fmt.Println("❌ 지원되지 않는 포맷입니다. --format=json|markdown|html 중 선택하세요.")
	}
}
