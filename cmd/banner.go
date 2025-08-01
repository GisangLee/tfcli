package cmd

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintBanner() {
	c := color.New(color.FgHiGreen).Add(color.Bold)

	logo := `
	_______ ______  _____ _      _____ _____ 
   |__   __|  ____|/ ____| |    |_   _/ ____|
	  | |  | |__  | |    | |      | || |     
	  | |  |  __| | |    | |      | || |     
	  | |  | |____| |____| |____ _| || |____ 
	  |_|  |______|\_____|______|_____\_____|
`

	c.Println(logo)
	c.Println("🌱 tfcli - Terraform 프로젝트를 손쉽게 관리하는 CLI 도구")
	fmt.Println()
	fmt.Println("🧑‍💻 만든 사람: Gisang Lee (https://github.com/GisangLee)")
	fmt.Println()
}
