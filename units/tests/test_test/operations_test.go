package test_test

import (
	"testing"

	"github.com/cf-deployment/units/helpers"
)

const testDirectory = "operations/test"

var testTests = map[string]helpers.OpsFileTestParams{
	"add-datadog-firehose-nozzle.yml": {
		Vars: []string{"datadog_api_key=XYZ", "datadog_metric_prefix=foo.bar", "traffic_controller_external_port=8443"},
	},
	"add-oidc-provider.yml": {},
	"add-persistent-isolation-segment-diego-cell-bosh-lite.yml": {
		Ops: []string{"add-persistent-isolation-segment-diego-cell.yml", "add-persistent-isolation-segment-diego-cell-bosh-lite.yml"},
	},
	"add-persistent-isolation-segment-diego-cell.yml": {},
	"add-persistent-isolation-segment-router.yml":     {},
	"alter-ssh-proxy-redirect-uri.yml":                {},
	"enable-nfs-test-server.yml":                      {},
	"enable-nfs-test-ldapserver.yml": {
		Ops: []string{"../enable-nfs-volume-service.yml", "enable-nfs-test-server.yml", "enable-nfs-test-ldapserver.yml"},
	},
	"enable-smb-test-server.yml": {
		Vars: []string{"smb-password=FOO.PASS", "smb-username=BAR.USER"},
	},
}

func TestTest(t *testing.T) {
	cfDeploymentHome, err := helpers.SetPath()
	if err != nil {
		t.Fatalf("setup: %v", err)
	}

	suite := helpers.NewSuiteTest(cfDeploymentHome, testDirectory, testTests)
	suite.EnsureTestCoverage(t)
	suite.InterpolateTest(t)
}
