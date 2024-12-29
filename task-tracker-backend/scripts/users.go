package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type User struct{
    Name string `json:"username"`
    Email string `json:"email"`
    Pass string `json:"password"`
}

func CreateTempUsers(){
    var Users []User
    
    for i:=0; i < 5; i++{
        var u User
        u.Name = "sampleuser" + strconv.Itoa(i) 
        u.Email = u.Name + "@yahoo.com"
        u.Pass = "12345678"
        Users = append(Users, u)
    }

    for _, user := range Users{
        jsondata, err := json.Marshal(user)
        if err != nil{
            log.Fatal(err)
            return
        }
        
        reader := bytes.NewReader(jsondata)
        resp, err := http.Post("http://10.186.173.128:8090/users/create","application/json",reader)
        if err != nil{
            log.Fatal(err)
            return
        }
        
        if resp.StatusCode != http.StatusCreated{
            log.Println("Can not create user. Server replied with: ", resp.Status, resp.StatusCode)
        }
        log.Println("User : ", user.Name, " created")
    }

}

func main(){
    CreateTempUsers()
}
