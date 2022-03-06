package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomUserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *MyCustomUserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("invalid session ID")
	}

	return nil
}

func (c *MyCustomUserClaims) CreateToken() (string, error) {
	// we are using HMAC(type SigningMethodHMAC), since SigningMethodHS512 is also of type-SigningMethodHMAC,
	// And Claims is a Interface implementing Valid() method, and c is of type-MyCustomUserClaims, MyCustomUserClaims is reciver type of Valid() method.
	// which means MyCustomUserClaims has Valid() method
	// so MyCustomUserClaims type is Claims-type
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken,err := t.SignedString(NEED_KEY) // NEED_KEY can be created by HMAC
	if err !=nil {
		fmt.Println(err.Error())
	}

	return signedToken, nil
}

func parseToken(signedToken string) (*MyCustomUserClaims, error) {
	t, err := jwt.ParseWithClaims(signedToken, &MyCustomUserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() == jwt.SigningMethodES512.Alg() {
			return nil, fmt.Errorf("invalid Signing algorithm")
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error in parseToken")
	}

	if !t.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	return t.Claims.(*MyCustomUserClaims), nil
}


func main() {
	
}