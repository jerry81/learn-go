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

## troubleshooting

- problem: wrong version, even after using goenv use
  - solution: may need to source zshrc first to apply the shim