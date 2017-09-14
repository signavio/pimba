# Pimba!

Pimba is a small and simple CLI tool to easily publish and serve static files.

## Install

[Binary releases](https://github.com/signavio/pimba/releases/) are available.
Please, download the suitable binary for your system.

If you have Go installed, you may also build and install it using `go get`:

```
$ go get [-u] github.com/signavio/pimba
```

## Serving

To serve the Pimba API and the static files, execute:

```
$ pimba api --storage /path/to/data/storage
```

It's also possible to set the port, passing the option `--port <port-number>`.

For further help on how to serve files, execute `pimba help api`.

## Pushing

To push files to the Pimba server, first enter the directory that you
would like to publish and execute:

```
$ cd /path/to/publish
$ pimba push --server pimba.server.host:port
```

For further help on how to push, execute `pimba help push`.

## Configuration

It's also possible to have a configuration file for Pimba. Refer to
[pimba.yaml.sample](pimba.yaml.sample) and create a config file `.pimba.yaml`
in your home directory. The configuration file will set defaults
based on your preferences.

```
$ cp pimba.yaml.sample $HOME/.pimba.yaml
$ vim $HOME/.pimba.yaml #edit as you like
```

## Maintainers

Stephano Zanzin - [@microwaves](https://github.com/microwaves)

## License

Please, refer to the [LICENSE](LICENSE) file.
