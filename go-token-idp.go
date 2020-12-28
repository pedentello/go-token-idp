package main

import (
    "fmt"
    "log"
    "net/http"
    "net/url"
    "io/ioutil"
    "encoding/json"
    "github.com/Kong/go-pdk"

)

type token struct {
    Access_token string `json:"access_token"`
}

// it represents to config parameters into the config.yml
type Config struct {
    URL_IDP string
    Client_id string
    Client_secret string
    Username string
    Password string
}

func New() interface{} {
    return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
	
   url_idp := conf.URL_IDP
   client_id := conf.Client_id
   client_secret := conf.Client_secret
   username := conf.Username
   password := conf.Password

   kong.Log.Info(fmt.Sprintf("URL %s ", url_idp))
   
   //We make HTTP request using the POST function

   resp, err := http.PostForm(url_idp, 
   url.Values{"grant_type": {"password"}, "client_id": {client_id}, "client_secret": {client_secret},
   "username": {username}, "password": {password} } )
   
   if err != nil {
      log.Fatal("ooopsss an error occurred in POST, please try again")
   }
   
   defer resp.Body.Close()

   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
       panic(err)
   }
   
   newToken := token{}
   
   jsonErr := json.Unmarshal(body, &newToken)
   if jsonErr != nil {
        log.Fatal(jsonErr)
   }
   
   kong.Log.Info(fmt.Sprintf("Access token %s ", newToken.Access_token))

   kong.ServiceRequest.AddHeader("Authorization", fmt.Sprintf("Bearer %s", newToken.Access_token))
  
   
}
