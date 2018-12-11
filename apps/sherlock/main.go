package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/spf13/viper"
	"net/url"
	"time"
)

// This function runs in O(N + M) time complexity, where the length of N and M
// corresponds to the length of before and after maps respectively.
func compare(before, after map[int64]string) map[int64]bool {
	differences := make(map[int64]bool)

	// O(N)
	for key := range before {
		// O(1)
		if _, ok := after[key]; !ok {
			differences[key] = true
		}
	}

	// O(M)
	for key := range after {
		// O(1)
		if _, ok := before[key]; !ok {
			differences[key] = true
		}
	}

	return differences
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

	baseFollowers := getFollowers(twitterClient)
	for {
		fmt.Println("will sleep for a minute now...")
		time.Sleep(1 * time.Minute)

		followers := getFollowers(twitterClient)
		diffs := compare(baseFollowers, followers)
		fmt.Println("Number of diffs: ", len(diffs))
		for userId := range diffs {
			if name, ok := baseFollowers[userId]; ok {
				// unfollow happened
				fmt.Println("Unfollow: ", userId, name)
			} else {
				// follow happened
				followerName := followers[userId]
				fmt.Println("Follow: ", userId, followerName)
			}
		}

		baseFollowers = followers
	}
}

func getFollowers(twitterClient *anaconda.TwitterApi) map[int64]string {
	usersMap := make(map[int64]string)
	followerIds := twitterClient.GetFollowersIdsAll(url.Values{})
	for page := range followerIds {
		if page.Error != nil {
			// TODO: Don't panic here, return err instead
			panic(page.Error)
		}

		// you hit a limit here if you send all the ids, need to loop through here
		// see https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-lookup.html
		// "up to 100 are allowed in a single request."
		followerCount := len(page.Ids)
		chunks := followerCount / 100
		lastChunk := followerCount % 100

		for i := 0; i <= chunks; i++ {
			start := i * 100
			end := start + 100
			if i == chunks {
				if lastChunk == 0 {
					break
				}
				end = start + lastChunk
			}

			users, err := twitterClient.GetUsersLookupByIds(page.Ids[start:end], url.Values{})
			if err != nil {
				// TODO: Don't panic here, return err instead
				panic(err)
			}

			for _, user := range users {
				usersMap[user.Id] = user.Name
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

	return usersMap
}
