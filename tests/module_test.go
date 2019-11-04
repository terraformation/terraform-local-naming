package test

import (
  "fmt"
  "github.com/gruntwork-io/terratest/modules/random"
  "github.com/gruntwork-io/terratest/modules/terraform"
  test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
  "github.com/stretchr/testify/assert"
  "strings"
  "testing"
)

func TestNamingModuleWithSuffix(t *testing.T) {
  t.Parallel()

  // A unique ID we can use to namespace resources so we don't clash with anything
  // already in the AWS account or tests running in parallel
  uniqueID := random.UniqueId()
  namingOptions := map[string]string{"suffix": uniqueID}

  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../", "exemples/default")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,

      // Variables to pass to our Terraform code using -var options
      Vars: map[string]interface{}{
        "naming_options": namingOptions,
      },
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, strings.ToLower(uniqueID), rendered)

  })
}

func TestNamingModuleWithPrefix(t *testing.T) {
  t.Parallel()

  // A unique ID we can use to namespace resources so we don't clash with anything
  // already in the AWS account or tests running in parallel
  uniqueID := random.UniqueId()
  namingOptions := map[string]string{"prefix": uniqueID}

  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../", "exemples/default")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,

      // Variables to pass to our Terraform code using -var options
      Vars: map[string]interface{}{
        "naming_options": namingOptions,
      },
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, strings.ToLower(uniqueID), rendered)

  })
}

func TestNamingModuleWithResourceName(t *testing.T) {
  t.Parallel()

  // A unique ID we can use to namespace resources so we don't clash with anything
  // already in the AWS account or tests running in parallel
  uniqueID := random.UniqueId()
  namingOptions := map[string]string{"resource_name": uniqueID}
  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../", "exemples/default")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,

      // Variables to pass to our Terraform code using -var options
      Vars: map[string]interface{}{
        "naming_options": namingOptions,
      },
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, strings.ToLower(uniqueID), rendered)
  })
}

func TestNamingModuleWithAll(t *testing.T) {
  t.Parallel()

  // A unique ID we can use to namespace resources so we don't clash with anything
  // already in the AWS account or tests running in parallel
  resourceName := random.UniqueId()
  suffix := random.UniqueId()
  prefix := random.UniqueId()
  namingOptions := map[string]string{"resource_name": resourceName, "suffix": suffix, "prefix": prefix}
  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../", "exemples/default")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,

      // Variables to pass to our Terraform code using -var options
      Vars: map[string]interface{}{
        "naming_options": namingOptions,
      },
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, strings.ToLower(fmt.Sprintf("%s-%s-%s", prefix, resourceName, suffix)), rendered)
  })
}

func TestNamingModuleWithoutHyphen(t *testing.T) {
  t.Parallel()

  // A unique ID we can use to namespace resources so we don't clash with anything
  // already in the AWS account or tests running in parallel
  resourceName := random.UniqueId()
  suffix := random.UniqueId()
  prefix := random.UniqueId()
  namingOptions := map[string]string{"resource_name": resourceName, "suffix": suffix, "prefix": prefix, "hyphen": "false"}
  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../", "exemples/default")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,

      // Variables to pass to our Terraform code using -var options
      Vars: map[string]interface{}{
        "naming_options": namingOptions,
      },
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, strings.ToLower(fmt.Sprintf("%s%s%s", prefix , resourceName, suffix)), rendered)
  })
}

func TestNamingModuleUpper(t *testing.T) {
  t.Parallel()

  // A unique ID we can use to namespace resources so we don't clash with anything
  // already in the AWS account or tests running in parallel
  resourceName := random.UniqueId()
  suffix := random.UniqueId()
  prefix := random.UniqueId()
  namingOptions := map[string]string{"resource_name": resourceName, "suffix": suffix, "prefix": prefix, "lower": "false"}
  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../", "exemples/default")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,

      // Variables to pass to our Terraform code using -var options
      Vars: map[string]interface{}{
        "naming_options": namingOptions,
      },
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, strings.ToUpper(fmt.Sprintf("%s-%s-%s", prefix, resourceName, suffix)), rendered)
  })
}

func TestNamingModuleWhithoutSpace(t *testing.T) {
  t.Parallel()

  // A unique ID we can use to namespace resources so we don't clash with anything
  // already in the AWS account or tests running in parallel
  resourceName := random.UniqueId()
  suffix := random.UniqueId()
  prefix := random.UniqueId()
  namingOptions := map[string]string{"resource_name": resourceName, "suffix": fmt.Sprintf("%s ",suffix), "prefix": prefix}
  terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../", "exemples/default")

  defer test_structure.RunTestStage(t, "teardown", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    terraform.Destroy(t, terraformOptions)
  })

  test_structure.RunTestStage(t, "setup", func() {
    terraformOptions := &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: terraformDir,

      // Variables to pass to our Terraform code using -var options
      Vars: map[string]interface{}{
        "naming_options": namingOptions,
      },
    }
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

  })

  test_structure.RunTestStage(t, "validate", func() {
    terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
    rendered := terraform.OutputRequired(t, terraformOptions, "rendered")

    assert.Equal(t, strings.ToLower(fmt.Sprintf("%s-%s-%s", prefix, resourceName, suffix)), rendered)
  })
}
