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

var jwtSecret = []byte(viper.GetString("SECRET_KEY"))

func JWTInterceptor(ctx context.Context) (context.Context, error) {
	// ดึง metadata จาก request
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

func UnaryJWTInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	secureMethods := map[string]bool{
		"/pb.UserService/CreateCustomer": true,
	}

	log.Println("info.FullMethod : ", info.FullMethod, " _ ", !secureMethods[info.FullMethod])
	if !secureMethods[info.FullMethod] {
		log.Println("check auth")
		// ใช้ JWTInterceptor เพื่อตรวจสอบ JWT
		ctx, err := JWTInterceptor(ctx)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
	log.Println("no check auth")
	return handler(ctx, req)
}
