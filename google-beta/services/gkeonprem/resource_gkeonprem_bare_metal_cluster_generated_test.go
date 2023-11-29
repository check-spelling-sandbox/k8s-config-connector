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

package gkeonprem_test

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

func TestAccGkeonpremBareMetalCluster_gkeonpremBareMetalClusterBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGkeonpremBareMetalClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeonpremBareMetalCluster_gkeonpremBareMetalClusterBasicExample(context),
			},
			{
				ResourceName:            "google_gkeonprem_bare_metal_cluster.cluster-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccGkeonpremBareMetalCluster_gkeonpremBareMetalClusterBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gkeonprem_bare_metal_cluster" "cluster-basic" {
  provider = google-beta
  name = "tf-test-my-cluster%{random_suffix}"
  location = "us-west1"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  bare_metal_version = "1.12.3"
  network_config {
    island_mode_cidr {
      service_address_cidr_blocks = ["172.26.0.0/16"]
      pod_address_cidr_blocks = ["10.240.0.0/13"]
    }
  }
  control_plane {
    control_plane_node_pool_config {
      node_pool_config {
        labels = {}
        operating_system = "LINUX"
        node_configs {
          labels = {}
          node_ip = "10.200.0.9"
        }
      }
    }
  }
  load_balancer {
    port_config {
      control_plane_load_balancer_port = 443
    }
    vip_config {
      control_plane_vip = "10.200.0.13"
      ingress_vip = "10.200.0.14"
    }
    metal_lb_config {
      address_pools {
        pool = "pool1"
        addresses = [
          "10.200.0.14/32",
          "10.200.0.15/32",
          "10.200.0.16/32",
          "10.200.0.17/32",
          "10.200.0.18/32",
          "fd00:1::f/128",
          "fd00:1::10/128",
          "fd00:1::11/128",
          "fd00:1::12/128"
        ]
        avoid_buggy_ips = true
        manual_assign = true
      }
    }
  }
  storage {
    lvp_share_config {
      lvp_config {
        path = "/mnt/localpv-share"
        storage_class = "local-shared"
      }
      shared_path_pv_count = 5
    }
    lvp_node_mounts_config {
      path = "/mnt/localpv-disk"
      storage_class = "local-disks"
    }
  }
  security_config {
    authorization {
      admin_users {
        username = "admin@hashicorptest.com"
      }
    }
  }
}
`, context)
}

func TestAccGkeonpremBareMetalCluster_gkeonpremBareMetalClusterManuallbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGkeonpremBareMetalClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeonpremBareMetalCluster_gkeonpremBareMetalClusterManuallbExample(context),
			},
			{
				ResourceName:            "google_gkeonprem_bare_metal_cluster.cluster-manuallb",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccGkeonpremBareMetalCluster_gkeonpremBareMetalClusterManuallbExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gkeonprem_bare_metal_cluster" "cluster-manuallb" {
  provider = google-beta
  name = "tf-test-cluster-manuallb%{random_suffix}"
  location = "us-west1"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  bare_metal_version = "1.12.3"
  network_config {
    island_mode_cidr {
      service_address_cidr_blocks = ["172.26.0.0/16"]
      pod_address_cidr_blocks = ["10.240.0.0/13"]
    }
  }
  control_plane {
    control_plane_node_pool_config {
      node_pool_config {
        labels = {}
        operating_system = "LINUX"
        node_configs {
          labels = {}
          node_ip = "10.200.0.9"
        }
      }
    }
  }
  load_balancer {
    port_config {
      control_plane_load_balancer_port = 443
    }
    vip_config {
      control_plane_vip = "10.200.0.13"
      ingress_vip = "10.200.0.14"
    }
    manual_lb_config {
      enabled = true
    }
  }
  storage {
    lvp_share_config {
      lvp_config {
        path = "/mnt/localpv-share"
        storage_class = "local-shared"
      }
      shared_path_pv_count = 5
    }
    lvp_node_mounts_config {
      path = "/mnt/localpv-disk"
      storage_class = "local-disks"
    }
  }
  security_config {
    authorization {
      admin_users {
        username = "admin@hashicorptest.com"
      }
    }
  }
  binary_authorization {
    evaluation_mode = "DISABLED"
  }
  upgrade_policy {
    policy = "SERIAL"
  }
}
`, context)
}

