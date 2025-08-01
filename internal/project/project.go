package project

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/GisangLee/tfcli/internal/utils"
)

func HandleCreateProject() {
	cspPrompt, _ := utils.PromptCSP()
	createProjectStructure(cspPrompt)
}

func createProjectStructure(csp string) {
	root := filepath.Join(".", csp)
	modules := filepath.Join(root, "modules")
	envRoot := filepath.Join(root, "environment")
	envs := []string{"dev", "stage", "prod"}

	// 디렉토리 생성
	dirs := append([]string{modules}, func() []string {
		result := []string{}
		for _, e := range envs {
			result = append(result, filepath.Join(envRoot, e))
		}
		return result
	}()...)

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("❌ 디렉토리 생성 실패: %s\n", dir)
		} else {
			fmt.Printf("✅ 디렉토리 생성됨: %s\n", dir)
		}
	}

	// 각 환경별 파일 생성
	for _, env := range envs {
		envPath := filepath.Join(envRoot, env)

		createEnvFile(envPath, "main.tf", `// main.tf
// 이 파일에 해당 환경에서 사용할 모듈을 참조하세요.
`)

		switch csp {
		case "aws":
			createEnvFile(envPath, "providers.tf", `provider "aws" {
  region  = var.aws_region
  profile = var.aws_profile
}

terraform {
  backend "local" {
    path = "state/terraform.tfstate"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~>6.0"
    }
  }
}`)

			createEnvFile(envPath, "variables.tf", `// variables.tf
// 환경에 따라 필요한 변수를 정의하세요.

variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "ap-northeast-2"
}

variable "aws_profile" {
  description = "AWS CLI profile name"
  type        = string
}
`)

			tfvars := fmt.Sprintf(`# terraform.tfvars
# 실제 값들을 설정하세요
aws_region  = "ap-northeast-2"
aws_profile = "%s"
`, env)
			createEnvFile(envPath, "terraform.tfvars", tfvars)

		case "ncp":
			createEnvFile(envPath, "providers.tf", `terraform {
  required_providers {
    ncloud = {
      source = "NaverCloudPlatform/ncloud"
    }
  }
  required_version = ">= 0.13"
}

provider "ncloud" {
  access_key  = var.access_key
  secret_key  = var.secret_key
  region      = var.region
  site        = var.site
  support_vpc = true
}`)

			createEnvFile(envPath, "variables.tf", `// variables.tf
// NCP 자격증명 변수 정의

variable "access_key" {
  type = string
}
variable "secret_key" {
  type = string
}
variable "region" {
  type = string
}
variable "site" {
  type = string
}
`)

			createEnvFile(envPath, "terraform.tfvars", `# terraform.tfvars
# 실제 값들을 설정하세요
access_key = "REPLACE_ME"
secret_key = "REPLACE_ME"
region     = "KR"
site       = "public"
`)

		default: // GCP 또는 기타 CSP는 기본 스켈레톤만
			createEnvFile(envPath, "variables.tf", `// variables.tf
// 필요한 변수를 여기에 정의하세요.
`)
			createEnvFile(envPath, "terraform.tfvars", `# terraform.tfvars
# 실제 값들을 설정하세요
`)
		}
	}

	fmt.Println("📁 프로젝트 기본 구조가 생성되었습니다.")
}

func createEnvFile(dir, filename, content string) {
	fullPath := filepath.Join(dir, filename)
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Printf("⚠️  이미 존재함 (건너뜀): %s\n", fullPath)
		return
	}
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		fmt.Printf("❌ 파일 생성 실패: %s (%v)\n", fullPath, err)
	} else {
		fmt.Printf("📄 생성됨: %s\n", fullPath)
	}
}
