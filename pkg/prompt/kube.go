package prompt

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/talal/go-bits/color"
	yaml "gopkg.in/yaml.v2"
)

// getKube iterates over the config files defined in the "KUBECONFIG"
// environment variable and returns the current kubernetes context and
// namespace (if defined).
//
// If multiple files have 'CurrentContext' value then the first one is
// returned.
//
// If the 'CURRENT_KUBE_CTX' environment variable is defined then its value is
// returned instead.
func getKube() string {
	if getBoolEnv("MIMIR_DISABLE_KUBE") {
		return ""
	}

	if ctx := os.Getenv("CURRENT_KUBE_CTX"); ctx != "" {
		return color.Sprintf(color.Yellow, ctx)
	}

	configPaths := filepath.SplitList(os.Getenv("KUBECONFIG"))
	if len(configPaths) == 0 {
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

	for _, path := range configPaths {
		parseKubeConfig(path, &data)
		if data.CurrentContext != "" {
			break
		}
	}
	if data.CurrentContext == "" {
		return ""
	}

	for _, c := range data.Contexts {
		if (c.Name == data.CurrentContext) && (c.Context.Namespace != "") {
			return color.Sprintf(color.Yellow, "(%v/%v)",
				strings.TrimSpace(data.CurrentContext), strings.TrimSpace(c.Context.Namespace))
		}
	}

	return color.Sprintf(color.Yellow, "(%v)", data.CurrentContext)
}

func parseKubeConfig(path string, data interface{}) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		// only handle errors that are not related to file's existence
		if !os.IsNotExist(err) {
			handleError(err)
		}
		return
	}

	err = yaml.Unmarshal(buf, data)
	handleError(err)
}
