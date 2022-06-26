# reference-repo-go
This repo is meant to contain simple go code to be used as a reference for other go devs.

Currently, the implementation is very basic to include an endpoint for integer addition.
The selection of a "calculator" implementation was arbitrary,
and chosen as a simple example use case to demonstrate.


## To run

```bash
make run
```

## To test

```bash
make test
```

### Endpoints

- POST `/api/v1/add`: returns the total from the addition of the int slice payload

```bash
curl -X POST localhost:8080/api/v1/add -i -d '{"numbers":[1,2,3]}'
```
