package test

import (
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformProdCreatesLocalDirectoryAndFile(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "/home/gocheme/Desktop/mock-project/MyTerragruntProject/mock/prod",
		NoColor:      true,
		Vars: map[string]interface{}{
			"environment": "prod",
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Deploy using Terraform
	terraform.InitAndApply(t, terraformOptions)

	expectedDir := "/home/gocheme/Desktop/mock-project/MyTerragruntProject/mock/prod"
	expectedFile := filepath.Join(expectedDir, "sample-config-prod.conf")

	// Verify if the directory and file exist
	assert.DirExists(t, expectedDir)
	assert.FileExists(t, expectedFile)
}
