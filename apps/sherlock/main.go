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
	defer twitterClient.Close()

	followerIds := twitterClient.GetFollowersIdsAll(url.Values{})

	for page := range followerIds {
		if page.Error != nil {
			panic(page.Error)
		}

		// you hit a limit here if you send all the ids, need to loop through here
		// see https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-lookup.html
		// "up to 100 are allowed in a single request."
		users, err := twitterClient.GetUsersLookupByIds(page.Ids[:30], url.Values{})
		if err != nil {
			panic(err)
		}

		for _, user := range users {
			fmt.Println(user.Name)
		}
	}
}
