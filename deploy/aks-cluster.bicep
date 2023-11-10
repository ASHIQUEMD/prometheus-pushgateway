resource aks 'Microsoft.ContainerService/managedClusters@2023-03-01' = {
  name: 'test-cluster'
  location: 'eastus'
  properties: {
    dnsPrefix: 'test-cluster-dns'
    agentPoolProfiles: [
      {
        name: 'agentpool'
        count: 3
        vmSize: 'standard_d2s_v5'
        osType: 'Linux'
        mode: 'System'
      }
    ]
    azureMonitorProfile: {
      metrics: {
        enabled: true
        kubeStateMetrics: {
          metricLabelsAllowlist: '*'
        }
      }
    }
  }
}
