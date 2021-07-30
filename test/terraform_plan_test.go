package test

import (
	"fmt"
	"github.com/greenexpomarker88/terraform-aws-s3-bucket/pkg/terraform"
	terratest "github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTerraformPlan(t *testing.T) {
	t.Parallel()
	tests := map[string]TestCase{
		"no_create": {
			Create: false,
			Bucket: "my-bucket-123",
		},
		"typical": {
			Create: true,
			Bucket: "my-bucket-789",
			Acl:    "private",
			Tags: map[string]string{
				"env":        "staging",
				"managed-by": "terraform",
			},
			ServerSideEncryptionEnabled:   true,
			ServerSideEncryptionAlgorithm: "aws:kms",
			VersioningEnabled:             true,
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			projRoot := ".."
			exampleDirectoryRelativeToRoot := "examples/complete"
			testDirectory := test_structure.CopyTerraformFolderToTemp(t, projRoot, exampleDirectoryRelativeToRoot)

			terraformOptions := &terratest.Options{
				TerraformDir: testDirectory,
				PlanFilePath: fmt.Sprintf("%v/tfplan", testDirectory),
				Vars: map[string]interface{}{
					"create_bucket":      tc.Create,
					"bucket":             tc.Bucket,
					"tags":               tc.Tags,
					"versioning_enabled": tc.VersioningEnabled,
				},
			}

			test_structure.SaveTerraformOptions(t, testDirectory, terraformOptions)
			terratest.InitAndPlan(t, terraformOptions)

			show := terratest.Show(t, terraformOptions)
			repo := terraform.NewRepository(show)
			bucketModule, err := repo.GetModule("module.my-bucket")

			t.Run("module creation", func(t *testing.T) {
				t.Parallel()
				if tc.Create {
					assert.NotNil(t, bucketModule)
					assert.Nil(t, err)
				} else {
					assert.Nil(t, bucketModule)
					assert.NotNil(t, err)
				}
			})

			t.Run("module tags", func(t *testing.T) {
				t.Parallel()
				for k, v := range tc.Tags {
					for _, r := range bucketModule.GetTaggableResources() {
						var tags map[string]interface{}
						if r.AttributeValues["tags"] == nil {
							tags = map[string]interface{}{}
						} else {
							tags = r.AttributeValues["tags"].(map[string]interface{})
						}

						val, ok := tags[k]
						assert.True(t, ok && val == v, "resource %v should have a tag: key=%v value=%v", r.Address, k, v)
					}
				}
			})
		})
	}
}
