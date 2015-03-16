# ip

The ip package has a utility that allows the requestor to get a full IP Address when passed an incoming HTTP Request object.

### Usage

After importing `github.com/kyani-inc/go-utils/ip` the usage is as easy as:

```
import (
    "net/http"
    "github.com/kyani-inc/go-utils/ip"
)

func main() {
    req, _ := http.NewRequest("GET", "http://example.com", nil)
    req.Header.Add("X-Forwarded-For", "127.0.0.1")

    // Returns "127.0.0.1"
    addr := ip.Client(req)
}
```

### Todo

- ~~Add Tests~~
- ~~Verify IPv6 Support (a.k.a. Tests)~~
