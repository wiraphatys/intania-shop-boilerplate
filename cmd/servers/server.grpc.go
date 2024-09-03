package servers

import (
	"context"
	"fmt"
	"time"

	"github.com/wiraphatys/intania-shop-boilerplate/internal/products"
	"github.com/wiraphatys/intania-shop-boilerplate/pkg/config"
	"github.com/wiraphatys/intania-shop-boilerplate/pkg/grpccon"
	"github.com/wiraphatys/intania-shop-boilerplate/proto/product_pb"
	"github.com/wiraphatys/intania-shop-boilerplate/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type GrpcServer struct {
	cfg    config.Config
	logger *zap.Logger
	db     *gorm.DB
}

func NewGrpcServer(cfg config.Config, logger *zap.Logger, db *gorm.DB) *GrpcServer {
	return &GrpcServer{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}

func (srv *GrpcServer) Start(
	productHandler *products.ProductGrpcHandler,
) {
	url := fmt.Sprintf("%v:%v", srv.cfg.GetServer().Host, srv.cfg.GetServer().Port)
	jwtConfig := srv.cfg.GetJwt()

	grpcServer, listener := grpccon.NewGrpcServer(&jwtConfig, url)
	product_pb.RegisterProductServiceServer(grpcServer, productHandler)

	reflection.Register(grpcServer)
	go func() {
		srv.logger.Sugar().Infof("ESC Intania Shop Backend starting at port %v", srv.cfg.GetServer().Port)

		if err := grpcServer.Serve(listener); err != nil {
			srv.logger.Fatal("Failed to start ESC Intania Shop Backend service", zap.Error(err))
		}
	}()

	wait := utils.GracefulShutdown(context.Background(), 2*time.Second, srv.logger, map[string]utils.Operation{
		"server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
		"database": func(ctx context.Context) error {
			sqlDB, err := srv.db.DB()
			if err != nil {
				return nil
			}
			return sqlDB.Close()
		},
	})

	<-wait

	grpcServer.GracefulStop()
	srv.logger.Info("Closing the listener")
	listener.Close()
	srv.logger.Info("ESC Intania Shop Backend service has been shutdown gracefully")
}
