package modules

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

func GetExeFile() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("%v %v", exePath, err)
		return "", err
	}
	// 返回上级目录
	yamlFile := filepath.Dir(filepath.Dir(filepath.Dir(exePath)))
	// 这个路径是手动指定的，经常需要修改
	configFile := fmt.Sprintf("D:\\go_vscode\\src\\github.com\\Unit-testing\\config\\application.yaml", yamlFile)
	return configFile, nil
}

// 通过Viper获取配置文件
func GetViperConfig(configFileName string) (*viper.Viper, error) {
	configFile, _ := GetExeFile()
	viper.SetConfigFile(configFile)
	content, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Println("ioutil err:%v", err)
		return nil, err
	}
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		fmt.Println("viper.ReadConfig: %v", err)
		return nil, err
	}
	config := viper.Sub(configFileName)
	return config, nil
}
