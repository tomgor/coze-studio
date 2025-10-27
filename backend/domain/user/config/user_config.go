package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	SimpleSSO    []SimpleSSOConfig `yaml:"simple_sso"`
}


type SimpleSSOConfig struct {
    Platform     string `json:"platform" mapstructure:"platform" yaml:"platform"`
    PublicKey    string `json:"public_key" mapstructure:"public_key" yaml:"public_key"`
    PrivateKey   string `json:"private_key" mapstructure:"private_key" yaml:"private_key"`
    FixedOrgID   string `json:"fixed_org_id" mapstructure:"fixed_org_id" yaml:"fixed_org_id"`
    FixedRoleId  string `json:"fixed_role_id" mapstructure:"fixed_role_id" yaml:"fixed_role_id"`
    InitPassword string `json:"init_password" mapstructure:"init_password" yaml:"init_password"`
}


// GlobalConfig 是一个全局变量，用于存储加载后的配置。
// 注意：如果您的应用是严格面向接口和依赖注入的，最好通过构造函数传递配置，而不是使用全局变量。
var GlobalConfig *Config

// InitConfig 加载并解析 config.yaml 文件。
// 这里的 filePath 应该是相对于运行时的路径。
func InitConfig(filePath string) error {
	// 1. 读取文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败 (%s): %w", filePath, err)
	}

	fmt.Printf("文件内容：%v \n", data)

	// 2. 解析 YAML 到结构体
	if err := yaml.Unmarshal(data, &GlobalConfig); err != nil {
		return fmt.Errorf("解析 YAML 失败: %w", err)
	}


	// 可以在这里添加一些配置校验逻辑
	fmt.Printf("详细信息 (使用 %%v): %+v\n", GlobalConfig)

	return nil
}

// GetConfig 暴露给模块内的其他代码获取配置
func GetConfig() *Config {
	return GlobalConfig
}
