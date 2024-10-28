# CodeGraph

### Usage

Compile binary:

```
task --list-all
task compile
```

Example usage:

- find all files in example folder
- pass them all codegraph binary

```
find ./examples/ -type f -name "*.js" | xargs ./codegraph | jq
```
