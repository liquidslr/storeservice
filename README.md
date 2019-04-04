## Simple KV (key-value) store web service

#### Using the web server

##### To create a new key value pair

```
curl -d '{"key":"newkey", "value":"newvalue"}' -H "Content-Type: application/json" -X POST http://localhost:3000/api/set/value/
```

##### To get value of a key

```
 curl localhost:3000/api/value/{key}
```

#### Building the application

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

Flags:
  -h, --help   help for store

Use "app [command] --help" for more information about a command.
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
