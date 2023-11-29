// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package compute_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
)

func TestAccComputeImageIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/compute.imageUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImageIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_image_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s roles/compute.imageUser", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccComputeImageIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_compute_image_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s roles/compute.imageUser", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeImageIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/compute.imageUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeImageIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_image_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s roles/compute.imageUser user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeImageIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/compute.imageUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImageIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_compute_image_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_compute_image_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeImageIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_compute_image_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeImageIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/compute.imageUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImageIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_image_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s roles/compute.imageUser %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeImageIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/compute.imageUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImageIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_image_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s roles/compute.imageUser", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_image_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s roles/compute.imageUser %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_image_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s roles/compute.imageUser %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeImageIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/compute.imageUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImageIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_image_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s roles/compute.imageUser user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeImageIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/compute.imageUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImageIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_image_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s roles/compute.imageUser user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_image_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s roles/compute.imageUser user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_image_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s roles/compute.imageUser user:admin@hashicorptest.com %s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeImageIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           acctest.RandString(t, 10),
		"role":                    "roles/compute.imageUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := acctest.Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImageIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_compute_image_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", tpgresource.CheckGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_compute_image_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/global/images/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-image%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeImageIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }
}

resource "google_compute_image_iam_member" "foo" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeImageIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_compute_image_iam_policy" "foo" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_compute_image_iam_policy" "foo" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  depends_on = [
    google_compute_image_iam_policy.foo
  ]
}
`, context)
}

func testAccComputeImageIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }
}

data "google_iam_policy" "foo" {
}

resource "google_compute_image_iam_policy" "foo" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeImageIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }
}

resource "google_compute_image_iam_binding" "foo" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeImageIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }
}

resource "google_compute_image_iam_binding" "foo" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccComputeImageIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }
}

resource "google_compute_image_iam_binding" "foo" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccComputeImageIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }
}

resource "google_compute_image_iam_binding" "foo" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_compute_image_iam_binding" "foo2" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_compute_image_iam_binding" "foo3" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccComputeImageIamMember_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }
}

resource "google_compute_image_iam_member" "foo" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccComputeImageIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }
}

resource "google_compute_image_iam_member" "foo" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_compute_image_iam_member" "foo2" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_compute_image_iam_member" "foo3" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccComputeImageIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_compute_image_iam_policy" "foo" {
  project = google_compute_image.example.project
  image = google_compute_image.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
