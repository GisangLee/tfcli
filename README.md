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
curl -L https://github.com/GisangLee/tfcli/releases/download/v1.1.0/tfcli-darwin-amd64-v1.0.0 -o tfcli
```

### 🔹 macOS (Intel)

```bash
curl -L https://github.com/GisangLee/tfcli/releases/download/v1.1.0/tfcli-darwin-arm64-v1.0.0 -o tfcli
```


### 🔹 Linux

```bash
curl -L https://github.com/GisangLee/tfcli/releases/download/v1.1.0/tfcli-linux-amd64-v1.0.0 -o tfcli
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
```bash
$ tfcli

╔══════════════════════════════════════════════════════╗
║          🚀 Terraform Automation CLI v1.0            ║
║      CSP 환경을 손쉽게 선택하고 작업하세요!      ║
║     Made with ☕️ & 💻 by DevOps JSON      ║
╚══════════════════════════════════════════════════════╝

✨ 시작해볼까요?

Use the arrow keys to navigate: ↓ ↑ → ← 
? ✨ 실행할 작업을 선택하세요: 
  ▸ create-project
    create-template
    tf-job 
```

> create-project<br>- CSP 및 환경(dev/stage/prod)에 따른 디렉토리 자동 생성


> create-template<br>- vpc, alb, eks 등 리소스별 템플릿 생성


> tf-job<br>- Terraform 명령어 실행 (init, plan, apply, destroy 등)

### 생성되는 디렉토리 예시
```
aws/
├── modules/
└── environment/
    ├── dev/
    ├── stage/
    └── prod/
└── modules/
    ├── vpc/
```
