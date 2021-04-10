package collector

import "github.com/willfantom/mininet-exporter/mininet"

const (
	namespace = "mininet"
)

func getNamespace(client *mininet.Client) string {
	return namespace + "_" + client.Name
}
