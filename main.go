package main

func main() {
	showBanner()

	rootPrompt := selectRootMenu()
	switch rootPrompt {
	case "create-project":
		handleCreateProject()
	case "create-template":
		handleCreateTemplate()
	case "tf-job":
		handleTfJob()
	default:
		println("❌ 알 수 없는 작업입니다.")
	}
}
