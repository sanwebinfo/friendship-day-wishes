# Friendship Day Wishes

A simple Go application that displays Happy Friendship Day ASCII art and a text greeting with the recipient's name and a social media sharing URL.  

> Just My Fun Project 🙂 and Learning Golang - Happy Friendship Day Wishes 2024  

```yaml

wishes@Your Name:~💚$


 ★─▀██▀▀▀█
 ★──██▄█
 ★──██▀█
 ★─▄██ ANTASTIC Friend ★★★


 Friendship is the compass
 that guides us
 through life's storm

```

> Happy Friendship Day ASCII Text Greeting Art - Friendship Day Greeting Generator With Name.  

## Features

- Friendship Day ASCII art and Text Greeting with name
- Shareable URL for social media sharing
- Supports both HTML and plain text responses
- Proper Error handling and Validations

## Setup

- Clone or Download the Repo

```sh

git clone https://github.com/sanwebinfo/friendship-day-wishes.git
cd friendship-day-wishes

```

- Start the Server

```sh
go run wish.go
```

## Usage

- Send a GET request to the `/wish` endpoint with a name query parameter:

```sh

## Browser View

http://localhost:6054/wish/web/?name=YourName

```

- cURL Request

```sh

curl "http://localhost:6054/wish/text?name=John-Doe"

or

curl -G -d "name=John Doe" http://localhost:6054/wish/text

or

curl -G --data-urlencode "name=John Doe" http://localhost:6054/wish/text

```

- httpie

```sh
http -b GET "http://localhost:6054/wish/text" "name==John Doe"
```

## HTML Response

If the Accept header includes `text/html`, you will get a formatted HTML response.

## Plain Text Response

If the Accept header includes `text/plain`, you will get a plain text response with ANSI color codes.

## Build Package

- Run Make file to build a package for your Systems

```sh
make build
```

## Packges Build for  

Linux, Apple, Windows and Android - `/makefile`  

- Linux-386
- Linux-arm-7
- Linux-amd64
- Linux-arm64
- Andriod-arm64
- windows-386
- windows-amd64
- darwin-amd64
- darwin-arm64

```sh
chmod +x wish
./wish
```

## LICENSE

MIT
