# golang-tips
Golang tips

Some tips for golang programming.


## Goroutine && Channel ##

### goroutine ###
- cluster_service
    > Service cluster have multiple server node.
    > All server provide same services.
    > Send requests to service cluster and get the only one fastest response.

    > In example, we create a cluster with 4 server node. Each node need 1/2/3/4 seconds to handle request.
    > We get the fastest response from cluster.
