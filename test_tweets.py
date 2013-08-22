# -*- coding: utf-8 -*-
import tweets

txt = u'“butt dial”'
print txt
a = tweets.format_tweet('foo', txt)
print a
assert a == u'%\n@foo: “butt dial”\n', u"got {}".format(a).encode('utf-8')
