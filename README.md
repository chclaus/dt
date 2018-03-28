[![Build Status](https://travis-ci.org/chclaus/dt.svg?branch=master)](https://travis-ci.org/chclaus/dt)
[![Go Report](https://goreportcard.com/badge/github.com/chclaus/dt)](https://goreportcard.com/report/github.com/chclaus/dt)

# Intro
The dev toolbelt is a fast and flexible tool to handle common use
cases of your daily work. Use cases you generally use some online
tools.

# Table of Contents

- [Example](#example)
- [Commands](#commands)
  * [URI](#uri-command)
  * [Base64](#base64-command)
  * [Hash](#hash-command)
  * [JWT](#jwt-command)
  # [Random](#random-command)
  * [Date](#date-command)
  * [HTML](#html-command)
  * [Server](#server)
- [Contributing](#contributing)
- [License](#license)

# Example
        Usage:
          dt [command]

        Available Commands:
          base64      Encodes or decodes a string to it's base64 representation
          date        Basic date operations
          hash        Hashes an arbitrary input with different hash algorithms
          help        Help about any command
          html        Escapes a html string and vice versa
          jwt         Allows decoding of a jwt
          server      Starts a simple web server to serve static content
          random      Generates random numbers and strings
          uri         Encodes or decodes an URI
          version     Prints the current version of the dt

        Flags:
          -h, --help   help for dt

        Use "dt [command] --help" for more information about a command.

# Commands

## URI command
- Encodes an URI to a safe representation
- Decodes an already encoded URI

## Base64 command
- Encodes a string to it's base64 representation
- Decodes a base64 string to it's plain representation

The base64 command can be used with different encodings:
- Standard, follows RFC 4648
- Standard Raw, follows RFC 4648 but with without padding
- URL, follows the alternative RFC 4648
- URL Raw, follows the alternative RFC 4648 but with without padding

## Hash command
Returns the hash representation of an input in different hash formats:
- md5
- sha1
- sha256
- sha3_256
- sha3_512
- sha512

## JWT command
Decodes a jwt and pretty prints the resulting json.

Further TODOs:
- Signature validation
- JWT creation

## Random command
Generates random strings and numbers. Currently supported functions are:
- Generates a random string, based on an alphabet with a specific length.
- Generates a random string with alphanumeric letters.
- Generates a random string with a alphanumeric and special characters.
- Generates random numbers with a specific length.

## Date command
Date conversions:
- time millis to RFC 3339
- time nanos to RFC 3339
- RFC 3339 to time millis

## Regex command
TODO

## HTML command
Escapes and unescapes HTML
- transforms a string to an escaped html sequence
- reads a file and prints an escaped html sequence
- unescapes an escaped html sequence to it's raw representation


## Password command
Generate passwords and password hashes

TODO implement

## Server
Starts a simple web server to serve static content. You can specify
hostname and port and must set a folder to serve.


# Get dt
## Install
If you've go installed and your `$GOPATH` is set, you can easily install
dt with:

        go get github.com/chclaus/dt

## Build
To build the project you can clone the repository with...

        git clone git@github.com:chclaus/dt

... and build and install the binary with

        cd dt && go install

## Build with Docker
To build the repository with docker, do the following:

        git clone git@github.com:chclaus/dt
        cd dt
        docker build -t dt .
        docker run --rm dt version

# Contributing
Everyone is welcome to create pull requests for this repository. If you're
new to github, take a look [here](https://help.github.com/categories/collaborating-with-issues-and-pull-requests/)
to get an idea of it.

If you've got an idea of a function that should find it's way to this
project, but you won't implement it by yourself, please create a new
issue.

# License
dt (dev-toolbelt) is released under the MIT license. See [LICENSE.txt](LICENSE.txt)
