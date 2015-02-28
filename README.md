# Git My Key!

Simple util to fetch a users public key from github.com's API.

### Install

#### Via `go get`

```
go get -u github.com/jmervine/git-my-key
```

> Note: ensure `GOBIN` environment variable is set.
>
> Something like:
>
> `$ test "$GOBIN" || (mkdir ~/.gobin && export GOBIN=~/.gobin)`

#### Download

Currently I only have the `linux/x86_64` binaries ready.

```
$ curl -sS -O http://static.mervine.net/go/linux/x86_64/git-my-key && chmod 755 git-my-key
$ ./git-my-hey help
```

### Use

```
NAME:
   ./git-my-key - fetch users public key from github.com

USAGE:
   ./git-my-key [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
  Joshua Mervine <joshua@mervine.net> - <unknown@email>

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --username, -u 	target user's github username
   --output, -o 	public key output file, default: {username}.pub
   --append, -a		append key to output file
   --help, -h		show help
   --version, -v	print the version

```
