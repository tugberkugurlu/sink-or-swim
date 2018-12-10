package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/spf13/viper"
	"net/url"
)

func main() {
	viper.SetEnvPrefix("SHERLOCK")
	viper.AutomaticEnv()

	accessToken := viper.GetString("TWITTER__ACCESS__TOKEN")
	accessTokenSecret := viper.GetString("TWITTER__ACCESS__TOKEN__SECRET")
	consumerKey := viper.GetString("TWITTER__CONSUMER__KEY")
	consumerSecret := viper.GetString("TWITTER__CONSUMER__SECRET")

	twitterClient := anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)
	followerIds := twitterClient.GetFollowersIdsAll(url.Values{})

	for page := range followerIds {
		if page.Error != nil {
			panic(page.Error)
		}

		for _, id := range page.Ids {
			fmt.Println(id)
		}
	}
}
