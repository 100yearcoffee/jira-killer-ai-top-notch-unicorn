package stats

import (
	"context"
	statsv1 "go-api/stats/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *grpc.ClientConn
	client statsv1.StatsServiceClient
}

func NewClient(target string) (*Client, error) {
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		client: statsv1.NewStatsServiceClient(conn),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetStats(ctx context.Context) (*statsv1.GetStatsResponse, error) {
	return c.client.GetStats(ctx, &statsv1.GetStatsRequest{})
}
