# Lazy tweets

Simple program to show/clean up inactive twitter's friendship

## Overview 

This is a simple program written in golang that allow you to find out among your twitter "friends" who is inactive for a long time.
An account is consider inactive when is not twitting since 300 days but you can change this value with something further or something closer.

## Prerequisites

 - Go Compiler 1.6 or greater [docs](https://golang.org/)
 - Git client

## Installation

    >> cd $GOPATH/src
    >> mkfir -p github.com/uolter
    >> git clone https://github.com/uolter/Lazy-tweets.git
    >> cd Lazy-tweets
    >> go get .
    >> go build -o lazytweets

## Configuration

Create a file called **.env** to allow the program to load  loads environment variables.
Edit the file and fill it with your [twitter configuration keys and access tokens](https://apps.twitter.com/).

.env

    export CONSUMER_KEY=<consumer key>
    export CONSUMER_SECRET=<consumer secret>
    export ACCESS_TOKEN=<access toke>
    export ACCESS_SECRET=<access secret>

## Run the program

    >> # Show the help
    >> ./lazytweets --help
    
    >> # Get lazy friends
    >> ./lazytweets 
    
    >> # Get inactive friends since the last 20 days
    >> ./lazytweets -inactive_after 20
    >> # Unfollow lazy friends
    >> ./lazytweets -unfollow

## Options

	Usage of ./lazytweets:
	  -alsologtostderr
    	log to standard error as well as files
	  -inactive_after int
    	Inactieve after n days. (default 300)
	  -log_backtrace_at value
	    	when logging hits line file:N, emit a stack trace (default :0)
	  -log_dir string
    	If non-empty, write log files in this directory
	  -logtostderr
    	log to standard error instead of files
	  -stderrthreshold value
    	logs at or above this threshold go to stderr
	  -unfollow
    	Unfollow inactive users.
	  -user_count int
    	Users to fetch from twitter. (default 200)
	  -v value
    	log level for V logs
	  -vmodule value
    	comma-separated list of pattern=N settings for file-filtered logging
