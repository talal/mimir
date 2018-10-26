package prompt

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// GetKube iterates over the config files defined in the "KUBECONFIG" environment
// variable and returns the current kubernetes context and namespace.
func GetKube() string {
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

	return withColor(red, context)
}

func getKubeCtx(configPath string) string {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		//non-existence is acceptable, just make the caller continue
		// with the next configPath
		if !os.IsNotExist(err) {
			handleError(err)
		}
		return ""
	}

	var data struct {
		Contexts []struct {
			Context struct {
				Cluster   string `yaml:"cluster"`
				Namespace string `yaml:"namespace"`
				User      string `yaml:"user"`
			} `yaml:"context"`
			Name string `yaml:"name"`
		} `yaml:"contexts"`
		CurrentContext string `yaml:"current-context"`
	}
	err = yaml.Unmarshal(buf, &data)
	handleError(err)

	if data.CurrentContext == "" {
		return ""
	}

	for _, c := range data.Contexts {
		if c.Name == data.CurrentContext {
			return fmt.Sprintf("(%v/%v)", strings.TrimSpace(data.CurrentContext),
				strings.TrimSpace(c.Context.Namespace))
		}
	}

	return strings.TrimSpace(data.CurrentContext)
}

// GetOSCloud uses the "CURRENT_OS_CLOUD" enviroment variable to return
// the current OpenStack cloud.
func GetOSCloud() string {
	cloudName := os.Getenv("CURRENT_OS_CLOUD")
	if cloudName == "" {
		return ""
	}

	return withColor(bBlack, cloudName)
}
