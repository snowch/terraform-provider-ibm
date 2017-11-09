package ibm

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMNetworkInterfaceSGAttachment(t *testing.T) {
	hostname := acctest.RandString(16)

	configInstance := "ibm_compute_vm_instance.tfuatvmwithgroups"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIBMComputeVmInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccTestAccIBMNetworkInterfaceSGAttachmentConfig(hostname),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_interface_sg_attachment", "security_group_id", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "public_security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(
						configInstance, "private_security_group_ids.#", "1"),
				),
			},
		},
	})
}

func testAccTestAccIBMNetworkInterfaceSGAttachmentConfig(hostname string) string {
	v := fmt.Sprintf(`
		data "ibm_security_group" "allowssh" {
			name        = "allow_ssh"
		  } 
		  resource "ibm_compute_vm_instance" "tfuatvm" {
			hostname                   = "%s"
			domain                     = "tfvmuatsg.com"
			os_reference_code          = "DEBIAN_7_64"
			datacenter                 = "wdc07"
			network_speed              = 10
			hourly_billing             = true
			private_network_only       = false
			cores                      = 1
			memory                     = 1024
			disks                      = [25, 10, 20]
			dedicated_acct_host_only   = true
			local_disk                 = false
			ipv6_enabled               = true
			secondary_ip_count         = 4
			notes                      = "VM notes"
		  }
		 
		  rsource "ibm_network_interface_sg_attachment"  "ifcsg" {
			security_group_id = "{data.ibm_security_group.allowssh.id}"
			network_interface_id = "${ibm_compute_vm_instance.tfuatvm.public_interface_id}"
		  }
		  `, hostname)
	fmt.Println(v)
	os.Exit(1)
	return v
}
