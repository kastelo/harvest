# go-harvest

go-harvest is a Go client library for accessing the [Harvest API](https://github.com/harvesthq/api).

## Usage

```go
import "github.com/PeterVerschuure/go-harvest"
```

Construct a new Harvest client, then use the client to access the Harvest API.

```go
apiClient := harvest.NewAPIClientWithBasicAuth(
	"YOUR_USERNAME", "YOUR_PASSWORD", "YOUR_SUBDOMAIN")

clients := apiClient.Client.List()
client := apiClient.Client.Find(123123)

people := apiClient.People.List()
person := apiClient.People.Find(123123)
```

## License
This library is distributed under the MIT-style license found in the [LICENSE](./LICENSE) file.
