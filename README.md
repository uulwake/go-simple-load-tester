# Load Test Simple Server

Create a simple stress test for a server.

Steps:

1. Run server

```bash
go run server/main.go
```

2. Run tester. You can specify number of concurrent users. Default concurrent users is 10 users

```bash
go run tester/main.go
```

specify number of concurrent users

```bash
go run tester/main.go -c 100
```
