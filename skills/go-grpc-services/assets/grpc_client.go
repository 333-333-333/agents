// api/booking/internal/booking/infrastructure/grpc_caregiver_client.go
package infrastructure

import (
	"api/booking/internal/booking/domain"
	pb "api/caregiver/proto/caregiverv1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// CaregiverClient implements domain.CaregiverGateway port.
type CaregiverClient struct {
	client pb.CaregiverServiceClient
	conn   *grpc.ClientConn
}

func NewCaregiverClient(addr string) (*CaregiverClient, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to caregiver service: %w", err)
	}

	return &CaregiverClient{
		client: pb.NewCaregiverServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *CaregiverClient) GetCaregiver(ctx context.Context, id string) (*domain.Caregiver, error) {
	resp, err := c.client.GetCaregiver(ctx, &pb.GetCaregiverRequest{CaregiverId: id})
	if err != nil {
		return nil, mapGRPCToDomainError(err)
	}
	return toDomain(resp.GetCaregiver()), nil
}

func (c *CaregiverClient) Close() error {
	return c.conn.Close()
}

func mapGRPCToDomainError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.NotFound:
		return domain.ErrCaregiverNotFound
	case codes.Unavailable:
		return fmt.Errorf("caregiver service unavailable: %w", err)
	default:
		return fmt.Errorf("caregiver service error: %w", err)
	}
}
