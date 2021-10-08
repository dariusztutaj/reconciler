package invoker

import (
	"github.com/kyma-incubator/reconciler/pkg/cluster"
	"github.com/kyma-incubator/reconciler/pkg/model"
	"github.com/kyma-incubator/reconciler/pkg/scheduler/reconciliation"
	"github.com/stretchr/testify/require"
	"testing"
)

var clusterStateMock = &cluster.State{
	Cluster: &model.ClusterEntity{
		Version:    1,
		Cluster:    "testCluster",
		Contract:   1,
		Kubeconfig: "abc...",
	},
	Configuration: &model.ClusterConfigurationEntity{
		Version:        1,
		Cluster:        "testCluster",
		ClusterVersion: 1,
		KymaVersion:    "1.2.3",
	},
	Status: &model.ClusterStatusEntity{
		ID:             1,
		Cluster:        "testCluster",
		ClusterVersion: 1,
		ConfigVersion:  1,
		Status:         model.ClusterStatusReconcilePending,
	},
}

func requireOperationState(t *testing.T, reconRepo reconciliation.Repository, opEntity *model.OperationEntity, state model.OperationState) {
	opUpdated, err := reconRepo.GetOperation(opEntity.SchedulingID, opEntity.CorrelationID)
	require.NoError(t, err)
	require.Equal(t, state, opUpdated.State)
	if opUpdated.State == model.OperationStateFailed ||
		opUpdated.State == model.OperationStateError ||
		opUpdated.State == model.OperationStateClientError {
		require.NotEmpty(t, opUpdated.Reason)
	} else {
		require.Empty(t, opUpdated.Reason)
	}
}