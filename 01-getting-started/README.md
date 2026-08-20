# getting started topics

## install go

- can install with homebrew
```
brew install go
```
- found it is downloaded and put into correct location /usr/local/go/bin
- but its not on the path (in zshrc)
- however, /etc/paths.d/go created by go installer handles pathing

- also can use goenv (like rvm)
```
brew install goenv
# add these to .zshrc
export GOENV_ROOT="$HOME/.goenv"
export PATH="$GOENV_ROOT/bin:$PATH"
eval "$(goenv init -)"

goenv install 1.25.x + goenv use 1.25.x
# use activates, install downloads
```

## build and run programs
1.  go build <filename>
2.  ./<filename>

or

1.  go run <filename>

## manage packages

- package manager built into go itself
- go.mod, go.sum
```
go get github.com/gin-gonic/gin@v1.10.0   # add/upgrade a dep (edits go.mod)
go mod tidy                                # sync go.mod/go.sum with actual imports (run this often)
go mod download                            # fetch deps into module cache
go mod verify                              # check checksums
go list -m all
```

1.  write code that has import
2.  go mod init <module name> (creates go.mod)
3.  go mod tidy (creates go.sum and downloads deps)

## troubleshooting

- problem: wrong version, even after using goenv use
  - solution: may need to source zshrc first to apply the shim