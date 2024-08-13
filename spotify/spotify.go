package spotify

import (
	"fmt"
	"net/http"
)

type APICredentials struct {
    
    
}

func Test() {
    fmt.Println("Test!")
}

func Auth_spotify_api() {
    req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", nil)
    if err != nil {
        panic("Could not create requesT")
    }
    req.Header.Add("Content-Type", "application/json")
}
