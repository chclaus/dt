# Intro
The current project state is a very early draft...

# Table of Contents

- [Overview](#overview)
- [Commands](#commands)
  * [URI](#uri-command)
- [Contributing](#contributing)
- [License](#license)

# Overview
The dev toolbelt is a fast and flexible tool to handle common use
cases like encodings, string, jwt, ... operations.

Use cases you generally use some online tools - BAH!

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

# Contributing
tbd.

# License
dt (dev-toolbelt) is released under the MIT license. See LICENSE.txt
