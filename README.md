# pixie-bot

A bot for pixie and [#pixie-lang](https://botbot.me/freenode/pixie-lang/).

## quickstart

you will need to have [docker](https://www.docker.com/) installed.

    # get the code
    $ git clone git@github.com:heyLu/pixie-bot.git
    $ cd pixie-bot
    # fetch dependencies (just an irc library at the moment)
    $ go get -d
    # build and run
    $ go build bot.go
    $ ./bot
    connected, hi there.
    ...

if you want to run a custom pixie version, you need to customize the docker
image used. to do this, recompile pixie with your changes and use `Dockerfile.pixie`
to build a custom image.

## todo

* a better name
* some thoughts regarding security
* something cat-related

## resources

* [documentation for goirc](http://godoc.org/github.com/fluffle/goirc/client)
* [lazybot](https://github.com/Raynes/lazybot)
