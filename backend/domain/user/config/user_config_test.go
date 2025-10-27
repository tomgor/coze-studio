package config

import (
	"testing"
	// ... 导入其他依赖，如 request, response model
)


func Test_loadConfig(t *testing.T) {

	err :=InitConfig("../../../conf/user/config.yaml")
	if err != nil {
		t.Errorf("加载用户配置失败: %v", err)
		return
	}

	t.Logf("记载用户配置成功: %v", GlobalConfig)

}
