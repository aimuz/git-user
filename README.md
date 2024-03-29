# git-user[beta]

> Switching identities on Git is easy, but the trouble is that different repositories may need different identities to commit and pull.

git multi user management

## Install

Users with go environment

```zsh
go install github.com/aimuz/git-user/cmd/git-user
```

Verify installation

```zsh
git user -v
```

## Usage

```zsh
git user
```

```text
git multi user management

Usage:
 git user [command] [flags]
 git user [command]

Examples:
 git user create --title example --user example --email example@example.com
 git user create --title example --user example --email example@example.com -i ~/.ssh/id_rsa
 git user list
 git user use example

Commands:
 git user clean       Clean current repo username and email configuration
 git user create      create a new git user
 git user help        Help about any command
 git user list        list all git users
 git user use         Switch the current repo git user

Flags:
      --data string    (default "/Users/aimuz/.config/git-user/user.yaml")
  -h, --help          help for user
  -v, --version       version for user

Use "user [command] --help" for more information about a command.

```

### Create User

Help:

```zsh
$ git user create --help 
Usage:
 git user create [flags]

Aliases:
 create, a

Examples:
 git user create --title example --user example --email example@example.com
 git user create --title example --user example --email example@example.com -i ~/.ssh/id_rsa

Flags:
      --email string           git user email
  -h, --help                   help for create
  -i, --identity_file string   The certificate corresponding to the user. If it is blank, the default value will be used
      --title string           if it is empty, username will be used
      --user string            git user name

Global Flags:
      --data string    (default "/Users/aimuz/.config/git-user/user.yaml")
```

```zsh
$ git user create --title example --user example --email example@example.com
Successfully created example user
```

### List All Users

```zsh
$ git user list                                                             
TITLE           USER            EMAIL                   IDENTITY FILE   GPG KEY  
example         example         example@example.com     
```

### Switch Current Repository User

`git user use [title]`

```zsh
$ git user use example 
User set successfully
```

## TODO

- [ ] AutoEnable

      Automatically enable the specified account
- [x] GPG support
- [ ] SSH Identity
