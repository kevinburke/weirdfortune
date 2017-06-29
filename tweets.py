import pprint
import sys
import HTMLParser

from clint.textui import colored
import tweepy

# == OAuth Authentication ==
#
# This mode of authentication is the new preferred way
# of authenticating with Twitter.

# The consumer keys can be found on your application's Details
# page located at https://dev.twitter.com/apps (under "OAuth settings")
consumer_key=""
consumer_secret=""

# The access tokens can be found on your applications's Details
# page located at https://dev.twitter.com/apps (located
# under "Your access token")
access_token=""
access_token_secret=""

auth = tweepy.OAuthHandler(consumer_key, consumer_secret)
auth.set_access_token(access_token, access_token_secret)

api = tweepy.API(auth)

def format_tweet(username, tweet):
    return u'%\n@{}: {}\n'.format(username, tweet).encode('utf-8')

def add_tweet(username, tweet):
    with open('games/weirdfortunes/weirdfortunes', 'a') as f:
        f.write(format_tweet(username, tweet))

# If the authentication was successful, you should
# see the name of the account print out

import argparse

parser = argparse.ArgumentParser()
parser.add_argument('--id', required=False)
args = parser.parse_args()
print args

def main():
    h = HTMLParser.HTMLParser()
    oldest_id = long(args.id) if args.id else sys.maxint
    try:
        for item in tweepy.Cursor(api.user_timeline).items(2000):
            if item.id >= oldest_id:
                print "seen this tweet already"
                continue
            if not hasattr(item, 'retweeted_status'):
                continue
            username = item.retweeted_status.user.screen_name
            tweet = h.unescape(item.retweeted_status.text)
            if 'http://' in tweet:
                continue
            if 'https://' in tweet:
                continue
            if username == 'fanfiction_txt':
                continue
            print colored.red('@{}'.format(username))
            print colored.blue(tweet)
            print "Add this tweet?"
            answer = raw_input()
            if answer == 'y' or answer == '':
                add_tweet(username, tweet)
    except:
        if 'item' in locals():
            print "Current tweet ID is {}".format(item.id)
        raise

if __name__ == "__main__":
    main()
