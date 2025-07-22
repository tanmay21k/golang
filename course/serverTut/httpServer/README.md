We use the net/http module for the purpose of creating the HTTP serer in go. 



## Handlers
Handlers are basically the functions that are used for performing different functionalities at different routes

Handlers are interfaces created to perform operations.
Handlers take parameters as ResponseWriter and *Request

## ListenAndServe

The method is present inside the net/http module and is used for creating a basic HTTP server in Golang.

The method takes the input as portnumber to be hosted upon and the handler, often times the function will look like 

_ , err := http.ListenAndServe(":2000", nil)
if err != nil {
    log.Fatal(err)
}