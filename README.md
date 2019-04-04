## Simple KV (key-value) store web service

Some featurs of the appication

- Creat key value pairs
- Get value of a key
- Watch for changes in the key value pairs stored

#### Building the application

##### Quick Setup

```
docker pull liquidslr/storeservice
docker run --rm -p 3000:3000 liquidslr/storeservice
```

##### Build using docker

```
docker build -t  storeservice .
docker run --rm -p 3000:3000 storeservice
```

##### Locally

The command below will create a binary file of the service

```
go build
```

##### Various options

```
Usage:
  store [command]

Available Commands:
  get         Prints value of key stored in db
  help        Help about any command
  put         Create a key value pair
  start       Start the web server, The default port is 3000
  watch       Watch for changes in key value pairs

Flags:
  -h, --help   help for store

Use "store [command] --help" for more information about a command.
```

### Using the cli

##### Starting web server

```
./storeservice start
```

##### To create a new key value pair

```
./storeservice put --key=newkey --value=newvalue
```

##### To get an existing value with a key

```
./storeservice get --key=newkey
```

##### To watch for changes in existing key value pairs

```
./storeservice watch
```

### Using the web server

##### To create a new key value pair

```
curl -d '{"key":"newkey", "value":"newvalue"}' -H "Content-Type: application/json" -X POST http://localhost:3000/api/set/value/
```

##### To get value of a key

```
 curl localhost:3000/api/value/{key}
```

##### To get all key value pairs

```
 curl localhost:3000/api/get/all/
```
