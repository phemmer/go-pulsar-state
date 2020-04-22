package pulsar

import (
	"context"
	"net/url"
	"strings"
	"sync"

	"github.com/apache/pulsar/pulsar-function-go/pf"
	"github.com/phemmer/errors"

	"github.com/phemmer/go-pulsar-state/bookkeeper/kv"
	"github.com/phemmer/go-pulsar-state/bookkeeper/proto/storage"
	"github.com/phemmer/go-pulsar-state/bookkeeper/proto/stream"
)

// Start is a convenience wrapper around pf.Start to avoid the need for importing multiple libraries.
func Start(f interface{}) {
	pf.Start(f)
}

type FunctionState struct {
	*pf.FunctionContext

	client *kv.Client
	table  *kv.Table
}

type FunctionStateContext struct {
	*FunctionState
	context.Context
}

var fcfsMapMutex sync.RWMutex
var FCFSMap = map[*pf.FunctionContext]*FunctionState{}

// The state storage URL isn't passed to go programs like it is for python and java.
var StateStorageServiceURL = "bk://127.0.0.1:4181"

func FromContext(ctx context.Context) (*FunctionState, bool) {
	// The bool on return is superfluous, but for compatability we keep it.

	fc, _ := pf.FromContext(ctx)
	if fc == nil {
		return nil, false
	}

	fcfsMapMutex.RLock()
	fs := FCFSMap[fc]
	fcfsMapMutex.RUnlock()
	if fs == nil {
		fs = &FunctionState{
			FunctionContext: fc,
		}
		fcfsMapMutex.Lock()
		FCFSMap[fc] = fs
		fcfsMapMutex.Unlock()
	}
	return fs, true
}

func (fs *FunctionState) getClient(ctx context.Context) (*kv.Client, error) {
	if fs.client != nil {
		return fs.client, nil
	}

	u, err := url.Parse(StateStorageServiceURL)
	if err != nil {
		return nil, errors.F(err, "parsing state storage url")
	}

	kvNS := fs.GetFuncTenant() + "_" + fs.GetFuncNamespace()
	strings.ReplaceAll(kvNS, "-", "_")

	fs.client, err = kv.NewClient(ctx, u.Host, kvNS)
	return fs.client, err
}
func (fs *FunctionState) getTable(ctx context.Context) (*kv.Table, error) {
	if fs.table != nil {
		return fs.table, nil
	}

	cl, err := fs.getClient(ctx)
	if err != nil {
		return nil, err
	}

	fs.table, err = cl.GetTable(ctx, fs.GetFuncName())
	if err != nil && errors.Contains(err, kv.StorageError(storage.StatusCode_STREAM_NOT_FOUND)) {
		// https://github.com/apache/pulsar/blob/v2.5.0/pulsar-functions/instance/src/main/java/org/apache/pulsar/functions/instance/JavaInstanceRunnable.java#L337-L340
		fs.table, err = cl.CreateTable(ctx, fs.GetFuncName(), &stream.StreamConfiguration{
			MinNumRanges:     4,
			InitialNumRanges: 4,
			StorageType:      stream.StorageType_TABLE,
			KeyType:          stream.RangeKeyType_HASH,
			RetentionPolicy: &stream.RetentionPolicy{
				TimePolicy: &stream.TimeBasedRetentionPolicy{
					RetentionMinutes: -1,
				},
			},
		})
		err = errors.F(err, "creating table")
	}
	return fs.table, err
}

func (fs *FunctionState) IncrCounter(ctx context.Context, key string, amount int64) (int64, error) {
	tbl, err := fs.getTable(ctx)
	if err != nil {
		return 0, err
	}

	return tbl.Incr(ctx, key, amount)
}

func (fs *FunctionState) GetCounter(ctx context.Context, key string) (int64, error) {
	tbl, err := fs.getTable(ctx)
	if err != nil {
		return 0, err
	}

	return tbl.GetInt(ctx, key)
}

func (fs *FunctionState) DelCounter(ctx context.Context, key string) error {
	tbl, err := fs.getTable(ctx)
	if err != nil {
		return err
	}

	return tbl.Delete(ctx, key)
}

func (fs *FunctionState) PutState(ctx context.Context, key string, value []byte) error {
	tbl, err := fs.getTable(ctx)
	if err != nil {
		return err
	}

	return tbl.Put(ctx, key, value)
}

func (fs *FunctionState) GetState(ctx context.Context, key string) ([]byte, error) {
	tbl, err := fs.getTable(ctx)
	if err != nil {
		return nil, err
	}

	return tbl.Get(ctx, key)
}
