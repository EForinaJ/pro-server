package jwt

import (
	"os"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenToken 生成JWT
func GenAdminToken(userId int64, username string) (string, error) {

	expire, err := g.Cfg().Get(gctx.New(), "jwt.admin.expire")
	if err != nil {
		return "", err
	}

	// 创建一个我们自己声明的数据
	claims := CustomClaims{
		userId,
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire.Duration() * time.Hour)), //有效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                    //签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                    //生效时间
			Issuer:    os.Getenv("JWT_ISSUER"),                                           //签发人
			Subject:   "admin",                                                           //主题
			ID:        gconv.String(userId),                                              //JWT ID用于标识该JWT
			Audience:  []string{"admin"},                                                 //用户
		},
	}

	//使用指定的加密方式和声明类型创建新令牌
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret, err := g.Cfg().Get(gctx.New(), "jwt.admin.secret")
	if err != nil {
		return "", err
	}
	//获得完整的、签名的令牌
	token, err := tokenStruct.SignedString(secret.Bytes())
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseAdminToken(token string) (*CustomClaims, error) {
	secret, err := g.Cfg().Get(gctx.New(), "jwt.admin.secret")
	if err != nil {
		return nil, err
	}

	//解析、验证并返回token。
	tokenObj, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret.Bytes(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenObj.Claims.(*CustomClaims); ok && tokenObj.Valid {
		return claims, nil
	} else {
		return nil, gerror.New("jwt-token get filde")
	}
}
