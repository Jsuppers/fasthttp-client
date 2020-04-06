[![Build Status](https://travis-ci.com/Jsuppers/fasthttp-client.svg?branch=master)](https://travis-ci.com/Jsuppers/fasthttp-client)
[![Coverage Status](https://coveralls.io/repos/github/Jsuppers/fasthttp-client/badge.svg?branch=master&service=github)](https://coveralls.io/github/Jsuppers/fasthttp-client?branch=master)

# fasthttp-client
fasthttp-client is a service which sends one billion json messages to an endpoint

## message format
```
{
    "text": "hello world", 
    "content_id": x, 
    "client_id":y,
    "timestamp": now
}
```
where 
* x is a counter from 1 to 1 billion  
* y is a random number between 1 and 10 
* now is right now with millisecond precision
    
## how to run
```
    git clone https://github.com/Jsuppers/fasthttp-client.git
    docker build -t fasthttp-client .
    docker run --rm -it fasthttp-client
```
This will starting sending messages to http://172.17.0.1:8080, if you want to send to another location, simple set the SERVER_ADDRESS environment variable.
e.g.
```
    docker run --rm -it --env SERVER_ADDRESS=http://localhost:8080 fasthttp-client
```
