package jwt

import (
	"TikTok/biz/dal/db"
	"TikTok/biz/model/user"
	"TikTok/biz/pack"
	"TikTok/biz/service"
	"TikTok/pkg/constants"
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"log"
	"time"
)

var (
	identityKey               = "userid"
	AccessTokenJwtMiddleware  *jwt.HertzJWTMiddleware
	RefreshTokenJwtMiddleware *jwt.HertzJWTMiddleware
)

func AccessTokenJwt() {
	var err error
	AccessTokenJwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:                       "Video",
		Key:                         []byte("AccessToken_key"),
		Timeout:                     time.Hour,
		MaxRefresh:                  time.Hour,
		WithoutDefaultTokenHeadName: true,
		TokenLookup:                 "header: Access-Token",
		IdentityKey:                 identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*db.User); ok {
				return jwt.MapClaims{
					identityKey: v.Id,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[AccessTokenJwtMiddleware.IdentityKey]
		},

		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			pack.SendFailResponse_UsedInJWT(c, code, message)
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.Set("Access-Token", token)
		},

		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct user.LoginRequest
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			users, err := db.LoginCheck(ctx, loginStruct.Username, loginStruct.Password)
			if err != nil {
				return nil, err
			}
			return users, nil
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}

func RefreshTokenJwt() {
	var err error
	RefreshTokenJwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "video zone",
		Key:         []byte(constants.RefreshTokenKey),
		Timeout:     time.Hour * 72,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: identityKey,
		TokenLookup: "header: Refresh-Token",
		//往令牌中添加的信息
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*db.User); ok {
				return jwt.MapClaims{
					identityKey: v.Id,
				}
			}
			return jwt.MapClaims{}
		},
		//从令牌中提取信息
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c) // 是从 JWT 令牌中提取 claims 的函数
			log.Printf("claims: %+v", claims)

			// 检查 claims[identityKey] 是否存在
			userID, exists := claims[identityKey]
			if !exists {
				log.Println("claims['userid'] 不存在")
				return nil
			}
			// 将 userID 存储到上下文中
			c.Set(constants.ContextUserId, userID)
			return userID
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			pack.SendFailResponse(c, errors.New(message))
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.Set("Refresh-Token", token)
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct user.LoginRequest
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			users, err := db.LoginCheck(ctx, loginStruct.Username, loginStruct.Password)
			if err != nil {
				return nil, err
			}
			return users, nil
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}

func IsAccessTokenAvailable(ctx context.Context, c *app.RequestContext) bool {
	claims, err := AccessTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return false
	}
	//JWT 的 Claims 中通常包含一个 exp 字段，表示 Token 的过期时间（Unix 时间戳）。
	switch v := claims["exp"].(type) { //switch-case处理解析出来的时间类型并与当前时间做比较
	case nil:
		return false
	case float64:
		if int64(v) < AccessTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return false
		}
		if n < AccessTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	default:
		return false
	}
	c.Set("JWT_PAYLOAD", claims) //将令牌存入上下文
	identity := AccessTokenJwtMiddleware.IdentityHandler(ctx, c)

	if identity != nil {
		c.Set(AccessTokenJwtMiddleware.IdentityKey, identity) //将用户id解析出存入上下文
	}
	if !AccessTokenJwtMiddleware.Authorizator(identity, ctx, c) { //
		return false
	}

	return true

}

func IsRefreshTokenAvailable(ctx context.Context, c *app.RequestContext) bool {
	claims, err := RefreshTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return false
	}

	switch v := claims["exp"].(type) {
	case nil:
		return false
	case float64:
		if int64(v) < RefreshTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return false
		}
		if n < RefreshTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	default:
		return false
	}

	c.Set("JWT_PAYLOAD", claims)
	identity := RefreshTokenJwtMiddleware.IdentityHandler(ctx, c)
	if identity != nil {
		c.Set(RefreshTokenJwtMiddleware.IdentityKey, identity)
	}
	if !RefreshTokenJwtMiddleware.Authorizator(identity, ctx, c) {
		return false
	}

	return true
}

func GenerateAccessToken(c *app.RequestContext) {
	data := service.GetUserIDFromContext(c)
	tokenString, _, _ := AccessTokenJwtMiddleware.TokenGenerator(data)
	c.Header("New-Access-Token", tokenString)
}

func Init() {
	AccessTokenJwt()
	RefreshTokenJwt()
	errInit := AccessTokenJwtMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("AccessTokenJwtMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	errInit = RefreshTokenJwtMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("RefreshTokenJwtMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}
