package tests

import (
  "github.com/gruntwork-io/terratest/modules/terraform"
  "github.com/gruntwork-io/terratest/modules/test-structure"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestNamingModuleDefault(t *testing.T) {
  t.Parallel()

  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../../", "examples/default")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, "prefix-resource_name-suffix", rendered)

  })
}
func TestNamingModuleWithoutSuffix(t *testing.T) {
  t.Parallel()

  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../../", "examples/without_suffix")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Regexp(t,"^prefix-resource_name-.*$", rendered)
    //assert.Equal(t, "prefix-resource_name", rendered)

  })
}
func TestNamingModuleWithoutPrefix(t *testing.T) {
  t.Parallel()

  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../../", "examples/without_prefix")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, "resource_name-suffix", rendered)

  })
}
func TestNamingModuleWithoutResourceName(t *testing.T) {
  t.Parallel()

  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../../", "examples/without_resource_name")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, "prefix-suffix", rendered)

  })
}
func TestNamingModuleLowerCase(t *testing.T) {
  t.Parallel()

  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../../", "examples/lower")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, "prefix-resource_name-suffix", rendered)

  })
}
func TestNamingModuleUpperCase(t *testing.T) {
  t.Parallel()

  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../../", "examples/upper")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, "PREFIX-RESOURCE_NAME-SUFFIX", rendered)

  })
}
func TestNamingModuleSeparator(t *testing.T) {
  t.Parallel()

  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../../", "examples/separator")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, "prefix_resource_name_suffix", rendered)

  })
}
func TestNamingModuleLength(t *testing.T) {
  t.Parallel()

  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../../", "examples/length")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, "prefix-resource", rendered)

  })
}
