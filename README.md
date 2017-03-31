# solidfire-gotypes
Collection of Golang Types for interacting with the SolidFire API via Golang

## What are solidfire-gotypes?
These are just helpers for you to use when interacting with the SolidFire API
via golang.  It's not a full set of bindings implementing API calls, but rather
just provides the grunt work of creating Object, Request and Response types.

## Example usage
Common example is the SolidFire ListActiveVolumes API call.  Here's an example
of how you would use these types to issue a request to a SolidFire API to
perform this operation:

```golang
...
    endpoint := "https://adminLogin:password@10.117.36.101/json-rpc/8.0"
    var req gosf.ListVolumesForAccountRequest
    req.AccountID = 75475
    fmt.Printf("Request is: %+v", req)
    vols, err := gosf.ListVolumesForAccount(endpoint, &req, false)
    if err != nil {
        fmt.Printf("awww... bummer dude: %+v", err)
        os.Exit(1)
    }
    fmt.Printf("Volume objects in response: %+v", vols)
...
}
```

See the examples directory for a full example including an example json-rpc
over http implementation.

