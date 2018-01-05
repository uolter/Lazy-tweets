package main

import (	
	"time"
	"os"
	"flag"
	
	log "github.com/golang/glog"
	"github.com/subosito/gotenv"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const (
	DATE_LAYOUT string = time.RubyDate
)

var inactive_after int
var user_count int
var unfollow bool

var p = log.Info

func init() {
	flag.IntVar(&inactive_after, "inactive_after", 300, "Inactieve after n days.")
	flag.IntVar(&user_count, "user_count", 200, "Users to fetch from twitter.")
	flag.BoolVar(&unfollow, "unfollow", false, "Unfollow iniactive users.")
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
	

	for _, user := range friends.Users {
		last_post_days_ago := last_post_days(user.Status.CreatedAt)

		if last_post_days_ago > inactive_after {
			p( user.Name ,"(@" + user.ScreenName + ")")
			p(last_post_days_ago , " days ago: " , user.Status.Text)
			
			if unfollow == true {
				client.Friendships.Destroy(&twitter.FriendshipDestroyParams{UserID: user.ID,})
				p("Unfollowed!!")
			}
			
		}
	}

	log.Flush()
}