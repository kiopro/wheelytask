# wheelytask

1. make build_containers // build docker containers
2. docker-compose up // up services

API: GET request for price with example JSON payload

```shell
curl --request GET \
  --url 'http://localhost:4567/get_price?=' \
  --header 'content-type: application/json' \
  --data '{
	"start_point" : {
		"lat" : 55.737142,
		"long": 37.532897
	},
	"end_point" : {
		"lat" : 55.760672,
		"long": 37.626933
	}
}'
```
