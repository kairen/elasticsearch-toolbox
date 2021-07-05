# Elasticsearch Toolbox
A tool to maintain an Elasticsearch cluster.

* [x] Rotate indices.
* [x] Backup indices(repository and snapshot).
* [ ] Restore indices.

## Building from Source
Clone repo into your go path under `$GOPATH/src`:

```sh
$ git clone https://github.com/kairen/elasticsearch-toolbox.git $GOPATH/src/github.com/kairen/elasticsearch-toolbox
$ cd $GOPATH/src/github.com/kairen/elasticsearch-toolbox
$ make
```

## Usage

```sh 
$ ./out/elasticsearch-toolbox
Maintaining and operating an Elasticsearch cluster.

Usage:
  elasticsearch-toolbox [command]

Available Commands:
  help        Help about any command
  repository  Create a repository to place snapshots
  rotate      Delete expired indices
  snapshot    Create a snapshot for backuping indices
  version     Print the version information

Flags:
      --alsologtostderr                  log to standard error as well as files
      --endpoints strings                Endpoints of elasticsearch. (default [http://elasticsearch:9200])
  -h, --help                             help for elasticsearch-toolbox
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --password string                  Password for basic auth.
      --retry-count int                  The number of retry for deleting request. (default 5)
      --sniffer                          Enable client to use a sniffing process for finding all nodes of your cluster.
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
      --tls.ca string                    SSL Certificate Authority file used to secure elasticsearch communication.
      --tls.cert string                  SSL certification file used to secure elasticsearch communication.
      --tls.key string                   SSL key file used to secure elasticsearch communication.
      --tls.skip-host-verify             (insecure) Skip server's certificate chain and host name verification
      --username string                  Username for basic auth.
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging

Use "elasticsearch-toolbox [command] --help" for more information about a command.
```
