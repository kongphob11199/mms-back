package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Secret key สำหรับตรวจสอบ JWT
var jwtSecret = []byte(viper.GetString("SECRET_KEY"))

// Middleware สำหรับตรวจสอบ JWT
func JWTInterceptor(ctx context.Context) (context.Context, error) {
	// ดึง metadata จาก request
	log.Println("JWTInterceptor : ", viper.GetString("SECRET_KEY"))
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	// ดึงค่า Authorization header
	authHeader := md["authorization"]
	if len(authHeader) == 0 {
		return nil, errors.New("authorization token not provided")
	}

	// ตรวจสอบและแยก Bearer Token
	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
	if tokenString == authHeader[0] {
		return nil, errors.New("malformed token")
	}

	// ตรวจสอบ JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// ตรวจสอบว่า Signing Method ตรงกันหรือไม่
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// เพิ่มข้อมูลจาก JWT ลงใน context หากต้องการ
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		ctx = context.WithValue(ctx, "user", claims["user"])
	}

	return ctx, nil
}

// UnaryInterceptor สำหรับ gRPC
func UnaryJWTInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// ระบุ RPC methods ที่ต้องการตรวจสอบ JWT
	secureMethods := map[string]bool{
		"/your.package.Service/ProtectedMethod": true, // ใส่ชื่อ method ที่ต้องการ
	}

	// ตรวจสอบว่า method นี้ต้องการ JWT หรือไม่
	log.Println("info.FullMethod : ", info.FullMethod)
	if !secureMethods[info.FullMethod] {
		// ใช้ JWTInterceptor เพื่อตรวจสอบ JWT
		ctx, err := JWTInterceptor(ctx)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}

	// หาก method นี้ไม่ต้องการ JWT ก็ผ่านได้เลย
	return handler(ctx, req)
}
