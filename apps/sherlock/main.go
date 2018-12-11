package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/spf13/viper"
	"net/url"
)

type Difference struct {
	Id int64
	IsFollower bool
	IsUnfollower bool
}

// see https://mrekucci.blogspot.com/2015/07/dont-abuse-mathmax-mathmin.html
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func compare(before, after []int64) (differences []int64) {
	minCount := min(len(before), len(after))
	for i := 0; i < minCount; i++ {
	}
	return
}

func getDifferences(before, after map[int64]string) (differences []Difference) {
	return
}

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
		followerCount := len(page.Ids)
		chunks := followerCount / 100
		lastChunk := followerCount % 100

		usersMap := make(map[int64]string)
		for i := 0; i <= chunks; i++ {
			start := i * 100
			end := start + 100
			if i == chunks {
				end = start + lastChunk
			}

			users, err := twitterClient.GetUsersLookupByIds(page.Ids[start:end], url.Values{})
			if err != nil {
				panic(err)
			}

			for _, user := range users {
				usersMap[user.Id] = user.Name
				fmt.Println(user.Name)
			}
		}

		fmt.Println("Number of followers retrieved: ", len(usersMap))

		// 1: put followers into a map in first iteration
		// 2: sleep for a minute
		// 3: put followers into another map on second iteration
		// 4: diff and return a slice of values that are only present in one of the input slices
		// 5: check that slice one by one and see where the value exists.
		// 		If it exists inside the source map, it means unfollow
		//      If it exists inside the new map, it means follow
		// 6: record follow/unfollow events
		// 7: make the new map the source map
		// 8: goto 2
	}
}
