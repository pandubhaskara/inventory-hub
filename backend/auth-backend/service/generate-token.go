package service

import (
	"fmt"
	"strconv"
	"time"

	"credential-auth/config"
	"credential-auth/database"
	"credential-auth/model"

	jwt "github.com/golang-jwt/jwt/v4"
)

type myClaims struct {
	jwt.StandardClaims
	Id    uint
	Email string
}

func GenerateToken(user model.User, ClientID int) (map[string]string, error) {
	acDuration, _ := strconv.ParseInt(config.Jwt.ATDuration, 10, 64)
	accessToken := model.AccessToken{
		UserID:    int(user.ID),
		ClientID:  ClientID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add((time.Second * time.Duration(acDuration))),
		IsActive:  true,
	}
	database.Db.Create(&accessToken)

	newAccessToken, err := token(int32(user.ID), int32(accessToken.ID), acDuration)
	if err != nil {
		return nil, err
	}

	rtDuration, _ := strconv.ParseInt(config.Jwt.RTDuration, 10, 64)
	refreshToken := model.RefreshToken{
		AccessTokenID: int(accessToken.ID),
		IssuedAt:      time.Now(),
		ExpiredAt:     time.Now().Add((time.Second * time.Duration(rtDuration))),
		IsActive:      true,
	}
	database.Db.Create(&refreshToken)
	fmt.Println(refreshToken.ID)

	newRefreshToken, err := token(int32(accessToken.ID), int32(refreshToken.ID), rtDuration)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	}, nil
}

func token(ID int32, TokenID int32, duration int64) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Subject:   fmt.Sprint(ID),
			Id:        fmt.Sprint(TokenID),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add((time.Second * time.Duration(duration))).Unix(),
			Issuer:    config.Jwt.Issuer,
		},
	)

	signedToken, err := token.SignedString([]byte(config.Jwt.Signature))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
