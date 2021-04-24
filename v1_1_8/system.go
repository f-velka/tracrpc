package v1_1_8

import (
	"errors"

	tracrpc "github.com/f-velka/go-trac-rpc"
)

const (
	system_multicall        string = "system.multicall"
	system_list_methods     string = "system.listMethods"
	system_method_help      string = "system.methodHelp"
	system_method_signature string = "system.methodSignature"
	system_get_API_version  string = "system.getAPIVersion"
)

// SystemService represents system API service.
type SystemService struct {
	rpc tracrpc.RpcClient
}

// newSystemService creates new SystemService instance.
func newSystemService(rpc tracrpc.RpcClient) (*SystemService, error) {
	if rpc == nil {
		return nil, errors.New("rpc client cannot be nil")
	}

	return &SystemService{
		rpc: rpc,
	}, nil
}

// TODO: system.multicall

// ListMethods calls system.listMethods.
func (s *SystemService) ListMethods() ([]string, error) {
	var reply []string
	if err := s.rpc.Call(system_list_methods, nil, &reply); err != nil {
		return nil, err
	}

	return reply, nil
}

// MethodHelp calls system.methodHelp.
func (s *SystemService) MethodHelp(methodName *string) (string, error) {
	args := packArgs(methodName)
	var reply string
	if err := s.rpc.Call(system_method_help, args, &reply); err != nil {
		return "", err
	}

	return reply, nil
}

// MethodSignature calls system.methodSignature.
func (s *SystemService) MethodSignature(methodName *string) ([]string, error) {
	args := packArgs(methodName)
	var reply []string
	if err := s.rpc.Call(system_method_signature, args, &reply); err != nil {
		return nil, err
	}

	return reply, nil
}

// GetAPIVersion calls system.getAPIVersion.
func (s *SystemService) GetAPIVersion() ([]int, error) {
	var reply []int
	if err := s.rpc.Call(system_get_API_version, nil, &reply); err != nil {
		return nil, err
	}

	return reply, nil
}
