# Fuux - simplest storage server

##### How to run
```
go run fuux/cmd/server -conf config.json
```

##### Config sample
```json
{
  "resource": [
    {
      "name": "zzz",
      "upload": {
        "secret": "xxx"
      },
      "download": {
        "secret": "aaa"
      }
    }
  ]
}
```