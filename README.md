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

## todo

* a better name
* some thoughts regarding security
* something cat-related
