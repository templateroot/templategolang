# reomte folder ~/.local/...
# provider: account id: xxx mail: xx@xx;


CGO_ENABLED=0 GOOS=freebsd  go build  -ldflags "-w -s" -buildvcs=false
scp binName username@hosturl:~/
rm binName