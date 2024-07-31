package interceptor

import (
	"context"
	"log/slog"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type MethodName string

type Permission string

type PermisionOfMethod map[MethodName]Permission

type AuthFunc func(ctx context.Context, token string, device string) (map[string]string, error)

func checkPermission(mPermissions map[string]string, pm Permission) bool {
	// kiểm tra toàn quyền
	ad, ok := mPermissions["*"]
	if ok && ad != "" {
		return true
	}
	// kiểm tra quyền hạn
	ad, ok = mPermissions[string(pm)]
	if !ok {
		return false
	} else {
		if ad != "" {
			return true
		} else {
			return false
		}
	}
}

func AuthUnaryIntercepter(authFunc AuthFunc, pom PermisionOfMethod) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		parts := strings.Split(info.FullMethod, "/")
		methodName := parts[len(parts)-1]
		pm, ok := pom[MethodName(methodName)]
		if ok {
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, status.Error(codes.InvalidArgument, "missing metadata")
			}
			token, ok := md["authorization"]
			if !ok {
				return nil, status.Error(codes.Internal, "token required")
			}
			device, ok := md["device"]
			if !ok {
				return nil, status.Error(codes.Internal, "device required")
			}
			mPermissions, err := authFunc(ctx, token[0], device[0])
			if err != nil {
				log := slog.With("method", "AuthUnaryIntercepter")
				logGroup := log.With("type", "UnaryException", "func", "AuthFunc").WithGroup("data").With("token", token[0], "device", device[0])
				logGroup.Error(err.Error())
				return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
			}
			if mPermissions == nil {
				return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
			}
			check := checkPermission(mPermissions, pm)
			if !check {
				return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
			}

		}
		return handler(ctx, req)
	}
}
func AuthStreamIntercepter(authFunc AuthFunc, pom PermisionOfMethod) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		parts := strings.Split(info.FullMethod, "/")
		methodName := parts[len(parts)-1]
		pm, ok := pom[MethodName(methodName)]
		if ok {
			md, ok := metadata.FromIncomingContext(ss.Context())
			if !ok {
				return status.Error(codes.InvalidArgument, "missing metadata")
			}
			token, ok := md["authorization"]
			if !ok {
				return status.Error(codes.Internal, "token required")
			}
			device, ok := md["device"]
			if !ok {
				return status.Error(codes.Internal, "device required")
			}
			mPermissions, err := authFunc(ss.Context(), token[0], device[0])
			if err != nil {
				log := slog.With("method", "AuthUnaryIntercepter")
				logGroup := log.With("type", "UnaryException", "func", "AuthFunc").WithGroup("data").With("token", token[0], "device", device[0])
				logGroup.Error(err.Error())
				return status.Errorf(codes.Unauthenticated, "unauthorized")
			}
			if mPermissions == nil {
				return status.Errorf(codes.Unauthenticated, "unauthorized")
			}
			check := checkPermission(mPermissions, pm)
			if !check {
				return status.Errorf(codes.Unauthenticated, "unauthorized")
			}
		}
		return handler(srv, ss)
	}
}
