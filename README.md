# go-harvest

go-harvest is a Go client library for accessing the Harvest API

## Disclaimer

This is a fork of an in-progress repository. I'm developing it both to
satisfy my API needs and to learn Go.

![Gopher n00b](http://i.imgur.com/ZGD7g81.gif)

n00b code ahead.

## Usage

```go
import "github.com/bennylope/go-harvest/harvest"
```

Construct a new Harvest client, then use the client to access the Harvest API.

```go
apiClient := harvest.NewAPIClientWithBasicAuth("YOUR_USERNAME", "YOUR_PASSWORD", "YOUR_SUBDOMAIN")

clients := apiClient.Client.List()
client := apiClient.Client.Find(123123)

people := apiClient.People.List()
person := apiClient.People.Find(123123)
```

## License
This library is distributed under the MIT-style license found in the [LICENSE](./LICENSE) file.
