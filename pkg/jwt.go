package pkg

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userID string) string
	VerifyToken(token string) (*jwtCustomClaim, error)
}

type jwtCustomClaim struct {
	SID string `json:"pid"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "ebook",
		secretKey: "f69f0580a78ff1aa841cdde6780f76e22164694c6f64ef11dccc396ff158baa8",
	}
}
func (s *jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    s.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, e := token.SignedString([]byte(s.secretKey))
	if e != nil {
		panic(e)
	}
	return t
}
func (maker *jwtService) VerifyToken(token string) (*jwtCustomClaim, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected: %v", token.Header["alg"])
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &jwtCustomClaim{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && verr != nil {
			return nil, err
		}
		return nil, err
	}

	payload, ok := jwtToken.Claims.(*jwtCustomClaim)
	if !ok {
		return nil, err
	}

	return payload, nil
}
