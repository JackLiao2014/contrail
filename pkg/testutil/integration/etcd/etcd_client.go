package integrationetcd

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/models"
)

// Integration test settings.
const (
	Endpoint           = "localhost:2379"
	Prefix             = "contrail"
	etcdDialTimeout    = 10 * time.Second
	etcdRequestTimeout = 10 * time.Second
	etcdWatchTimeout   = 10 * time.Second

	AccessControlListSchemaID    = "access_control_list"
	ApplicationPolicySetSchemaID = "application_policy_set"
	NetworkIPAMSchemaID          = "network_ipam"
	ProjectSchemaID              = "project"
	SecurityGroupSchemaID        = "security_group"
	VirtualNetworkSchemaID       = "virtual_network"
)

// EtcdClient is etcd client extending etcd.clientv3 with test functionality and using etcd v3 API.
type EtcdClient struct {
	*clientv3.Client
	log *logrus.Entry
}

// NewEtcdClient is a constructor of etcd client.
// After usage Close() needs to be called to close underlying connections.
func NewEtcdClient(t *testing.T) *EtcdClient {
	l := pkglog.NewLogger("etcd-client")
	l.WithFields(logrus.Fields{"endpoint": Endpoint, "dial-timeout": etcdDialTimeout}).Debug("Connecting")
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{Endpoint},
		DialTimeout: etcdDialTimeout,
	})
	require.NoError(t, err, "connecting etcd failed")

	return &EtcdClient{
		Client: c,
		log:    l,
	}
}

// Close closes connection to etcd.
func (e *EtcdClient) Close(t *testing.T) {
	err := e.Client.Close()
	assert.NoError(t, err, "closing etcd.clientv3.Client failed")
}

// DeleteNetworkIPAM deletes NetworkIPAM resource.
func (e *EtcdClient) DeleteNetworkIPAM(t *testing.T, uuid string, opts ...clientv3.OpOption) {
	e.DeleteKey(t, JSONEtcdKey(NetworkIPAMSchemaID, uuid), opts...)
}

// DeleteProject deletes Project resource.
func (e *EtcdClient) DeleteProject(t *testing.T, uuid string, opts ...clientv3.OpOption) {
	e.DeleteKey(t, JSONEtcdKey(ProjectSchemaID, uuid), opts...)
}

// DeleteSecurityGroup deletes SecurityGroup resource.
func (e *EtcdClient) DeleteSecurityGroup(t *testing.T, uuid string, opts ...clientv3.OpOption) {
	e.DeleteKey(t, JSONEtcdKey(SecurityGroupSchemaID, uuid), opts...)
}

// Clear recursively deletes all keys starting with "etcd.path" prefix.
func (e *EtcdClient) Clear(t *testing.T) (revision int64) {
	return e.DeleteKey(t, "/"+viper.GetString("etcd.path"), clientv3.WithPrefix())
}

// GetKey gets etcd key.
func (e *EtcdClient) GetKey(t *testing.T, key string, opts ...clientv3.OpOption) *clientv3.GetResponse {
	ctx, cancel := context.WithTimeout(context.Background(), etcdRequestTimeout)
	defer cancel()

	r, err := e.Get(ctx, key, opts...)
	assert.NoError(t, err, fmt.Sprintf("getting etcd resource failed\n response: %+v", r))
	e.log.WithFields(logrus.Fields{"prefix": key, "response": r}).Debug("Received etcd Get response")
	return r
}

// DeleteKey deletes etcd key.
func (e *EtcdClient) DeleteKey(t *testing.T, key string, opts ...clientv3.OpOption) (revision int64) {
	ctx, cancel := context.WithTimeout(context.Background(), etcdRequestTimeout)
	defer cancel()

	r, err := e.Delete(ctx, key, opts...)
	assert.NoError(t, err, fmt.Sprintf("deleting etcd resource from etcd failed\n response: %+v", r))
	return r.Header.Revision
}

// WatchKey watches value changes for provided key and returns collect method that collect captured values.
func (e *EtcdClient) WatchKey(
	key string, opts ...clientv3.OpOption,
) (collect func() []string) {
	return e.WatchKeyN(key, 0, 0, opts...)
}

// WatchKeyN watches value changes for provided key n times and returns collect method
// that collects the captured values.
// If there were less than n events then it waits until timeout passes.
func (e *EtcdClient) WatchKeyN(
	key string, n int, timeout time.Duration, opts ...clientv3.OpOption,
) (collect func() []string) {
	resultChan := make(chan []string)
	ctx, cancel := context.WithCancel(context.Background())
	wchan := e.Client.Watch(ctx, key, opts...)

	go func() {
		var result []string

		for val := range wchan {
			for _, ev := range val.Events {
				result = append(result, string(ev.Kv.Value))
				if n > 0 && len(result) >= n {
					cancel()
				}
			}
		}

		resultChan <- result
		close(resultChan)
	}()

	return func() (vals []string) {
		select {
		case <-ctx.Done():
		case <-time.After(timeout):
			cancel()
		}
		return <-resultChan
	}
}

// WatchResource spawns a watch on specified resource.
func (e *EtcdClient) WatchResource(
	schemaID, uuid string, opts ...clientv3.OpOption,
) (clientv3.WatchChan, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), etcdWatchTimeout)
	w := e.Watch(ctx, JSONEtcdKey(schemaID, uuid), opts...)
	return w, ctx, cancel
}

// CheckKeyDoesNotExist checks that there is no value on given key.
func (e *EtcdClient) CheckKeyDoesNotExist(t *testing.T, key string) {
	gr := e.GetKey(t, key, clientv3.WithPrefix())
	assert.Equal(t, int64(0), gr.Count, fmt.Sprintf("key %v should be empty", key))
}

// GetString gets a string value in Etcd
func (e *EtcdClient) GetString(t *testing.T, key string) (value string, revision int64) {
	err := e.Client.Sync(context.Background())
	assert.NoError(t, err)

	kvHandle := clientv3.NewKV(e.Client)
	response, err := kvHandle.Get(context.Background(), key)
	require.NoError(t, err)
	require.NotEmpty(t, response.Kvs)

	return string(response.Kvs[0].Value), response.Header.Revision
}

// ExpectValue gets key and checks if value and revision match.
func (e *EtcdClient) ExpectValue(t *testing.T, key string, value string, revision int64) {
	nextVal, nextRev := e.GetString(t, key)
	assert.Equal(t, value, nextVal)
	assert.Equal(t, revision, nextRev)
}

// JSONEtcdKey returns etcd key of JSON-encoded resource.
func JSONEtcdKey(schemaID, uuid string) string {
	return models.ResourceKey(schemaID, uuid)
}

// RetrieveCreateEvent blocks and retrieves create Event from given watch channel.
func RetrieveCreateEvent(ctx context.Context, t *testing.T, watch clientv3.WatchChan) *clientv3.Event {
	events := RetrieveWatchEvents(ctx, t, watch)
	if assert.Equal(t, 1, len(events)) {
		assert.True(t, events[0].IsCreate())
		return events[0]
	}
	return nil
}

// RetrieveWatchEvents blocks and retrieves events from given watch channel.
func RetrieveWatchEvents(ctx context.Context, t *testing.T, watch clientv3.WatchChan) []*clientv3.Event {
	wr := <-watch
	assert.NoError(t, wr.Err(), "watching etcd key failed")
	if errors.Cause(ctx.Err()) == context.DeadlineExceeded {
		assert.Fail(t, "watching etcd key timed out")
	}

	return wr.Events
}
