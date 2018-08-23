//
// This file is automatically generated by scripts/generate-plugins.go -- Do not edit!
//

package command

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/packer/plugin"

	alicloudecsbuilder "github.com/hashicorp/packer/builder/alicloud/ecs"
	amazonchrootbuilder "github.com/hashicorp/packer/builder/amazon/chroot"
	amazonebsbuilder "github.com/hashicorp/packer/builder/amazon/ebs"
	amazonebssurrogatebuilder "github.com/hashicorp/packer/builder/amazon/ebssurrogate"
	amazonebsvolumebuilder "github.com/hashicorp/packer/builder/amazon/ebsvolume"
	amazoninstancebuilder "github.com/hashicorp/packer/builder/amazon/instance"
	azurearmbuilder "github.com/hashicorp/packer/builder/azure/arm"
	cloudstackbuilder "github.com/hashicorp/packer/builder/cloudstack"
	digitaloceanbuilder "github.com/hashicorp/packer/builder/digitalocean"
	dockerbuilder "github.com/hashicorp/packer/builder/docker"
	filebuilder "github.com/hashicorp/packer/builder/file"
	googlecomputebuilder "github.com/hashicorp/packer/builder/googlecompute"
	hypervisobuilder "github.com/hashicorp/packer/builder/hyperv/iso"
	hypervvmcxbuilder "github.com/hashicorp/packer/builder/hyperv/vmcx"
	lxcbuilder "github.com/hashicorp/packer/builder/lxc"
	lxdbuilder "github.com/hashicorp/packer/builder/lxd"
	ncloudbuilder "github.com/hashicorp/packer/builder/ncloud"
	nullbuilder "github.com/hashicorp/packer/builder/null"
	oneandonebuilder "github.com/hashicorp/packer/builder/oneandone"
	openstackbuilder "github.com/hashicorp/packer/builder/openstack"
	oracleclassicbuilder "github.com/hashicorp/packer/builder/oracle/classic"
	oracleocibuilder "github.com/hashicorp/packer/builder/oracle/oci"
	parallelsisobuilder "github.com/hashicorp/packer/builder/parallels/iso"
	parallelspvmbuilder "github.com/hashicorp/packer/builder/parallels/pvm"
	profitbricksbuilder "github.com/hashicorp/packer/builder/profitbricks"
	qemubuilder "github.com/hashicorp/packer/builder/qemu"
	scalewaybuilder "github.com/hashicorp/packer/builder/scaleway"
	tritonbuilder "github.com/hashicorp/packer/builder/triton"
	virtualboxisobuilder "github.com/hashicorp/packer/builder/virtualbox/iso"
	virtualboxovfbuilder "github.com/hashicorp/packer/builder/virtualbox/ovf"
	vmwareisobuilder "github.com/hashicorp/packer/builder/vmware/iso"
	vmwarevmxbuilder "github.com/hashicorp/packer/builder/vmware/vmx"
	alicloudimportpostprocessor "github.com/hashicorp/packer/post-processor/alicloud-import"
	amazonimportpostprocessor "github.com/hashicorp/packer/post-processor/amazon-import"
	artificepostprocessor "github.com/hashicorp/packer/post-processor/artifice"
	checksumpostprocessor "github.com/hashicorp/packer/post-processor/checksum"
	compresspostprocessor "github.com/hashicorp/packer/post-processor/compress"
	dockerimportpostprocessor "github.com/hashicorp/packer/post-processor/docker-import"
	dockerpushpostprocessor "github.com/hashicorp/packer/post-processor/docker-push"
	dockersavepostprocessor "github.com/hashicorp/packer/post-processor/docker-save"
	dockertagpostprocessor "github.com/hashicorp/packer/post-processor/docker-tag"
	googlecomputeexportpostprocessor "github.com/hashicorp/packer/post-processor/googlecompute-export"
	googlecomputeimportpostprocessor "github.com/hashicorp/packer/post-processor/googlecompute-import"
	manifestpostprocessor "github.com/hashicorp/packer/post-processor/manifest"
	shelllocalpostprocessor "github.com/hashicorp/packer/post-processor/shell-local"
	vagrantpostprocessor "github.com/hashicorp/packer/post-processor/vagrant"
	vagrantcloudpostprocessor "github.com/hashicorp/packer/post-processor/vagrant-cloud"
	vspherepostprocessor "github.com/hashicorp/packer/post-processor/vsphere"
	vspheretemplatepostprocessor "github.com/hashicorp/packer/post-processor/vsphere-template"
	ansibleprovisioner "github.com/hashicorp/packer/provisioner/ansible"
	ansiblelocalprovisioner "github.com/hashicorp/packer/provisioner/ansible-local"
	chefclientprovisioner "github.com/hashicorp/packer/provisioner/chef-client"
	chefsoloprovisioner "github.com/hashicorp/packer/provisioner/chef-solo"
	convergeprovisioner "github.com/hashicorp/packer/provisioner/converge"
	fileprovisioner "github.com/hashicorp/packer/provisioner/file"
	powershellprovisioner "github.com/hashicorp/packer/provisioner/powershell"
	puppetmasterlessprovisioner "github.com/hashicorp/packer/provisioner/puppet-masterless"
	puppetserverprovisioner "github.com/hashicorp/packer/provisioner/puppet-server"
	saltmasterlessprovisioner "github.com/hashicorp/packer/provisioner/salt-masterless"
	shellprovisioner "github.com/hashicorp/packer/provisioner/shell"
	shelllocalprovisioner "github.com/hashicorp/packer/provisioner/shell-local"
	windowsrestartprovisioner "github.com/hashicorp/packer/provisioner/windows-restart"
	windowsshellprovisioner "github.com/hashicorp/packer/provisioner/windows-shell"
)

