# udp
Simple UDP test client and server

Build the docker server container:
```
docker build -t server .
```

Run the docker server container:
```
docker run -p 8888:8888/udp server
```

Build a local client:
```
cd cmd/client
go build
```

Run parallel clients with different payloads:
```
./client -payload "foo"
```
and
```
./client -payload "bar"
```

