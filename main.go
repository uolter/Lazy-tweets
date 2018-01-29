package main

import (
	"flag"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	log "github.com/golang/glog"
	"github.com/subosito/gotenv"
)

const (
	DATE_LAYOUT string = time.RubyDate
)

var inactive_after int
var user_count int
var unfollow bool
var relationship bool
var exclude string

var p = log.Info

func init() {
	flag.IntVar(&inactive_after, "inactive_after", 300, "Inactieve after n days.")
	flag.IntVar(&user_count, "user_count", 200, "Users to fetch from twitter.")
	flag.BoolVar(&unfollow, "unfollow", false, "Unfollow inactive users.")
	flag.BoolVar(&relationship, "relationship", false, "Show which accunts you follow but they don't.")
	flag.StringVar(&exclude, "exclude", "", "List of account (comma separated to exclude)")
	flag.Lookup("logtostderr").Value.Set("true")

	gotenv.Load()
}

func last_post_days(date string) (days int) {

	t, err := time.Parse(DATE_LAYOUT, date)

	if err != nil {
		log.Error("Wrong date format", date, err)
		return 0
	}

	diff := time.Now().Sub(t)
	days = int(diff.Hours() / 24.0)
	return
}

func main() {

	flag.Parse()

	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	p("Friends")

	friends, _, _ := client.Friends.List(&twitter.FriendListParams{
		Count: user_count,
	})

	excludelist := strings.Split(exclude, ",")
	sort.Strings(excludelist)

	for _, user := range friends.Users {

		i := sort.SearchStrings(excludelist, user.ScreenName)
		if i < len(excludelist) && excludelist[i] == user.ScreenName {
			p("Skip ", user.ScreenName)
			continue
		}

		last_post_days_ago := last_post_days(user.Status.CreatedAt)

		if last_post_days_ago > inactive_after {
			p(user.Name, "(@"+user.ScreenName+")")
			p(last_post_days_ago, " days ago: ", user.Status.Text)
			if unfollow == true {
				client.Friendships.Destroy(&twitter.FriendshipDestroyParams{UserID: user.ID})
				p("Unfollowed!!")
			}

		}

		// User also following me
		if relationship == true {
			relashionship, _, err := client.Friendships.Show(&twitter.FriendshipShowParams{TargetID: user.ID})
			if err != nil {
				log.Error("Error in finding relationship ", err)
			} else if relashionship.Target.Following == false {
				p("", relashionship.Target.ScreenName, " (", relashionship.Target.IDStr, ")")
			}
		}
	}

	log.Flush()
}
