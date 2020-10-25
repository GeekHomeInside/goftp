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
➜ ./server
Hello I'm GoFTP Server
```

- launch client

```sh
➜ ./client 
Hello I'm goClient
Receive from:  127.0.0.1:3333 Hello Caller!, I'm GoFTP Server :)
```

- Now look Server side to see the connexion between server and client

```sh
➜ ./server
Hello I'm GoFTP Server
Connexion OK

recived from:
 127.0.0.1:43106
```