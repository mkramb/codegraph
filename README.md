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
task generate-examples
```

Will generate file as:

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
