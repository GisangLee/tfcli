package utils

func ParseTerraformArgs(task string) []string {
	switch task {
	case "init":
		return []string{"init", "-no-color"}
	case "fmt -recursive":
		return []string{"fmt", "-recursive", "-no-color"}
	case "plan":
		return []string{"plan", "-var-file=terraform.tfvars", "-no-color"}
	case "apply":
		return []string{"apply", "-auto-approve", "-no-color"}
	case "destroy":
		return []string{"destroy", "-auto-approve", "-no-color"}
	default:
		return []string{}
	}
}
