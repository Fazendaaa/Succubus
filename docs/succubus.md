# Succubus

## Configuration Description

### Top view

```yaml
[optional] image:
[optional] dockerfile:

[required] base:
...

[required] dev:
...

[optional] extended:
...
```

### Base

```yaml
...
[required] base:
[required]     run: ...
[required]     add: ...
[required]     test: ...
```

### Development

```yaml
...
[required] dev:
[required]     anal: ...
[required]     linter: ...
[required]     doc: ...
```

### Extended

```yaml
...
[optional] extended:
[optional]     <anyname>: ...
...
```

## Dockerfile

<!-- Explain:
1. Layers
2. Caching
3. Multi-stage
-->

### Default

### Multi-staged

<!-- Explain:
1. System tag
2. Runner tag
3. Everything in between
-->

#### System

#### Runner
