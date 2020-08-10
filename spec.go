package randomkeys

import "context"

type Interface interface {
	SearchCluster(ctx context.Context, clusterID string) (Cluster, error)
}