type PluginCommand struct {
	Meta
}

var Builders = map[string]packer.Builder{
	"alicloud-ecs":        new(alicloudecsbuilder.Builder),
	"amazon-chroot":       new(amazonchrootbuilder.Builder),
	"amazon-ebs":          new(amazonebsbuilder.Builder),
	"amazon-ebssurrogate": new(amazonebssurrogatebuilder.Builder),
	"amazon-ebsvolume":    new(amazonebsvolumebuilder.Builder),
	"amazon-instance":     new(amazoninstancebuilder.Builder),
	"azure-arm":           new(azurearmbuilder.Builder),
	"cloudstack":          new(cloudstackbuilder.Builder),
	"digitalocean":        new(digitaloceanbuilder.Builder),
	"docker":              new(dockerbuilder.Builder),
	"file":                new(filebuilder.Builder),
	"googlecompute":       new(googlecomputebuilder.Builder),
	"hyperv-iso":          new(hypervisobuilder.Builder),
	"hyperv-vmcx":         new(hypervvmcxbuilder.Builder),
	"lxc":                 new(lxcbuilder.Builder),
	"lxd":                 new(lxdbuilder.Builder),
	"ncloud":              new(ncloudbuilder.Builder),
	"null":                new(nullbuilder.Builder),
	"oneandone":           new(oneandonebuilder.Builder),
	"openstack":           new(openstackbuilder.Builder),
	"oracle-classic":      new(oracleclassicbuilder.Builder),
	"oracle-oci":          new(oracleocibuilder.Builder),
	"parallels-iso":       new(parallelsisobuilder.Builder),
	"parallels-pvm":       new(parallelspvmbuilder.Builder),
	"profitbricks":        new(profitbricksbuilder.Builder),
	"qemu":                new(qemubuilder.Builder),
	"scaleway":            new(scalewaybuilder.Builder),
	"triton":              new(tritonbuilder.Builder),
	"virtualbox-iso":      new(virtualboxisobuilder.Builder),
	"virtualbox-ovf":      new(virtualboxovfbuilder.Builder),
	"vmware-iso":          new(vmwareisobuilder.Builder),
	"vmware-vmx":          new(vmwarevmxbuilder.Builder),
}

