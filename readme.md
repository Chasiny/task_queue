# Simple Task Queue

## Start

* go run main.go

## Add Task
```
curl -X POST localhost:8103 -d '{"id":"print-tim","cmd":"date","args":["-R"],"interval":5000}'
```

respon
```json
{"ok":true,"error":"","id":"print-time","cmd":"","args":null,"interval":0}
```

## Delete Task
```
curl -X DELETE localhost:8103 -d ' {"id":"print-time"}'
```

respon
```json
{"ok":true,"error":"","id":"print-time","cmd":"","args":null,"interval":0}
```
