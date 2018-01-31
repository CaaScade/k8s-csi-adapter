package identity

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

const (
	DriverName    = "doopty-driver"
	DriverVersion = "0.1.0"
)

type Server struct {
}

func (s *Server) GetSupportedVersions(
	ctx context.Context,
	req *csi.GetSupportedVersionsRequest) (
	*csi.GetSupportedVersionsResponse, error) {

	// Allow csp to handle this (? thecodeteam/csi-scalio/service/identity.go)
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Server) GetPluginInfo(
	ctx context.Context,
	req *csi.GetPluginInfoRequest) (
	*csi.GetPluginInfoResponse, error) {
	return &csi.GetPluginInfoResponse{
		Name:          DriverName,
		VendorVersion: DriverVersion,
	}, nil
}
