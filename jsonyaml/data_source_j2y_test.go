package jsonyaml

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testDataSourceJ2YConfig_basic = `
provider "yamldecode" {}

data "yamldecode_decode" "foo" {
  input =<<EOS
---
	true
EOS
}
`

func TestDataSourceJ2Y_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceJ2YConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.yamldecode_decode.foo", "boolean", "true"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string_array"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "object"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "number"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string"),
				),
			},
		},
	})
}