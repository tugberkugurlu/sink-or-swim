package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/spf13/viper"
	"net/url"
	"sync"
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

	baseFollowers, err := getFollowers(twitterClient)
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("will sleep for a minute now...")
		time.Sleep(1 * time.Minute)

		followers, err := getFollowers(twitterClient)
		if err != nil {
			panic(err)
		}

		diffs := compare(baseFollowers, followers)
		fmt.Println("Number of diffs: ", len(diffs))
		for userId := range diffs {
			// check the diff one to see where the value exists.
			// 		If it exists inside the source map, it means unfollow
			//      If it exists inside the new map, it means follow

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

type userWithId struct {
	Id int64
	Name string
}

// you hit a limit here if you send all the ids, need to loop through here
// see https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-lookup.html
// "up to 100 are allowed in a single request."
func getFollowers(twitterClient *anaconda.TwitterApi) (map[int64]string, error) {
	start := time.Now()
	usersMap := make(map[int64]string)
	followerIds := twitterClient.GetFollowersIdsAll(url.Values{})
	for page := range followerIds {
		if page.Error != nil {
			return nil, page.Error
		}

		followerCount := len(page.Ids)
		chunks := followerCount / 100
		lastChunk := followerCount % 100

		var wg sync.WaitGroup
		usersChan := make(chan userWithId, followerCount)

		for i := 0; i <= chunks; i++ {
			start := i * 100
			end := start + 100
			if i == chunks {
				if lastChunk == 0 {
					break
				}
				end = start + lastChunk
			}

			wg.Add(1)
			go func (userIds []int64) {
				defer wg.Done()
				users, err := twitterClient.GetUsersLookupByIds(userIds, url.Values{})
				if err != nil {
					// TODO: Do not panic here, propagate the error back to the main goroutine
					panic(err)
				}

				for _, user := range users {
					usersChan <- userWithId{Id: user.Id, Name: user.Name}
				}
			}(page.Ids[start:end])
		}

		go func() {
			wg.Wait()
			close(usersChan)
		}()

		for user := range usersChan {
			usersMap[user.Id] = user.Name
		}

		elapsed := time.Since(start)
		fmt.Println("Number of followers retrieved: ", len(usersMap))
		fmt.Printf("getFollowers took %s\n", elapsed)
	}

	return usersMap, nil
}
