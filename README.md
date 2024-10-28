# CodeGraph

### Usage

Compile binary:

```
task --list-all
task compile
```

Example usage:

- find all files in example folder
- pass them all to codegraph binary

```
find ./examples/ -type f -name "*.js" | xargs ./codegraph | jq
```

Will output as:

```
[
  {
    "type": "label",
    "value": "handler-consumer-execute-a",
    "filepath": "./examples/javascript/service-handler-consumer.js",
    "position": {
      "row": 2,
      "column": 3
    }
  }, ...
]
```
