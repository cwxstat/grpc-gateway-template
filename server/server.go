package server

import (
	"context"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pbExample "github.com/cwxstat/grpc-gateway-template/proto"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// Backend implements the protobuf interface
type Backend struct {
	mu         *sync.RWMutex
	namespaces map[string]*pbExample.Namespace
}

// New initializes a new Backend struct.
func New() *Backend {

	namespaces := map[string]*pbExample.Namespace{}

	// Grab current or defaults
	namespaces["system"] = &pbExample.Namespace{
		Name:       "system",
		CreateTime: timestamppb.Now(),
		Metadata:   &structpb.Struct{},
	}

	namespaces["default"] = &pbExample.Namespace{
		Name:       "default",
		CreateTime: timestamppb.Now(),
		Metadata:   &structpb.Struct{},
	}

	return &Backend{
		namespaces: namespaces,
		mu:         &sync.RWMutex{},
	}
}

// AddUser adds a namespace to the in-memory hash.
func (b *Backend) CreateNamespace(ctx context.Context, req *pbExample.CreateNamespaceRequest) (*pbExample.Namespace, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if value, ok := b.namespaces[req.GetNamespace()]; ok {
		return value, status.Errorf(codes.AlreadyExists, "namespace with name %q already exists", req.GetNamespace())

	}

	namespace := &pbExample.Namespace{
		Name:       req.GetNamespace(),
		CreateTime: timestamppb.Now(),
		Metadata:   req.GetMetadata(),
	}
	b.namespaces[req.GetNamespace()] = namespace

	return namespace, nil
}

// GetNamespace gets a namespace from the cluster.
func (b *Backend) GetNamespace(ctx context.Context, req *pbExample.GetNamespaceRequest) (*pbExample.Namespace, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if value, ok := b.namespaces[req.GetName()]; ok {
		return value, nil
	}

	return nil, status.Errorf(codes.NotFound, "namespace with name %q could not be found", req.GetName())
}

// DeleteNamespace deletes a namespace from the cluster.
func (b *Backend) DeleteNamespace(ctx context.Context, req *pbExample.GetNamespaceRequest) (*pbExample.Namespace, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if value, ok := b.namespaces[req.GetName()]; ok {
		delete(b.namespaces, req.GetName())
		return value, nil
	}

	return nil, status.Errorf(codes.NotFound, "namespace with name %q could not be found", req.GetName())
}

// ListNamespaces lists all namespaces in the cluster.
func (b *Backend) ListNamespaces(_ *pbExample.ListNamespaceRequest, srv pbExample.NamespaceService_ListNamespacesServer) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, namespace := range b.namespaces {
		err := srv.Send(namespace)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateNamespaces updates a namespace in the cluster.
func (b *Backend) UpdateNamespace(ctx context.Context, req *pbExample.UpdateNamespaceRequest) (*pbExample.Namespace, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if _, ok := b.namespaces[req.Namespace.Name]; ok {
		namespace := &pbExample.Namespace{
			Name: req.GetNamespace().Name,
			CreateTime: req.Namespace.GetCreateTime(),
			Metadata: req.GetNamespace().Metadata,
		}
		b.namespaces[req.Namespace.Name] = namespace
		return namespace, nil
	}

	return nil, status.Errorf(codes.NotFound, "namespace with name %q could not be found", req.Namespace.Name)

}
