# 简易的定时任务队列

## 启动

* go run main.go

## 添加任务
```
curl -X POST localhost:8103 -d '{"id":"print-tim","cmd":"date","args":["-R"],"interval":5000}'
```

respon
```json
{"ok":true,"error":"","id":"print-time","cmd":"","args":null,"interval":0}
```

## 删除任务
```
curl -X DELETE localhost:8103 -d ' {"id":"print-time"}'
```

respon
```json
{"ok":true,"error":"","id":"print-time","cmd":"","args":null,"interval":0}
```