func TestAccGkeonpremBareMetalCluster_gkeonpremBareMetalClusterBgplbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGkeonpremBareMetalClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeonpremBareMetalCluster_gkeonpremBareMetalClusterBgplbExample(context),
			},
			{
				ResourceName:            "google_gkeonprem_bare_metal_cluster.cluster-bgplb",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccGkeonpremBareMetalCluster_gkeonpremBareMetalClusterBgplbExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gkeonprem_bare_metal_cluster" "cluster-bgplb" {
  provider = google-beta
  name = "tf-test-cluster-bgplb%{random_suffix}"
  location = "us-west1"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  bare_metal_version = "1.12.3"
  network_config {
    island_mode_cidr {
      service_address_cidr_blocks = ["172.26.0.0/16"]
      pod_address_cidr_blocks = ["10.240.0.0/13"]
    }
    advanced_networking = true
    multiple_network_interfaces_config {
      enabled = true
    }
    sr_iov_config {
      enabled = true
    }
  }
  control_plane {
    control_plane_node_pool_config {
      node_pool_config {
        labels = {}
        operating_system = "LINUX"
        node_configs {
          labels = {}
          node_ip = "10.200.0.9"
        }
        taints {
          key = "test-key"
          value = "test-value"
          effect = "NO_EXECUTE"
        }
      }
    }
    api_server_args {
      argument = "test-argument"
      value = "test-value"
    }
  }
  load_balancer {
    port_config {
      control_plane_load_balancer_port = 443
    }
    vip_config {
      control_plane_vip = "10.200.0.13"
      ingress_vip = "10.200.0.14"
    }
    bgp_lb_config {
      asn = 123456
      bgp_peer_configs {
        asn = 123457
        ip_address = "10.0.0.1"
        control_plane_nodes = ["test-node"]
      }
      address_pools {
        pool = "pool1"
        addresses = [
          "10.200.0.14/32",
          "10.200.0.15/32",
          "10.200.0.16/32",
          "10.200.0.17/32",
          "10.200.0.18/32",
          "fd00:1::f/128",
          "fd00:1::10/128",
          "fd00:1::11/128",
          "fd00:1::12/128"
        ]
      }
      load_balancer_node_pool_config {
        node_pool_config {
          labels = {}
          operating_system = "LINUX"
          node_configs {
            labels = {}
            node_ip = "10.200.0.9"
          }
          taints {
            key = "test-key"
            value = "test-value"
            effect = "NO_EXECUTE"
          }
          kubelet_config {
            registry_pull_qps = 10
            registry_burst = 12
            serialize_image_pulls_disabled = true
          }
        }
      }
    }
  }
  storage {
    lvp_share_config {
      lvp_config {
        path = "/mnt/localpv-share"
        storage_class = "local-shared"
      }
      shared_path_pv_count = 5
    }
    lvp_node_mounts_config {
      path = "/mnt/localpv-disk"
      storage_class = "local-disks"
    }
  }
  security_config {
    authorization {
      admin_users {
        username = "admin@hashicorptest.com"
      }
    }
  }
  proxy {
    uri = "http://test-domain/test"
    no_proxy = ["127.0.0.1"]
  }
  cluster_operations {
    enable_application_logs = true
  }
  maintenance_config {
    maintenance_address_cidr_blocks = ["192.168.0.1/20"] 
  }
  node_config {
    max_pods_per_node = 10
    container_runtime = "CONTAINERD"
  }
  node_access_config {
    login_user = "test@example.com"
  }
  os_environment_config {
    package_repo_excluded = true
  }
}
`, context)
}

func testAccCheckGkeonpremBareMetalClusterDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_gkeonprem_bare_metal_cluster" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{GkeonpremBasePath}}projects/{{project}}/locations/{{location}}/bareMetalClusters/{{name}}")
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
				return fmt.Errorf("GkeonpremBareMetalCluster still exists at %s", url)
			}
		}

		return nil
	}
}
