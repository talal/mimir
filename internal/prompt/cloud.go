package prompt

import (
	"fmt"
	"os"

	"github.com/talal/go-bits/color"
)

// getOSCloud returns the current OpenStack cloud info using the standard
// OpenStack environment variables.
//
// If the 'CURRENT_OS_CLOUD' environment variable is defined then its value is
// returned instead.
func getOSCloud() string {
	if getBoolEnv("MIMIR_DISABLE_CLOUD") {
		return ""
	}

	if cloud := os.Getenv("CURRENT_OS_CLOUD"); cloud != "" {
		return color.Sprintf(color.Magenta, cloud)
	}

	cloudInfo := fmt.Sprintf("%s/%s@%s/%s@%s",
		getOSEnvVal("OS_REGION_NAME", ""),
		getOSEnvVal("OS_USERNAME", ""),
		getOSEnvVal("OS_USER_DOMAIN_NAME", "OS_USER_DOMAIN_ID"),
		getOSEnvVal("OS_PROJECT_NAME", "OS_PROJECT_ID"),
		getOSEnvVal("OS_PROJECT_DOMAIN_NAME", "OS_PROJECT_DOMAIN_ID"))

	if cloudInfo == "/@/@" {
		return ""
	}

	return color.Sprintf(color.Magenta, cloudInfo)
}

// getOSEnvVal takes two keys for OpenStack environment variables and returns
// the corresponding value for the first key, if no value exists, then the
// corresponding value for the second key is returned.
func getOSEnvVal(key1, key2 string) (val string) {
	val = os.Getenv(key1)
	if val == "" {
		val = os.Getenv(key2)
	}

	return
}
