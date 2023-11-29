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

package cloudbuild_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccCloudBuildBitbucketServerConfig_cloudbuildBitbucketServerConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudBuildBitbucketServerConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudBuildBitbucketServerConfig_cloudbuildBitbucketServerConfigExample(context),
			},
			{
				ResourceName:            "google_cloudbuild_bitbucket_server_config.bbs-config",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"config_id", "location"},
			},
		},
	})
}

func testAccCloudBuildBitbucketServerConfig_cloudbuildBitbucketServerConfigExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuild_bitbucket_server_config" "bbs-config" {
    config_id = "tf-test-bbs-config%{random_suffix}"
    location = "us-central1"
    host_uri = "https://bbs.com"
    secrets {
        admin_access_token_version_name = "projects/myProject/secrets/mybbspat/versions/1"
        read_access_token_version_name = "projects/myProject/secrets/mybbspat/versions/1"
        webhook_secret_version_name = "projects/myProject/secrets/mybbspat/versions/1"
    }
    username = "test"
    api_key = "<api-key>"
}
`, context)
}

func TestAccCloudBuildBitbucketServerConfig_cloudbuildBitbucketServerConfigPeeredNetworkExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedTestNetwork(t, "peered-network"),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudBuildBitbucketServerConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudBuildBitbucketServerConfig_cloudbuildBitbucketServerConfigPeeredNetworkExample(context),
			},
			{
				ResourceName:            "google_cloudbuild_bitbucket_server_config.bbs-config-with-peered-network",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"config_id", "location"},
			},
		},
	})
}

func testAccCloudBuildBitbucketServerConfig_cloudbuildBitbucketServerConfigPeeredNetworkExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {}

resource "google_project_service" "servicenetworking" {
  service = "servicenetworking.googleapis.com"
  disable_on_destroy = false
}
 
data "google_compute_network" "vpc_network" {
    name       = "%{network_name}"
    depends_on = [google_project_service.servicenetworking]
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          = "tf-test-private-ip-alloc%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = data.google_compute_network.vpc_network.id
}

resource "google_service_networking_connection" "default" {
  network                 = data.google_compute_network.vpc_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
  depends_on              = [google_project_service.servicenetworking]
}

resource "google_cloudbuild_bitbucket_server_config" "bbs-config-with-peered-network" {
    config_id = "tf-test-bbs-config%{random_suffix}"
    location = "us-central1"
    host_uri = "https://bbs.com"
    secrets {
        admin_access_token_version_name = "projects/myProject/secrets/mybbspat/versions/1"
        read_access_token_version_name = "projects/myProject/secrets/mybbspat/versions/1"
        webhook_secret_version_name = "projects/myProject/secrets/mybbspat/versions/1"
    }
    username = "test"
    api_key = "<api-key>"
    peered_network = replace(data.google_compute_network.vpc_network.id, data.google_project.project.name, data.google_project.project.number)
    ssl_ca = "-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----\n"
    depends_on = [google_service_networking_connection.default]
}
`, context)
}

func testAccCheckCloudBuildBitbucketServerConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloudbuild_bitbucket_server_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{CloudBuildBasePath}}projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("CloudBuildBitbucketServerConfig still exists at %s", url)
			}
		}

		return nil
	}
}
