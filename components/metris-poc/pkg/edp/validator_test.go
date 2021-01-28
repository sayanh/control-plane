package edp

import (
	"github.com/pkg/errors"
	"testing"
)

const (
	ValidData = `
{
  "resource_groups": [
    "group1",
    "group2"
  ],
  "compute": {
    "vm_types": [
      {
        "name": "Standard_D8_v3",
        "count": 3
      },
      {
        "name": "Standard_D6_v3",
        "count": 2
      }
    ],
    "provisioned_cpus": 24,  
    "provisioned_ram_gb": 96,
    "provisioned_volumes": {
      "size_gb_total": 150,
      "count": 3,
      "size_gb_rounded": 192
    }
  },
  "networking": {
    "provisioned_loadbalancers": 1,
    "provisioned_vnets": 2,
    "provisioned_ips": 3
  },
  "event_hub": {
    "number_namespaces": 3,
    "incoming_requests_pt1m": 3,
    "max_incoming_bytes_pt1m": 5600,
    "max_outgoing_bytes_pt1m": 5600,
    "incoming_requests_pt5m": 0,
    "max_incoming_bytes_pt5m": 0,
    "max_outgoing_bytes_pt5m": 0
  }
}
`
	InvalidDataMissingResourceGroups = `
{
  "compute": {
    "vm_types": [
      {
        "name": "Standard_D8_v3",
        "count": 3
      },
      {
        "name": "Standard_D6_v3",
        "count": 2
      }
    ],
    "provisioned_cpus": 24,  
    "provisioned_ram_gb": 96,
    "provisioned_volumes": {
      "size_gb_total": 150,
      "count": 3,
      "size_gb_rounded": 192
    }
  },
  "networking": {
    "provisioned_loadbalancers": 1,
    "provisioned_vnets": 2,
    "provisioned_ips": 3
  },
  "event_hub": {
    "number_namespaces": 3,
    "incoming_requests_pt1m": 3,
    "max_incoming_bytes_pt1m": 5600,
    "max_outgoing_bytes_pt1m": 5600,
    "incoming_requests_pt5m": 0,
    "max_incoming_bytes_pt5m": 0,
    "max_outgoing_bytes_pt5m": 0
  }
}
`
	InvalidDataMissingCompute = `
{
  "resource_groups": [
    "group1",
    "group2"
  ],
  "networking": {
    "provisioned_loadbalancers": 1,
    "provisioned_vnets": 2,
    "provisioned_ips": 3
  },
  "event_hub": {
    "number_namespaces": 3,
    "incoming_requests_pt1m": 3,
    "max_incoming_bytes_pt1m": 5600,
    "max_outgoing_bytes_pt1m": 5600,
    "incoming_requests_pt5m": 0,
    "max_incoming_bytes_pt5m": 0,
    "max_outgoing_bytes_pt5m": 0
  }
}
`
	InvalidDataMissingNetworking = `
{
  "resource_groups": [
    "group1",
    "group2"
  ],
  "compute": {
    "vm_types": [
      {
        "name": "Standard_D8_v3",
        "count": 3
      },
      {
        "name": "Standard_D6_v3",
        "count": 2
      }
    ],
    "provisioned_cpus": 24,  
    "provisioned_ram_gb": 96,
    "provisioned_volumes": {
      "size_gb_total": 150,
      "count": 3,
      "size_gb_rounded": 192
    }
  },
  "event_hub": {
    "number_namespaces": 3,
    "incoming_requests_pt1m": 3,
    "max_incoming_bytes_pt1m": 5600,
    "max_outgoing_bytes_pt1m": 5600,
    "incoming_requests_pt5m": 0,
    "max_incoming_bytes_pt5m": 0,
    "max_outgoing_bytes_pt5m": 0
  }
}
`
	InvalidDataMissingEventHub = `
{
  "resource_groups": [
    "group1",
    "group2"
  ],
  "compute": {
    "vm_types": [
      {
        "name": "Standard_D8_v3",
        "count": 3
      },
      {
        "name": "Standard_D6_v3",
        "count": 2
      }
    ],
    "provisioned_cpus": 24,  
    "provisioned_ram_gb": 96,
    "provisioned_volumes": {
      "size_gb_total": 150,
      "count": 3,
      "size_gb_rounded": 192
    }
  },
  "networking": {
    "provisioned_loadbalancers": 1,
    "provisioned_vnets": 2,
    "provisioned_ips": 3
  }
}
`
	MissingResourceGroupsErrorMsg = "Key: 'ConsumptionMetrics.ResourceGroups' Error:Field validation for 'ResourceGroups' failed on the 'required' tag"
	MissingComputeErrorMsg        = "Key: 'ConsumptionMetrics.Compute.VMTypes' Error:Field validation for 'VMTypes' failed on the 'required' tag"
)

func TestValidateConsumptionMetrics(t *testing.T) {
	testCases := []struct {
		name               string
		data               string
		expectedValidation bool
		expectedError      error
	}{
		{
			name:               "valid input data",
			data:               ValidData,
			expectedValidation: true,
			expectedError:      nil,
		},
		{
			name:               "invalid data input, missing resource groups",
			data:               InvalidDataMissingResourceGroups,
			expectedValidation: false,
			expectedError:      errors.New(MissingResourceGroupsErrorMsg),
		},
		{
			name:               "invalid data input, missing compute",
			data:               InvalidDataMissingCompute,
			expectedValidation: false,
			expectedError:      errors.New(MissingComputeErrorMsg),
		},
		{
			name:               "invalid data input, missing networking",
			data:               InvalidDataMissingNetworking,
			expectedValidation: true,
			expectedError:      nil,
		},
		{
			name:               "invalid data input, missing eventhub",
			data:               InvalidDataMissingEventHub,
			expectedValidation: true,
			expectedError:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			valid, err := ValidateConsumptionMetrics(tc.data)
			if err != nil && tc.expectedError == nil {
				t.Errorf("error should be: %v, got: %v", tc.expectedError, err)
				return
			}

			if tc.expectedError != nil && err == nil {
				t.Errorf("error should be: %v, got: %v", tc.expectedError, err)
				return
			}

			if tc.expectedError != nil && tc.expectedError.Error() != err.Error() {
				t.Errorf("error should be:\n%v\n, got:\n%v\n", tc.expectedError, err)
				return
			}

			if valid != tc.expectedValidation {
				t.Errorf("received wrong consumption metrics, expected: %v but got: %v",
					tc.expectedValidation, valid)
			}
		})
	}
}
