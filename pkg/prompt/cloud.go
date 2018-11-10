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

// GetOSCloud returns the current OpenStack cloud info using the "CURRENT_OS_CLOUD"
// enviroment variable, if set. If not, then individual OpenStack environment variables
// are used to get the cloud info.
func GetOSCloud() string {
	cloudInfo := os.Getenv("CURRENT_OS_CLOUD")
	if cloudInfo == "" {
		osRegion := getOSEnvVal("OS_REGION_NAME", "")
		osUser := getOSEnvVal("OS_USERNAME", "")
		osUserDomain := getOSEnvVal("OS_USER_DOMAIN_NAME", "OS_USER_DOMAIN_ID")
		osProject := getOSEnvVal("OS_PROJECT_NAME", "OS_PROJECT_ID")
		osProjectDomain := getOSEnvVal("OS_PROJECT_DOMAIN_NAME", "OS_PROJECT_DOMAIN_ID")

		// at least one value should be available
		if osRegion != "" || osUser != "" || osUserDomain != "" ||
			osProject != "" || osProjectDomain != "" {
			cloudInfo = fmt.Sprintf("%s/%s@%s/%s@%s", osRegion, osUser, osUserDomain,
				osProject, osProjectDomain)
		}
	}

	return withColor(bBlack, cloudInfo)
}

// getOSEnvVal takes two keys for OpenStack environment variables and returns the
// corresponding value for the first key, if no value exists, then the corresponding
// value for the second key is returned.
func getOSEnvVal(key1, key2 string) string {
	var val string
	if val = os.Getenv(key1); val == "" {
		val = os.Getenv(key2)
	}

	return val
}
