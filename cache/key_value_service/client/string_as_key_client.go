package client

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pineal-niwan/sensor/cache/key_value_service/pb"
	"google.golang.org/grpc"
	"sync"
	"time"
)

var (
	ErrClientNotReady = errors.New(`client not ready`)
)

type StringAsKeyClient struct {
	// grpc 连接
	conn *grpc.ClientConn
	// k-v client接口
	client pb.StringKeyValueServiceClient
	// rpc url
	url string
	// re-connection max back off (ms)
	maxBackOff time.Duration
	sync.RWMutex
}

//新建Client -- maxBackOff(ms)
func NewStringAsKeyClient(url string, maxBackOff int64) *StringAsKeyClient {
	return &StringAsKeyClient{
		url:        url,
		maxBackOff: time.Millisecond * time.Duration(maxBackOff),
	}
}

//关闭
func (c *StringAsKeyClient) Close() error {
	c.Lock()
	defer c.Unlock()

	if c.conn != nil {
		err := c.conn.Close()
		if err != nil {
			return err
		}
		//已经关闭了
		c.conn = nil
		c.client = nil
	}
	return nil
}

//连接
func (c *StringAsKeyClient) Dial(timeout time.Duration) error {
	c.Lock()
	defer c.Unlock()

	//拨号连接
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	conn, err := grpc.DialContext(ctx, c.url, grpc.WithInsecure(), grpc.WithBackoffMaxDelay(c.maxBackOff))
	defer cancel()

	if err != nil {
		return err
	}

	//设置client和conn
	c.client = pb.NewStringKeyValueServiceClient(conn)
	c.conn = conn
	return nil
}

// 设置缓存
func (c *StringAsKeyClient) Set(ctx context.Context, key string, value []byte) error {
	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return ErrClientNotReady
	}
	_, err := c.client.Set(ctx, &pb.MsgStringKeyValue{
		Key:   key,
		Value: value,
	})
	return err
}

// 设置缓存-带超时
func (c *StringAsKeyClient) SetWithTimeout(ctx context.Context, key string, value []byte, timeout int64) error {
	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return ErrClientNotReady
	}
	_, err := c.client.SetWithTimeout(ctx, &pb.MsgStringKeyValueTimeout{
		Timeout: timeout,
		Key:     key,
		Value:   value,
	})
	return err
}

// 设置缓存
func (c *StringAsKeyClient) SetIfKeyNotExist(ctx context.Context, key string, value []byte) error {
	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return ErrClientNotReady
	}
	_, err := c.client.SetIfKeyNotExist(ctx, &pb.MsgStringKeyValue{
		Key:   key,
		Value: value,
	})
	return err
}

// 设置缓存-带超时
func (c *StringAsKeyClient) SetWithTimeoutIfKeyNotExist(ctx context.Context, key string, value []byte, timeout int64) error {
	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return ErrClientNotReady
	}
	_, err := c.client.SetWithTimeoutIfKeyNotExist(ctx, &pb.MsgStringKeyValueTimeout{
		Timeout: timeout,
		Key:     key,
		Value:   value,
	})
	return err
}

// 设置缓存
func (c *StringAsKeyClient) SetIfExist(ctx context.Context, key string, value []byte) error {
	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return ErrClientNotReady
	}
	_, err := c.client.SetIfKeyExist(ctx, &pb.MsgStringKeyValue{
		Key:   key,
		Value: value,
	})
	return err
}

// 设置缓存-带超时
func (c *StringAsKeyClient) SetWithTimeoutIfExist(ctx context.Context, key string, value []byte, timeout int64) error {
	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return ErrClientNotReady
	}
	_, err := c.client.SetWithTimeoutIfKeyExist(ctx, &pb.MsgStringKeyValueTimeout{
		Timeout: timeout,
		Key:     key,
		Value:   value,
	})
	return err
}

// 获取缓存
func (c *StringAsKeyClient) Get(ctx context.Context, key string) ([]byte, error) {
	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return nil, ErrClientNotReady
	}
	msgValue, err := c.client.Get(ctx, &pb.MsgStringKey{
		Key: key,
	})

	if err != nil {
		return nil, err
	}

	if msgValue.Ok {
		return msgValue.Value, nil
	} else {
		return nil, nil
	}
}

// 获取缓存后刷新
func (c *StringAsKeyClient) GetThenRefresh(ctx context.Context, key string, timeout int64) ([]byte, error) {
	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return nil, ErrClientNotReady
	}
	msgValue, err := c.client.GetThenRefresh(ctx, &pb.MsgStringKeyTimeout{
		Timeout: timeout,
		Key:     key,
	})

	if err != nil {
		return nil, err
	}

	if msgValue.Ok {
		return msgValue.Value, nil
	} else {
		return nil, nil
	}
}

// 获取缓存长度
func (c *StringAsKeyClient) GetLen(ctx context.Context) (hashLen int, listLen int, err error) {
	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return 0, 0, ErrClientNotReady
	}
	msgLen, err := c.client.GetLen(ctx, &empty.Empty{})

	if err != nil {
		return 0, 0, err
	}
	return int(msgLen.HashLen), int(msgLen.ListLen), nil
}

// 清空缓存
func (c *StringAsKeyClient) Clear(ctx context.Context) error {
	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return ErrClientNotReady
	}
	_, err := c.client.Clear(ctx, &empty.Empty{})
	return err
}
