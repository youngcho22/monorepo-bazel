## query
# query targets available in a target..?
```
bazel query @npm//...

bazel query @npm//... | grep express // query express target in npm
```

# query depedency of a target
```
bazel query 'deps(//projects/calculator-api)' --noimplicit_deps
```

# query reverse dependenc of a target
```
bazel query 'rdeps(//..., //projects/calculator)'
```

# query with union
```
// query all projects that depend on calculator OR go_hello_world
bazel query 'rdeps(//..., (//projects/calculator union //projects/go_hello_world))'
```

## Tags
Bazel targets can be tagged with different IDs. With tags, targets can be filtered during bazel commands
```
bazel query 'attr(tags, "calc", //projects/calculator/...)'

bazel test --test_tag_filters calc
```

## Gazelle
```
go mod init ...
go mod tidy
bazel run //:gazelle-update-repos // this update dpes.bzl to include all deps in go.mod
```

## Java
```
need to run this after copy and pasting setup code int WORKSPACE.bazel
bazel run @maven//:pin
```