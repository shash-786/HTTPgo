### Http Server

1) ***main.go*** 
    - Using the Gin Framework First the gin.Engine was implemented 
    - The Default gin.Engine is Used  (It comes with Default Logger and Fault   Recovery System)
    - The Routes Were Defined For Creating (POST), Fetching (GET), Updating (POST), and deleting (GET) and then we Began the Server

2) ***server.go***
    - Defined the Server Struct with a simple Hash Map as Database. We don't use a separate db since this Sever _DOES NOT HAVE PERSISTENCE_
    - Defined User Struct with JSON tags for encoding and decoding purposes
    - Create a function `New` to initialize a server and **IMP** Initialize the map
    - Then Define the Handler Functions for the Routes


Inspired by Krishna Iyer [LINK](https://github.com/kicodelibrary/go-http-server-2022/) however this API was implemented using the gin Framework and not the gorilla mux framework

**To Learn More About the [Gin Framework](https://pkg.go.dev/github.com/gin-gonic/gin)**
