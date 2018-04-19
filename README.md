# wheelytask

1. make build_containers // for build docker containers
2. docker-compose up // for up services

```
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
