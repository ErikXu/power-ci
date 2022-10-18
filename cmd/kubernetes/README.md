# Kubernetes tools

## Usage

- Install kubernetes

``` bash
# Single ip
./power-ci k8s install --masters x.x.x.x --nodes x.x.x.x -p {password of all machines}

# Ip array
./power-ci k8s install --masters x.x.x.x,x.x.x.x --nodes x.x.x.x,x.x.x.x -p {password of all machines}

# Ip range
./power-ci k8s install --masters x.x.x.x-x.x.x.x --nodes x.x.x.x-x.x.x.x -p {password of all machines}
```
