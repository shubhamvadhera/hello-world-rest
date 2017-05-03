# "Hello World" REST API
### Simple “Hello World” REST API in Go to understand how handler and Http router work, and how to parse the request and send the response.

1. GET /hello/xxxx

Response:
Hello, xxxx!


2. POST /hello

Request:
{
   “name”: “foo”
}

Response:
{
   “greeting” : “Hello, foo!”
}
