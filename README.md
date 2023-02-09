# Multiple module example

This repo attempts to show various outcomes of multiple modules. There is a root module with one package: modexample. 
There are three binaries produced by this repo, all located in packages under `cmd`. 
Two of these (taggedmodule and untaggedmodule) declare their own `go.mod` and are therefore their own modules.
Another (nomodule) does not declare any module.

The `consuming` module then imports all the packages in this repo. Note the following:
- `nomodule` does not appear in the go.mod file. It does not have its own go.mod and is thus part of the `go-multiple-module-example` module.
- `untaggedmodule` appears in the go.mod, but without a version. It appears in the go.mod because it is its own module (it has a go.mod file). It lacks a version because there is no tag in the format `cmd/untaggedmodule/vX.Y.Z`.
- `taggedmodule` appears in the go.mod and has a version. It appears in the go.mod because it is its own module (it has a go.mod file). It has a version because there is a tag in the form `cmd/taggedmodule/vX.Y.Z` (in this case `cmd/taggedmodule/v1.0.1`).
