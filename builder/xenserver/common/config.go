//go:generate packer-sdc mapstructure-to-hcl2 -type Config,DiskConfig
package common

import (
	"time"

	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	CommonConfig        `mapstructure:",squash"`
	Comm                communicator.Config `mapstructure:",squash"`

	VCPUsMax       uint         `mapstructure:"vcpus_max"`
	VCPUsAtStartup uint         `mapstructure:"vcpus_atstartup"`
	VMMemory       uint         `mapstructure:"vm_memory"`
	Disks          []DiskConfig `mapstructure:"disks"`
	// Deprecated: use Disks instead
	DiskSize      uint              `mapstructure:"disk_size"`
	CloneTemplate string            `mapstructure:"clone_template"`
	VMOtherConfig map[string]string `mapstructure:"vm_other_config"`
	VMTags        []string          `mapstructure:"vm_tags"`

	ISOChecksum string   `mapstructure:"iso_checksum"`
	ISOUrls     []string `mapstructure:"iso_urls"`
	ISOUrl      string   `mapstructure:"iso_url"`
	ISOName     string   `mapstructure:"iso_name"`

	PlatformArgs map[string]string `mapstructure:"platform_args"`

	RawInstallTimeout string        `mapstructure:"install_timeout"`
	InstallTimeout    time.Duration ``
	SourcePath        string        `mapstructure:"source_path"`

	Firmware        string `mapstructure:"firmware"`
	SkipSetTemplate bool   `mapstructure:"skip_set_template"`

	ctx interpolate.Context
}

type DiskConfig struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Size        uint   `mapstructure:"size"`
	ReadOnly    bool   `mapstructure:"read_only"`
}

func (c Config) GetInterpContext() *interpolate.Context {
	return &c.ctx
}
