<h1 align="center" style="border-bottom: none;">:apple: Simplifier :apple:</h1>
<p align="center"><i>The first principle is that you must not fool yourself and you are the easiest person to fool.</i><br>–Richard Feynman</p>
<p align="center"><i>It's complicated.</i><br>–Facebook</p>

### Introduction

Simplifier is a tool for generating "ELI5ized" pieces of text from more complicated inputs. As of right now it is nothing more than a command line thesaurus.

### Setup

Word searches are done using the [Big Huge Thesaurus API][1]. You'll need an API key from there, which can then be set one of two ways:
- Passed as a flag: `simplifier --api-key <your api key> ...`
- Environment variable: `export BHT_API_KEY=<your api key>`

Using a cache will save bandwidth, latency, and request limits down the line when many repeated API calls are being made. Caching aside to a local Redis instance is currently hardcoded. The instance is assumed to be running at `localhost:6379`. If you have **Docker Compose** and **GNU Make**, you can simply run:

```sh
make startredis    # to start the Redis server, and...
make stopredis     # to stop the Redis server
```

### Usage

```sh
# to return a list of synonyms and other related words
simplifier thesaurus [word]
```


[1]: https://words.bighugelabs.com/api.php
