# Badger Embedded Database

This is a simple demo of the [Badger](https://github.com/dgraph-io/badger) Key Value Store. It provides a robust pure
Go embedded LSM key store with separately stored values. It is equal (on SSD) or barely slower (on HDD) than Bolt for 
reads and much faster for writes because it uses a write-ahead log (WAL). This also means a power loss could prevent 
the database from flushing all writes to disk upon shutdown.

## Build and Run

```
go build
.\badger.exe
```

[&#x2190; Back to Project List](../README.md)