package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func getKube() string {
	configPaths := filepath.SplitList(os.Getenv("KUBECONFIG"))

	var context string
	for _, configPath := range configPaths {
		context = getKubeCtx(configPath)
		if context != "" {
			break
		}
	}
	if context == "" {
		return ""
	}

	namespaceBytes, err := ioutil.ReadFile(filepath.Join(os.Getenv("HOME"), ".kubectl-namespace"))
	namespace := strings.TrimSpace(string(namespaceBytes))
	if err != nil {
		if !os.IsNotExist(err) {
			handleError(err)
		}
		namespace = ""
	}
	if namespace != "" {
		context = fmt.Sprintf("(%v/%v)", context, namespace)
	}

	return withColor(red, context)
}

func getKubeCtx(configPath string) string {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		//non-existence is acceptable, just make the caller continue with the next configPath
		if !os.IsNotExist(err) {
			handleError(err)
		}
		return ""
	}

	var data struct {
		CurrentContext string `yaml:"current-context"`
	}
	err = yaml.Unmarshal(buf, &data)
	handleError(err)
	return strings.TrimSpace(data.CurrentContext)
}

func getOSCloud() string {
	cloudName := os.Getenv("CURRENT_OS_CLOUD")
	if cloudName == "" {
		return ""
	}
	return withColor(bBlack, cloudName)
}
