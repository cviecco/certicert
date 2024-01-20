package server

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	acmecfg "github.com/Cloud-Foundations/golib/pkg/crypto/certmanager/config"
	//"github.com/Cloud-Foundations/golib/pkg/log"
)

type baseConfig struct {
	HttpAddress       string `yaml:"http_address"`
	TLSCertFilename   string `yaml:"tls_cert_filename"`
	TLSKeyFilename    string `yaml:"tls_key_filename"`
	ACME              acmecfg.AcmeConfig
	CAFKeyilename     string `yaml:"ssh_ca_filename"`
	Ed25519CAFilename string `yaml:"ed25519_ca_keyfilename"`
	//AutoUnseal                  autoUnseal `yaml:"auto_unseal"`
	HtpasswdFilename string `yaml:"htpasswd_filename"`
	ExternalAuthCmd  string `yaml:"external_auth_command"`
	ClientCAFilename string `yaml:"client_ca_filename"`
	//KeymasterPublicKeysFilename string `yaml:"keymaster_public_keys_filename"`
	//HostIdentity                 string        `yaml:"host_identity"`
	//DataDirectory                string        `yaml:"data_directory"`
	//SharedDataDirectory          string        `yaml:"shared_data_directory"`
	AdminUsers  []string `yaml:"admin_users"`
	AdminGroups []string `yaml:"admin_groups"`
	//PublicLogs                   bool          `yaml:"public_logs"`
}

type AppConfigFile struct {
	Base baseConfig
}

func LoadVerifyConfigFile(configFilename string) (*AppConfigFile, error) {
	return loadVerifyConfigFile(configFilename)
}

func loadVerifyConfigFile(configFilename string) (*AppConfigFile, error) {
	var config AppConfigFile
	if _, err := os.Stat(configFilename); os.IsNotExist(err) {
		return nil, fmt.Errorf("No config file: please re-run with -configHost")
	}
	source, err := os.ReadFile(configFilename)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file err=%s", err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse config filei %s", err)

	}
	// TODO: actually check things out
	return &config, nil
}
