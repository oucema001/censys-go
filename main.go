package main

import (
	"context"
	"fmt"

	"github.com/oucema001/censys-go/censys"
)

func main() {

	/*var keys struct {
		Appid     string `json:"appid"`
		Appsecret string `json:"appsecret"`
	}
	f, err := os.Open("keys.json")
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(f)
	dec.Decode(&keys)
	var req *http.Request
	req, err = http.NewRequest("GET", "https://censys.io/api/v1/data", strings.NewReader(""))
	if err != nil {
		log.Panic(err)
	}
	req.SetBasicAuth(keys.Appid, keys.Appsecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	io.Copy(os.Stdout, res.Body)
	var token oauth2.Token
	dec = json.NewDecoder(res.Body)
	err = dec.Decode(&token)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	conf := &oauth2.Config{
		/* 	ClientID:     keys.Appid,
		ClientSecret: keys.Appsecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/v3.2/dialog/oauth",
			TokenURL: "https://graph.facebook.com/v3.2/oauth/access_token",
		},
	}
	fbclient := conf.Client(ctx, &token)
	res2, err := fbclient.Get("https://graph.ffacebook.com/me/photos")
	if err != nil {
		panic(err)
	}
	defer res2.Body.Close()
	io.Copy(os.Stdout, res2.Body)*/
	/*client := censys.Client{
		ApiID:     "d8b4925a-2887-41d6-bbf5-6a00ea7dc529",
		ApiSecret: "q8qbfwgaEIJ3BYbLnVfHHzFIWvzd6Prb",
	}*/
	client := censys.NewClient(nil, "d8b4925a-2887-41d6-bbf5-6a00ea7dc529", "q8qbfwgaEIJ3BYbLnVfHHzFIWvzd6Prb")
	/*	p, err := client.GetProfile(context.Background())
			if err != nil {
				fmt.Println(err)
			}

		a, err := client.Search(context.Background(), "www.google.com", censys.WEBSITES)
		if err != nil {
			fmt.Println(err)
		}*/
	a, err := client.GetView(context.Background(), censys.WEBSITESVIEW, "google.com")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v+", a)
}
