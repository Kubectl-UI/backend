package config

import (
	"github.com/spf13/viper"
)

type KubectlUI struct {
	v *viper.Viper
}

const (
	ApplicationPort = "kubectlui.application_port"
	ConfigCommand   = "kubectlui.config_cmd"
	FilePath        = "kubectlui.config_file_path"
)

func (k *KubectlUI) setDefaults() {
	k.v.SetDefault(ApplicationPort, "8080")
	k.v.SetDefault(ConfigCommand, "kubectl")
	k.v.SetDefault(FilePath, "./examples")
}

func Load() *KubectlUI {
	k := &KubectlUI{v: viper.New()}

	k.setDefaults()
	return k
}

func (k *KubectlUI) ReplaceDefault(key string, value interface{}) {
	if value != "" && value != nil {
		k.v.Set(key, value)
	}
}

func (k *KubectlUI) Get(key string) interface{} {
	return k.v.Get(key)
}
