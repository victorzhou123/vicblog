package category

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	catesvc "github.com/victorzhou123/vicblog/category-server/domain/category/service"
	"github.com/victorzhou123/vicblog/common/controller/rpc"
)

func NewCategoryServer(cfg *Config) (catesvc.CategoryService, error) {

	conn, err := grpc.NewClient(cfg.toAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	categoryGrpcClient := rpc.NewCategoryServiceClient(conn)

	return catesvc.NewCategoryServer(cfg.toExpireTime(), categoryGrpcClient), nil
}
