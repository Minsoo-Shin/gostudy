package server

import (
	"context"
	api "github.com/Minsoo-Shin/proglog/api/v1"
	"github.com/Minsoo-Shin/proglog/internal/log"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	for scenario, fn := range map[string]func(
		t *testing.T,
		client api.LogClient,
		config *Config,
	){
		"produce/consume a message to/from the log succeeds": testProduceConsume,
		"produce/consume stream succeds":                     testProduceConsumeStream,
		"consume past log boundary fails":                    testConsumePastBoundary,
	} {
		t.Run(scenario, func(t *testing.T) {
			client, config, teardown := setupTest(t, nil)
			defer teardown()
			fn(t, client, config)
		})
	}
}

func setupTest(t *testing.T, fn func(*Config)) (
	client api.LogClient,
	cfg *Config,
	teardown func(),
) {
	t.Helper()

	l, err := net.Listen("tcp", ":0")
	assert.NoError(t, err)

	clientOptions := []grpc.DialOption{grpc.WithInsecure()}
	cc, err := grpc.Dial(l.Addr().String(), clientOptions...)
	assert.NoError(t, err)

	dir, err := ioutil.TempDir("", "server-test")
	assert.NoError(t, err)

	clog, err := log.NewLog(dir, log.Config{})
	assert.NoError(t, err)

	cfg = &Config{
		CommitLog: clog,
	}
	if fn != nil {
		fn(cfg)
	}
	server, err := NewGRPCServer(cfg)
	assert.NoError(t, err)

	go func() {
		server.Serve(l)
	}()

	client = api.NewLogClient(cc)
	return client, cfg, func() {
		server.Stop()
		cc.Close()
		clog.Remove()
	}
}

func testProduceConsume(t *testing.T, client api.LogClient, config *Config) {
	ctx := context.Background()

	want := &api.Record{
		Value: []byte("hello world"),
	}

	produce, err := client.Produce(ctx,
		&api.ProduceRequest{
			Record: want,
		})
	assert.NoError(t, err)

	consume, err := client.Consume(ctx, &api.ConsumeRequest{
		Offset: produce.Offset,
	})
	assert.NoError(t, err)
	assert.Equal(t, want.Value, consume.Record.Value)
	assert.Equal(t, want.Offset, consume.Record.Offset)
}

func testConsumePastBoundary(t *testing.T, client api.LogClient, config *Config) {
	ctx := context.Background()

	produce, err := client.Produce(ctx, &api.ProduceRequest{
		Record: &api.Record{
			Value: []byte("hello world"),
		},
	})
	assert.NoError(t, err)

	consume, err := client.Consume(ctx, &api.ConsumeRequest{
		Offset: produce.Offset + 1,
	})
	if consume != nil {
		t.Fatal("consume not nil")
	}
	got := grpc.Code(err)
	want := grpc.Code(api.ErrOffsetOutOfRange{}.GRPCStatus().Err())
	if got != want {
		t.Fatalf("got err: %v, want: %v", got, want)
	}
}

func testProduceConsumeStream(t *testing.T, client api.LogClient, config *Config) {
	ctx := context.Background()

	records := []*api.Record{{
		Value:  []byte("first message"),
		Offset: 0,
	}, {
		Value:  []byte("second message"),
		Offset: 1,
	}}

	{
		stream, err := client.ProduceStream(ctx)
		assert.NoError(t, err)

		for offset, record := range records {
			err = stream.Send(&api.ProduceRequest{
				Record: record,
			})
			assert.NoError(t, err)
			res, err := stream.Recv()
			assert.NoError(t, err)
			if res.Offset != uint64(offset) {
				t.Fatalf("got offset: %d, want: %d", res.Offset, offset)

			}
		}
	}

	{
		stream, err := client.ConsumeStream(ctx, &api.ConsumeRequest{Offset: 0})
		assert.NoError(t, err)

		for i, record := range records {
			res, err := stream.Recv()
			assert.NoError(t, err)
			assert.Equal(t, res.Record, &api.Record{
				Value:  record.Value,
				Offset: uint64(i),
			})
		}
	}
}
