/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim ÇOBANİ
	Date        : 06.03.2023  Pazartesi
	Time        : 17:50
	Notes       :

*/

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"sandbox.101.icibot.net/Config"
	"sandbox.101.icibot.net/models"
	"strings"
	"time"
)

// CheckToken : Token kontrolü yapar. Middleware olarak kullanılır.
func CheckToken(c *gin.Context) {
	// burayı kaldıracağız sonradan.
	log.Println("Access-Control-Allow-Origin eklendi.")
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
	}

	var tokenStr string = c.Request.Header.Get("Authorization")
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 10)
	tokenStr = strings.Replace(tokenStr, "bearer ", "", 10)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.AppSecretKey), nil
	})
	if err == nil && token.Valid {
		var user = models.User{}
		Config.Db.First(&user,
			int64(ParseToken(tokenStr, "id").(float64)),
		)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Lütfen tekrar sisteme giriş yapınız."})
			c.Abort()
		}
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Lütfen tekrar sisteme giriş yapınız."})
		c.Abort()
	}
}

// ParseToken : Token içerisindeki bilgileri parse eder.
func ParseToken(myToken string, ClaimsName string) interface{} {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.AppSecretKey), nil
	})

	if err == nil && token.Valid {
		return token.Claims.(jwt.MapClaims)[ClaimsName]
	} else {
		return ""
	}
}

// CreateToken : Token oluşturur.
func CreateToken(user models.User) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	Claims := make(jwt.MapClaims)
	Claims["id"] = user.ID
	Claims["exp"] = time.Now().Add(Config.TokenExpireDuration).Unix()
	Claims["expd"] = time.Now().Add(Config.TokenExpireDuration).Format("2006-01-02T15:04:05.000Z")
	Claims["iat"] = time.Now().Unix()
	token.Claims = Claims
	tokenString, err := token.SignedString([]byte(Config.AppSecretKey))
	return tokenString, err
}

// context içinde tanımlı olan kullanıcı bilgilerine ve firma bilgilerine erişim sağlanır.
func GetUser(c *gin.Context) models.User {
	var tokenStr string = c.Request.Header.Get("Authorization")
	return getUserByToken(tokenStr)
}

func getUserByToken(tokenStr string) models.User {
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 10)
	tokenStr = strings.Replace(tokenStr, "bearer ", "", 10)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.AppSecretKey), nil
	})

	if err == nil && token.Valid {
		var user = models.User{}
		Config.Db.First(&user, int64(ParseToken(tokenStr, "id").(float64)))
		return user

	} else {
		log.Println(err)
		return models.User{}
	}
	return models.User{}
}
