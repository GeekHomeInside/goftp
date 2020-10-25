# goftp

This repository is used to learn how to code in golang.

- [RFC 999](https://tools.ietf.org/html/rfc959)

I try to provide:

- One Binaries for Server
- One Binaries for Client


# How to build package? 

```sh
make build
```

# How to test?

- launch the server

```sh
➜ ./server 3333
Hello I'm GoFTP Server
Ready to handle connexion
```

- launch client

```sh
➜ ./client 127.0.0.1:3333
Hello I'm goClient
>> 
```

- Now look Server side to see the connexion between server and client

```sh client
>> test
->: goFTP Server received: test
```

```sh server
-> test
```

# How to stop the server

```sh
>> STOP
```

or `CTRL+C`