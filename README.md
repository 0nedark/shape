## Shape
A simple command that allows you to modify .json files with the provided key value arguments. Usefull if you want to modify config files in command line or with bash scripts.

## Instalation
run `go get github.com/0nedark/shape`

## Flags
`--file value`, `-f value` - specifies the file to be modified

`--mutable`, `-m` - allow adding of new keys to the file

## Examples
**Initial JSON**
```json
{"loglevel":"info","env":"stg"}
```
**update single key**
`shape -f ./config.json loglevel=debug`
```json
{"loglevel":"debug","env":"stg"}
```

**update multiple keys**
`shape -f ./config.json loglevel=debug env=dev`
```json
{"loglevel":"debug","env":"dev"}
```

**add new key**
`shape -m -f ./config.json example=a`
```json
{"loglevel":"info","env":"stg","example":"a"}
```
