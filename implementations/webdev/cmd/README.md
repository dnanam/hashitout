## HTTP Request methods

https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Methods

* GET
The GET method requests a representation of the specified resource. Requests using GET should only retrieve data and should not contain a request content.

* HEAD
The HEAD method asks for a response identical to a GET request, but without a response body.

* POST
The POST method submits an entity to the specified resource, often causing a change in state or side effects on the server.

* PUT
The PUT method replaces all current representations of the target resource with the request content.

* DELETE
The DELETE method deletes the specified resource.

* CONNECT
The CONNECT method establishes a tunnel to the server identified by the target resource.

* OPTIONS
The OPTIONS method describes the communication options for the target resource.

* TRACE
The TRACE method performs a message loop-back test along the path to the target resource.

* PATCH
The PATCH method applies partial modifications to a resource.


## Status Codes

Set an explicit status code, if you do not set on Go defaults to 200.

## Handler Function

## A function that helps to reload server when file system has a small change
 
## Content Type Headers
Headers are metadata that can be put on a request, they provide additional information about a web request.
The content type header tells what type of content we are sending, web/image/json/html/etc.

## Adding a Contact Page to the application

## Different ways to route the pages using ServeMux
Routing is what happens when we decide which page to show the user when a request comes to the server.

## Custom Routing
Since we can get the URL from the request, we can do custom routing using the request 

## URL Path vs RAW Path
Anything after ? will be additional data we are passing to the server. The path ends at ?.
Its also possible that what we put in query parameters can be in the path.
For example : localhost:3000/page?id=123 or localhost:3000/page/123
The RAWPath attribute will treat all encoded values that were part of the url path as encoded
and give us the encoded value.
The URL Path on the other hand actually handles decoding for us.

## Handler Type

### What is a handler
n the context of web requests, a handler is a piece of code that processes a request and generates a response. 
It's the component that handles the logic of a specific URL or resource, determining how to respond to a client's request.
The client request for example is show me localhost:9100/contact page

## Using Chi as the default router
If you were to write your own router, which is basically what looks at path and decides what do with the request,
there are chances that it will take quite a bit of time to implement the one which can handle all use cases and hence using a standard library like chi is helpful.