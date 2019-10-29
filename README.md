# Status protocol go bandwidth tests


## Description

This is a bandwidth test for status-protocol-go.
How it works:
New keys are always generated, each second (we could make this configurable) a message is sent to a random peer, this goes on until either we reach `MESSAGES` or we have been running for enough `SECONDS`. Both options are ignored if set to 0.

The tests are all run in the same container, so bandwidth usage is to be divided by the number of peers.

The code is very rough, I just got it working as quickly as I could, stats are recorded partly in containers and could be collated at the end of each tests (messages sent, the id of messages so we can check for dropped messages etc).


## Running

You need to have docker installed.


`./run.rb  -m 0 -s 20 -a 2 -d -g`

Where

`-m : Number of messages to be sent, if 0 it will be unlimited, otherwise the program will stop once it sends m messages from each peer.`

`-s : Number of seconds to run. 0 for disabling it`

`-a : Number of peers, needs to be at least 2`

`-d : Enable datasync (enabled in v1)`

`-g : Enable generic discovery topic (disabled in v1)`


Either `-m` or `-s` needs to be specified.


## Script

```
make build
```

```
./status-protocol-bandwidth-test -public-chat-id=testpublicchat -src=test2 -port=30304 -dst=test1 -messages=100
```
