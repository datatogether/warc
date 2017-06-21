# Warc
### proper go WARC package

This is a work in progress.


### Building

You'll need [goyacc](https://godoc.org/golang.org/x/tools/cmd/goyacc) to do the whole translate-yacc-file bit, grab that by running:

```shell
go get godoc.org/golang.org/x/tools/cmd/goyacc
go install godoc.org/golang.org/x/tools/cmd/goyacc
```

you'll now have the `goyacc` command, which the makefile will use.