# v1.1.0 (2018-11-09)

New features:
* Kubernetes context and OpenStack cloud info can be turned off by specifying the respective flags to 'false': `MIMIR_KUBE` and `MIMIR_OS_CLOUD`

Changes:
* If `CURRENT_OS_CLOUD` env variable is not available then the OpenStack cloud info is shown using the standard OpenStack environment variables. The scope of the info depends on the environment variables that are available.


# v1.0.0 (2018-11-02)

Initial release.
