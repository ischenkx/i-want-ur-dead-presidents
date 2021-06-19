package client

import entitiesGrpcGen "github.com/ischenkx/innotech-backend/services/entities/implementation/grpc/pb/generated"

type Client struct {
	entitiesGrpcGen.ProductsClient
}
