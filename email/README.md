# Email

[tutorials](../README.md) / email

Commands:

```sh
$ go run .
```

Protocols:

- IMAP
- POP3
- SMTP

## Outgoing messages

**SMTP** - used to send messages

```
Type  Security             Auth          Port
----  -------------------  ------------  -----
SMTP  no encryption        AUTH          587
SMTP  encrypted (TLS)      StartTLS      587
SMTP  encrypted (SSL)      SSL           465
```

## Incoming messages

**IMAP** - used to view and manipulate message on server

```
Type  Security             Auth          Port
----  -------------------  ------------  -----
IMAP  no encryption        AUTH          143
IMAP  encrypted (TLS)      StartTLS      143
IMAP  encrypted (SSL)      SSL           993
```

**POP3** - used to read and download messages to your local device

```
Type  Security             Auth          Port
----  -------------------  ------------  -----
POP3  no encryption        AUTH          110
POP3  encrypted (TLS)      StartTLS      110
POP3  encrypted (SSL)      SSL           995
```
