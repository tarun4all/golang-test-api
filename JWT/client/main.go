package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// TODO: shift to env
// cmd: set JWT_TOKEN = shhhhhhhhhhhhhhhh
// var key = os.Get("JWT_TOKEN")
var mySigningKey = []byte("shhhhhhhhhhhh")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["autorized"] = true
	claims["user"] = "Tarun Bansal"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("something went wrong")
		return "", err
	}

	return tokenString, err

}

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:3000/", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(body))
}

func createServer() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":3001", nil))
}

func main() {
	fmt.Println("Client")

	createServer()
}
