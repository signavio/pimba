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
$ pimba api --storage /path/to/data/storage --secret my-jwt-secret
```

The flag `--secret` is mandatory and it's the necessary key for signing
tokens for pushing to the Pimba buckets.

It's also possible to set the port, passing the flag `--port <port-number>`.

For further help, execute `pimba help api`.

## Pushing

To push files to the Pimba server, first enter the directory that you
would like to publish and execute:

```
$ cd /path/to/publish
$ pimba push --server pimba.server.host:port --name my-bucket-name
```

If the flag `--name` is not passed, Pimba will create a bucket with a random
string as the name.

After you did the first push to your bucket, use the returned token to be able
to update the bucket. Execute:

```
$ pimba push --server pimba.server.host:port --name my-bucket-name --token returned-token
```

Remember to save your token in a safe place, Pimba **doesn't store tokens**,
thus meaning that if you lose the token the bucket will become inaccessible.

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
