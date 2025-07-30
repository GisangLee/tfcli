package main

import (
	"fmt"

	"github.com/fatih/color"
)

func showBanner() {
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	fmt.Println()
	fmt.Println(cyan("╔══════════════════════════════════════════════════════╗"))
	fmt.Println(cyan("║"), bold("         🚀 Terraform Automation CLI v1.0           "), cyan("║"))
	fmt.Println(cyan("║"), yellow("     CSP 환경을 손쉽게 선택하고 작업하세요!     "), cyan("║"))
	fmt.Println(cyan("║"), green("    Made with ☕️ & 💻 by DevOps JSON     "), cyan("║"))
	fmt.Println(cyan("╚══════════════════════════════════════════════════════╝"))
	fmt.Println(magenta("\n✨ 시작해볼까요?\n"))
}
