package test

import (
  "fmt"
  "os"
  "testing"
)

func TestTerraformPlan(t *testing.T) {
  t.Parallel()
  tests := map[string]PlanTestCase{
    "typical": {
      TestCase{
        Create: false,
      },
    },
  }

  for name, tc := range tests {
    name := name
    tc := tc
    t.Run(name, func(t *testing.T) {
      t.Parallel()
        cwd, err := os.Getwd()
        if err != nil {
          t.Fatal(err)
        }

        testingDir := fmt.Sprintf("%v/.terraform_tests/plan/%v", cwd, name)

    })
  }
}
