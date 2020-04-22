package kv

import (
	"context"
	"encoding/binary"
	"fmt"
	"net/url"

	"github.com/phemmer/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/phemmer/go-pulsar-state/bookkeeper/proto/kv/rpc"
	"github.com/phemmer/go-pulsar-state/bookkeeper/proto/storage"
	"github.com/phemmer/go-pulsar-state/bookkeeper/proto/stream"
)

// https://github.com/apache/bookkeeper/blob/release-4.10.0/stream/clients/python/bookkeeper/common/constants.py#L18
var rootRangeID = 0
var rootRangeMetadata = metadata.New(map[string]string{
	"bk-rt-sc-id-bin": int64bin(int64(rootRangeID)),
})

func int64bin(v int64) string {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], uint64(v))
	return string(buf[:])
}

type respHdr interface {
	GetHeader() *rpc.ResponseHeader
}

type StorageError storage.StatusCode

func (se StorageError) Code() storage.StatusCode {
	return storage.StatusCode(se)
}
func (se StorageError) Error() string {
	return se.Code().String()
}

func errChk(resp respHdr, err error) error {
	if err != nil {
		return err
	}
	if resp == nil {
		return fmt.Errorf("empty response")
	}
	hdr := resp.GetHeader()
	if hdr.Code == storage.StatusCode_SUCCESS {
		return nil
	}
	return StorageError(hdr.Code)
}

type StorageClientSettings struct {
	ServiceURI url.URL
}

type Client struct {
	scs       StorageClientSettings
	conn      *grpc.ClientConn
	namespace string
}

func NewClient(ctx context.Context, addr, namespace string) (*Client, error) {
	//	if scs.ServiceURI.Scheme != "bk" {
	//		return nil, fmt.Errorf("only 'bk' service URIs supported")
	//	}

	//	conn, err := grpc.Dial(scs.ServiceURI.Host + ":" + scs.ServiceURI.Port())
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return nil, errors.F(err, "connecting")
	}

	cl := &Client{
		conn:      conn,
		namespace: namespace,
	}
	return cl, nil
}

func (cl *Client) GetTable(ctx context.Context, name string) (*Table, error) {
	//	delreq := &storage.DeleteStreamRequest{
	//		NsName: cl.namespace,
	//		Name: name,
	//	}

	req := &storage.GetStreamRequest{
		Id: &storage.GetStreamRequest_StreamName{&stream.StreamName{
			NamespaceName: cl.namespace,
			StreamName:    name,
		}},
	}

	rootCl := storage.NewRootRangeServiceClient(cl.conn)
	//	rootCl.DeleteStream(
	//		metadata.NewOutgoingContext(ctx, rootRangeMetadata),
	//		delreq)
	resp, err := rootCl.GetStream(
		metadata.NewOutgoingContext(ctx, rootRangeMetadata),
		req)
	if err != nil {
		return nil, errors.F(err, "GetStream")
	}
	if resp.Code != storage.StatusCode_SUCCESS {
		return nil, errors.F(StorageError(resp.Code), "retrieving stream")
	}

	return newTable(cl.conn, resp.StreamProps), nil
}

func (cl *Client) CreateTable(ctx context.Context, name string, conf *stream.StreamConfiguration) (*Table, error) {
	req := &storage.CreateStreamRequest{
		NsName:     cl.namespace,
		Name:       name,
		StreamConf: conf,
	}

	rootCl := storage.NewRootRangeServiceClient(cl.conn)
	resp, err := rootCl.CreateStream(
		metadata.NewOutgoingContext(ctx, rootRangeMetadata),
		req)
	if err != nil {
		return nil, errors.F(err, "CreateStream")
	}
	if resp.Code != storage.StatusCode_SUCCESS {
		return nil, errors.F(StorageError(resp.Code), "creating stream")
	}

	return newTable(cl.conn, resp.StreamProps), nil
}
