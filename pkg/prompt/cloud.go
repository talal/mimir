package prompt

import (
	"fmt"
	"os"

	"github.com/talal/go-bits/color"
)

// getOSCloud returns the current OpenStack cloud info using the
// "CURRENT_OS_CLOUD" environment variable, if set. If not, then individual
// OpenStack environment variables are used to get the cloud info.
func getOSCloud() string {
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
	if cloudInfo == "" {
		return ""
	}

	return color.Sprintf(color.Magenta, cloudInfo)
}

// getOSEnvVal takes two keys for OpenStack environment variables and returns
// the corresponding value for the first key, if no value exists, then the
// corresponding value for the second key is returned.
func getOSEnvVal(key1, key2 string) string {
	var val string
	if val = os.Getenv(key1); val == "" {
		val = os.Getenv(key2)
	}

	return val
}
