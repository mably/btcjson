btcjson
=======

[![Build Status](https://travis-ci.org/mably/btcjson.png?branch=master)]
(https://travis-ci.org/mably/btcjson)
[![tip for next commit](http://peer4commit.com/projects/130.svg)](http://peer4commit.com/projects/130)

Package btcjson implements the bitcoin JSON-RPC API.  There is a test
suite which is aiming to reach 100% code coverage.  See
`test_coverage.txt` for the current coverage (using gocov).  On a
UNIX-like OS, the script `cov_report.sh` can be used to generate the
report.  Package btcjson is licensed under the liberal ISC license.

This package is one of the core packages from btcd, an alternative full-node
implementation of bitcoin which is under active development by Conformal.
Although it was primarily written for btcd, this package has intentionally been
designed so it can be used as a standalone package for any projects needing to
communicate with a bitcoin client using the json rpc interface.
[BlockSafari](http://blocksafari.com) is one such program that uses
btcjson to communicate with btcd (or bitcoind to help test btcd).

## JSON RPC

Bitcoin provides an extensive API call list to control bitcoind or
bitcoin-qt through json-rpc.  These can be used to get information
from the client or to cause the client to perform some action.

The general form of the commands are:

```JSON
	{"jsonrpc": "1.0", "id":"test", "method": "getinfo", "params": []}
```

btcjson provides code to easily create these commands from go (as some
of the commands can be fairly complex), to send the commands to a
running bitcoin rpc server, and to handle the replies (putting them in
useful Go data structures).

## Sample Use

```Go
	msg, err := btcjson.CreateMessage("getinfo")
	reply, err := btcjson.RpcCommand(user, password, server, msg)
```

## Documentation

Full `go doc` style documentation for the project can be viewed online without
installing this package by using the GoDoc site
[here](http://godoc.org/github.com/mably/btcjson).

You can also view the documentation locally once the package is installed with
the `godoc` tool by running `godoc -http=":6060"` and pointing your browser to
http://localhost:6060/pkg/github.com/mably/btcjson

## Installation

```bash
$ go get github.com/mably/btcjson
```

## TODO

- Increase test coverage to 100%.

## License

Package btcjson is licensed under the liberal ISC License.
