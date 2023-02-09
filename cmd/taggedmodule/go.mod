module github.com/ericklaus-wf/go-multiple-module-example/cmd/taggedmodule

go 1.18

replace github.com/ericklaus-wf/go-multiple-module-example => ../../

// remember to update this version if you begin depending on additional features from the root
require github.com/ericklaus-wf/go-multiple-module-example v1.0.0
