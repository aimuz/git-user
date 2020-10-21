# git-user

git multi user management

## Install

Users with go environment

```bash
go get github.com/aimuz/git-user
```

Verify installation

```bash
git user -v
```

## Usage

```bash
git user
```

```
git multi user management

Usage:
 git user [command] [flags]
 git user [command]

Examples:
 git user create --title example --user example --email example@example.com
 git user list
 git user use example

Commands:
 git user clear       Clear current repo username and email configuration
 git user create      create a new git user
 git user help        Help about any command
 git user list        list all git users
 git user use         Switch the current repo git user

Flags:
      --data string    (default "$HOME/.config/git-user/user.yaml")
  -h, --help          help for user
  -v, --version       version for user

Use "user [command] --help" for more information about a command.
```

### Create User

Help:

```bash
$ git user create --help 
Usage:
 git user create [flags]

Aliases:
 create, a

Examples:
 git user create --title example --user example --email example@example.com

Flags:
      --email string   git user email
  -h, --help           help for create
      --title string   if it is empty, username will be used
      --user string    git user name

Global Flags:
      --data string    (default "$HOME/.config/git-user/user.yaml")
```

```bash
$ git user create --title example --user example --email example@example.com
Successfully created example user
```

### List All Users

```bash
$ git user list                                                             
TITLE   USER            EMAIL                   
example example         example@example.com     
```

### Switch Current Repository User

`git user use [title]`

```bash
$ git user use example 
User set successfully

```