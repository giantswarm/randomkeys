package randomkeystest

import (
	"context"

	"github.com/giantswarm/randomkeys/v3"
)

type Searcher struct {
}

func NewSearcher() *Searcher {
	return &Searcher{}
}

func (s *Searcher) SearchCluster(ctx context.Context, clusterID string) (randomkeys.Cluster, error) {
	return randomkeys.Cluster{}, nil
}
