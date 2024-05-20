package dataplex_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDataplexAspectType_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataplexAspectTypeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexAspectType_full(context),
			},
			{
				ResourceName:            "google_dataplex_aspect_type.test_aspect_type_full",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"aspect_type_id", "labels", "location", "terraform_labels"},
			},
			{
				Config: testAccDataplexAspectType_update(context),
			},
			{
				ResourceName:            "google_dataplex_aspect_type.test_aspect_type_basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"aspect_type_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccDataplexAspectType_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_aspect_type" "test_aspect_type_full" {
  aspect_type_id = "tf-test-aspect-type-full%{random_suffix}"
  project = "%{project_name}"
  location = "us-central1"

  labels = { "tag": "test-tf" }
  display_name = "terraform aspect type"
  description = "aspect type created by Terraform"
  metadata_template = <<EOF
{
  "type": "record",
  "name": "Schema",
  "recordFields": [
    {
      "name": "fields",
      "type": "array",
      "index": 1,
      "arrayItems": {
        "name": "field",
        "type": "record",
        "typeId": "field",
        "recordFields": [
          {
            "name": "name",
            "type": "string",
            "index": 1,
            "constraints": {
              "required": true
            }
          },
          {
            "name": "description",
            "type": "string",
            "index": 2
          },
          {
            "name": "dataType",
            "type": "string",
            "index": 3,
            "constraints": {
              "required": true
            }
          },
          {
            "name": "metadataType",
            "type": "enum",
            "index": 4,
            "constraints": {
              "required": true
            },
            "enumValues": [
              {
                "name": "BOOLEAN",
                "index": 1
              },
              {
                "name": "NUMBER",
                "index": 2
              },
              {
                "name": "STRING",
                "index": 3
              },
              {
                "name": "BYTES",
                "index": 4
              },
              {
                "name": "DATETIME",
                "index": 5
              },
              {
                "name": "TIMESTAMP",
                "index": 6
              },
              {
                "name": "GEOSPATIAL",
                "index": 7
              },
              {
                "name": "STRUCT",
                "index": 8
              },
              {
                "name": "OTHER",
                "index": 100
              }
            ]
          },
          {
            "name": "mode",
            "type": "enum",
            "index": 5,
            "enumValues": [
              {
                "name": "NULLABLE",
                "index": 1
              },
              {
                "name": "REPEATED",
                "index": 2
              },
              {
                "name": "REQUIRED",
                "index": 3
              }
            ]
          },
          {
            "name": "defaultValue",
            "type": "string",
            "index": 6
          },
          {
            "name": "annotations",
            "type": "map",
            "index": 7,
            "mapItems": {
              "name": "label",
              "type": "string"
            }
          },
          {
            "name": "fields",
            "type": "array",
            "index": 20,
            "arrayItems": {
              "name": "field",
              "type": "record",
              "typeRef": "field"
            }
          }
        ]
      }
    }
  ]
}
EOF
}
`, context)
}

func testAccDataplexAspectType_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_aspect_type" "test_aspect_type_basic" {
  aspect_type_id = "tf-test-aspect-type-basic%{random_suffix}"
  project = "%{project_name}"
  location = "us-central1"

  metadata_template = <<EOF
{
  "name": "tf-test-template",
  "type": "record",
  "recordFields": [
    {
      "name": "type",
      "type": "enum",
      "annotations": {
        "displayName": "Type",
        "description": "Specifies the type of view represented by the entry."
      },
      "index": 1,
      "constraints": {
        "required": true
      },
      "enumValues": [
        {
          "name": "VIEW",
          "index": 1
        }
      ]
    }
  ]
}
EOF
}
`, context)
}
