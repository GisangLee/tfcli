# 🚀 tfcli

Terraform 기반 인프라 작업을 자동화하는 경량 CLI 도구입니다.

- CSP (AWS / NCP / GCP) 별 프로젝트 구조 생성
- 리소스 템플릿 생성
- 터미널 기반 Terraform 명령 실행 (`init`, `plan`, `apply`, `destroy`)

---

## 📦 다운로드

아래 명령어로 플랫폼에 맞게 바이너리를 다운로드하세요:

### 🔹 macOS (Apple Silicon)

```bash
curl -L https://github.com/GisangLee/tfcli/releases/download/v1.2.0/tfcli-darwin-amd64-v1.2.0 -o tfcli
```

### 🔹 macOS (Intel)

```bash
curl -L https://github.com/GisangLee/tfcli/releases/download/v1.2.0/tfcli-darwin-arm64-v1.2.0 -o tfcli
```


### 🔹 Linux

```bash
curl -L https://github.com/GisangLee/tfcli/releases/download/v1.2.0/tfcli-linux-amd64-v1.2.0 -o tfcli
```

---
## 📦 설치
```bash
chmod +x tfcli
sudo mv tfcli /usr/local/bin/tfcli
```

---

## 📦 사용법
> 터미널에서 아래 명렁어로 실행하세요.

> ### tfcli OR tfcli --help
```bash
$ tfcli

        _______ ______  _____ _      _____ _____ 
   |__   __|  ____|/ ____| |    |_   _/ ____|
          | |  | |__  | |    | |      | || |     
          | |  |  __| | |    | |      | || |     
          | |  | |____| |____| |____ _| || |____ 
          |_|  |______|\_____|______|_____\_____|

🌱 tfcli - Terraform 프로젝트를 손쉽게 관리하는 CLI 도구

🧑‍💻 만든 사람: Gisang Lee (https://github.com/GisangLee)

╭────────────────────────────────────────────────────────╮
│ 🌱 tfcli - Terraform 프로젝트를 손쉽게 관리하는 CLI 도구  │
╰────────────────────────────────────────────────────────╯

🛠 사용법:
  tfcli [명령어] [옵션]

📚 사용 가능한 명령어:
  project     📁 TFCLI 프로젝트 구조 생성
  template    🧩 모듈 템플릿 생성
  module      📦 모듈 자동 참조
  tf          🚀 Terraform 작업 실행 (init/fmt --recursive/plan/apply/destroy)

🔧 옵션:
  -h, --help   도움말 출력
```


> ### tfcli project --help
```bash
$ tfcli project --help
╭──────────────────────────────────────────────╮
│ 📁 project - Terraform 프로젝트 구조 생성     │
╰──────────────────────────────────────────────╯

🧭 설명:
  AWS / NCP / GCP 환경의 Terraform 프로젝트 디렉토리를 빠르게 생성합니다.

🛠 사용법:
  tfcli project

📂 생성 구조:
  [csp]/
    ├─ modules/
    └─ environment/
         ├─ dev/
         ├─ stage/
         └─ prod/

🔧 옵션:
  -h, --help   도움말 출력
```

> ### tfcli template --help
```bash
$ tfcli template --help

╭──────────────────────────────────────────────╮
│ 🧩 template - Terraform 모듈 템플릿 생성      │
╰──────────────────────────────────────────────╯

🧭 설명:
  원하는 이름과 리소스 종류에 맞춰 템플릿을 생성합니다.

🛠 사용법:
  tfcli template

📁 생성되는 파일:
  main.tf, variables.tf, outputs.tf

🔧 옵션:
  -h, --help   도움말 출력
```

> ### tfcli module --help
```bash
$ tfcli module --help

╭────────────────────────────────────────────────────────╮
│ 📦 tfcli module - 모듈 자동 참조 도우미                          │
╰────────────────────────────────────────────────────────╯

🛠 사용법:
  tfcli module add [모듈명] --source=경로

📚 예시:
  tfcli module add vpc --source=./modules/vpc

🔧 옵션:
  -h, --help   도움말 출력
```

> ### tfcli tf --help
```bash
 $ tfcli tf --help
╭──────────────────────────────────────────────╮
│ 🚀 tf - Terraform 작업 실행                    │
╰──────────────────────────────────────────────╯

🧭 설명:
  Terraform의 주요 명령어를 선택적으로 실행합니다.

🛠 사용법:
  tfcli tf

⚙️ 실행 가능한 작업:
  init, fmt -recursive, plan, apply, destroy

🔧 옵션:
  -h, --help   도움말 출력
```