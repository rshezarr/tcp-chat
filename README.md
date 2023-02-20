# Net-Cat

## Description

This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

## Usage

1. Command to running tcp server (in root directory)

```
$ go run main.go
```

Server will start at **8989** port

-   Your also able to print your own port

Or, by audit case, you can run server by instruction below

```
$ go build
```

You will have a exe file that you can run through the command

```
$ ./TCPChat $PORT
```

2. Open another terminal, and connect to the chat

```
$ nc localhost $PORT
```

You'll see welcome logo and request for your name\
Next you can chat with your friends

## Authors

@Rshezarr (Rakhat)
