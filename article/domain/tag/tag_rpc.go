package tag

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/victorzhou123/vicblog/common/controller/rpc"
	tagsvc "github.com/victorzhou123/vicblog/tag-server/domain/tag/service"
)

func NewTagServer(cfg *Config) (tagsvc.TagService, error) {

	conn, err := grpc.NewClient(cfg.toAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	tagGrpcClient := rpc.NewTagServiceClient(conn)

	return tagsvc.NewTagServer(cfg.toExpireTime(), tagGrpcClient), nil
}
