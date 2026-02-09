// api/caregiver/internal/caregiver/infrastructure/handler/grpc.go
package handler

import (
	"api/caregiver/internal/caregiver/application"
	pb "api/caregiver/proto/caregiverv1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CaregiverGRPCHandler struct {
	pb.UnimplementedCaregiverServiceServer
	service *application.CaregiverService
}

func NewCaregiverGRPCHandler(service *application.CaregiverService) *CaregiverGRPCHandler {
	return &CaregiverGRPCHandler{service: service}
}

func (h *CaregiverGRPCHandler) GetCaregiver(ctx context.Context, req *pb.GetCaregiverRequest) (*pb.GetCaregiverResponse, error) {
	caregiver, err := h.service.GetByID(ctx, req.GetCaregiverId())
	if err != nil {
		return nil, mapDomainErrorToGRPC(err)
	}

	return &pb.GetCaregiverResponse{
		Caregiver: toProtobuf(caregiver),
	}, nil
}

func mapDomainErrorToGRPC(err error) error {
	// Map domain errors to gRPC status codes
	var notFound domain.NotFoundError
	if errors.As(err, &notFound) {
		return status.Error(codes.NotFound, err.Error())
	}
	return status.Error(codes.Internal, "internal error")
}
