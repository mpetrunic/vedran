# Vedran

> Polkadot chain load balancer.

### Architecture

_Vedran loadbalancer_ is used in conjunction with [Vedran daemon](https://github.com/NodeFactoryIo/vedran-daemon). Suppose the node owner wants to register to loadbalancer, than it is required to install and run _Vedran daemon_. Daemon executes the registration process and starts providing all relevant information (ping, metrics) to the _Vedran loadbalancer_. Please check [Vedran daemon repo](https://github.com/NodeFactoryIo/vedran-daemon) for more details on the daemon itself.


![Image of vedran architecture](./assets/vedran-arch.png)

### Get `vedran` package
1. Install [Golang](https://golang.org/doc/install) **1.13 or greater**
2. Run the command below
```
go get github.com/NodeFactoryIo/vedran
```
3. Run vedran from your Go bin directory. For linux systems it will likely be:
```
~/go/bin/vedran
```
Note that if you need to do this, you probably want to add your Go bin directory to your $PATH to make things easier!

## Usage

Load balancer is started using command `start`.

```
$ vedran -h

vedran is a command line interface for polkadot load balancer

Usage:
  vedran [command]

Available Commands:
  help     Help about any command
  start    Starts vedran load balancer

Use "vedran [command] --help" for more information about a command.
```

When running vedran load balancer, flags described below can be used to customize instance. 

```
--auth-secret string   [REQUIRED] Authentication secret used for generating tokens
--capacity int         [OPTIONAL] Maximum number of nodes allowed to connect, where -1 represents no upper limit (default -1)
--fee float32          [OPTIONAL] Value between 0-1 representing fee percentage (default 0.1)
--log-file string      [OPTIONAL] Path to logfile (default stdout)
--log-level string     [OPTIONAL] Level of logging (eg. info, warn, error) (default "error")
--name string          [OPTIONAL] Public name for load balancer, autogenerated name used if omitted (default "load-balancer-wAbqMEaavbwy")
--port int32           [OPTIONAL] Port on which load balancer will be started (default 4000)
--selection string     [OPTIONAL] Type of selection used for choosing nodes (default "round-robin")
--whitelist strings    [OPTIONAL] Comma separated list of node id-s, if provided only these nodes will be allowed to connect
```

### Demo

#### Requirements

- Install [Docker Engine](https://docs.docker.com/engine/install/)
- Install [Docker Compose](https://docs.docker.com/compose/install/)

**Run demo with `docker-compose up`**

This demo starts three separate dockerized components:
- _Polkadot node_
- _Vedran daemon_
- _Vedran loadbalancer_ 

## License

This project is licensed under Apache 2.0:
- Apache License, Version 2.0, ([LICENSE-APACHE](http://www.apache.org/licenses/LICENSE-2.0))