var Provisioners = map[string]packer.Provisioner{
	"ansible":           new(ansibleprovisioner.Provisioner),
	"ansible-local":     new(ansiblelocalprovisioner.Provisioner),
	"chef-client":       new(chefclientprovisioner.Provisioner),
	"chef-solo":         new(chefsoloprovisioner.Provisioner),
	"converge":          new(convergeprovisioner.Provisioner),
	"file":              new(fileprovisioner.Provisioner),
	"powershell":        new(powershellprovisioner.Provisioner),
	"puppet-masterless": new(puppetmasterlessprovisioner.Provisioner),
	"puppet-server":     new(puppetserverprovisioner.Provisioner),
	"salt-masterless":   new(saltmasterlessprovisioner.Provisioner),
	"shell":             new(shellprovisioner.Provisioner),
	"shell-local":       new(shelllocalprovisioner.Provisioner),
	"windows-restart":   new(windowsrestartprovisioner.Provisioner),
	"windows-shell":     new(windowsshellprovisioner.Provisioner),
}

var PostProcessors = map[string]packer.PostProcessor{
	"alicloud-import":      new(alicloudimportpostprocessor.PostProcessor),
	"amazon-import":        new(amazonimportpostprocessor.PostProcessor),
	"artifice":             new(artificepostprocessor.PostProcessor),
	"checksum":             new(checksumpostprocessor.PostProcessor),
	"compress":             new(compresspostprocessor.PostProcessor),
	"docker-import":        new(dockerimportpostprocessor.PostProcessor),
	"docker-push":          new(dockerpushpostprocessor.PostProcessor),
	"docker-save":          new(dockersavepostprocessor.PostProcessor),
	"docker-tag":           new(dockertagpostprocessor.PostProcessor),
	"googlecompute-export": new(googlecomputeexportpostprocessor.PostProcessor),
	"googlecompute-import": new(googlecomputeimportpostprocessor.PostProcessor),
	"manifest":             new(manifestpostprocessor.PostProcessor),
	"shell-local":          new(shelllocalpostprocessor.PostProcessor),
	"vagrant":              new(vagrantpostprocessor.PostProcessor),
	"vagrant-cloud":        new(vagrantcloudpostprocessor.PostProcessor),
	"vsphere":              new(vspherepostprocessor.PostProcessor),
	"vsphere-template":     new(vspheretemplatepostprocessor.PostProcessor),
}

var pluginRegexp = regexp.MustCompile("packer-(builder|post-processor|provisioner)-(.+)")

func (c *PluginCommand) Run(args []string) int {
	// This is an internal call (users should not call this directly) so we're
	// not going to do much input validation. If there's a problem we'll often
	// just crash. Error handling should be added to facilitate debugging.
	log.Printf("args: %#v", args)
	if len(args) != 1 {
		c.Ui.Error("Wrong number of args")
		return 1
	}

	// Plugin will match something like "packer-builder-amazon-ebs"
	parts := pluginRegexp.FindStringSubmatch(args[0])
	if len(parts) != 3 {
		c.Ui.Error(fmt.Sprintf("Error parsing plugin argument [DEBUG]: %#v", parts))
		return 1
	}
	pluginType := parts[1] // capture group 1 (builder|post-processor|provisioner)
	pluginName := parts[2] // capture group 2 (.+)

	server, err := plugin.Server()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error starting plugin server: %s", err))
		return 1
	}

	switch pluginType {
	case "builder":
		builder, found := Builders[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load builder: %s", pluginName))
			return 1
		}
		server.RegisterBuilder(builder)
	case "provisioner":
		provisioner, found := Provisioners[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load provisioner: %s", pluginName))
			return 1
		}
		server.RegisterProvisioner(provisioner)
	case "post-processor":
		postProcessor, found := PostProcessors[pluginName]
		if !found {
			c.Ui.Error(fmt.Sprintf("Could not load post-processor: %s", pluginName))
			return 1
		}
		server.RegisterPostProcessor(postProcessor)
	}

	server.Serve()

	return 0
}

func (*PluginCommand) Help() string {
	helpText := `
Usage: packer plugin PLUGIN

  Runs an internally-compiled version of a plugin from the packer binary.

  NOTE: this is an internal command and you should not call it yourself.
`

	return strings.TrimSpace(helpText)
}

func (c *PluginCommand) Synopsis() string {
	return "internal plugin command"
}
