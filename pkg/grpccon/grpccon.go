package grpccon

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wiraphatys/intania-shop-boilerplate/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type grpcAuth struct {
	secretKey string
}

func (g *grpcAuth) unaryAuthorization(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}

	authHeader, ok := md["auth"]
	if !ok {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}

	if len(authHeader) == 0 {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}

	// verify api-key
	token, err := jwt.Parse(authHeader[0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(g.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	_, ok = claims["sub"].(string)
	if !ok {
		return nil, errors.New("not found subject in token")
	}

	return handler(ctx, req)
}

func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)
	grpcAuth := &grpcAuth{
		secretKey: cfg.ApiSecretKey,
	}
	opts = append(opts, grpc.UnaryInterceptor(grpcAuth.unaryAuthorization))
	grpcServer := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error: Failed to listen: %v", err)
	}
	return grpcServer, lis
}
