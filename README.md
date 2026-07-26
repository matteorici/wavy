# Drone Battery Estimator

Drone Battery Estimator is a Go console application that estimates the remaining flight time of a drone using several environmental parameters.

The estimator considers battery charge, ambient temperature and payload weight to produce an approximate flight duration.

## Features

- Battery percentage input
- Temperature adjustment
- Payload correction
- Flight time estimation
- Console summary
- Report generation

## Example

Input

Charge: 82
Temperature: 18
Payload: 0.8

Output

Estimated Flight Time: 24.3 minutes

## Run

```bash
go run .
```
