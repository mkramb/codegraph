# CodeGraph

### Functionality

The concept is to integrate markers within your codebase to enable the generation of a code graph that maps services and dependencies. This approach can be applied across various programming languages and infrastructure as code. It provides a service map without relying on telemetry, allowing teams to visually understand the actual workings of the code. With this insight, you can deliver new features and implement refactors faster and with fewer bugs.

### Usage

Example usage:

- find all files in example folder
- pass them all to codegraph binary
- lastly, load visualizer in the browser

```
task examples-generate
task examples-visualizer
```

Which will generate a file as:

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

And visiting the visualizer app on http://localhost:3000/, would render a code graph:
![plot](./assets/visualizer.png)

## Development setup

Install & build:

```
task setup-pnpm
task compile
```

To list all tasks:

```
task --list-all
```
