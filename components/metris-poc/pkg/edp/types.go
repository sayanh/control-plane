package edp

type ConsumptionMetrics struct {
	ResourceGroups []string `json:"resource_groups" validate:"required"`
	Compute        struct {
		VMTypes []struct {
			Name  string `json:"name" validate:"required"`
			Count int    `json:"count" validate:"numeric"`
		} `json:"vm_types" validate:"required"`
		ProvisionedCpus    int `json:"provisioned_cpus" validate:"numeric"`
		ProvisionedRAMGb   int `json:"provisioned_ram_gb" validate:"numeric"`
		ProvisionedVolumes struct {
			SizeGbTotal   int `json:"size_gb_total" validate:"numeric"`
			Count         int `json:"count" validate:"numeric"`
			SizeGbRounded int `json:"size_gb_rounded" validate:"numeric"`
		} `json:"provisioned_volumes" validate:"required"`
	} `json:"compute" validate:"required"`
	Networking struct {
		ProvisionedLoadbalancers int `json:"provisioned_loadbalancers" validate:"numeric"`
		ProvisionedVnets         int `json:"provisioned_vnets" validate:"numeric"`
		ProvisionedIps           int `json:"provisioned_ips" validate:"numeric"`
	} `json:"networking" validate:"required"`
	EventHub struct {
		NumberNamespaces     int `json:"number_namespaces" validate:"numeric"`
		IncomingRequestsPt1M int `json:"incoming_requests_pt1m" validate:"numeric"`
		MaxIncomingBytesPt1M int `json:"max_incoming_bytes_pt1m" validate:"numeric"`
		MaxOutgoingBytesPt1M int `json:"max_outgoing_bytes_pt1m" validate:"numeric"`
		IncomingRequestsPt5M int `json:"incoming_requests_pt5m" validate:"numeric"`
		MaxIncomingBytesPt5M int `json:"max_incoming_bytes_pt5m" validate:"numeric"`
		MaxOutgoingBytesPt5M int `json:"max_outgoing_bytes_pt5m" validate:"numeric"`
	} `json:"event_hub" validate:"required"`
}
