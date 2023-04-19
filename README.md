# GoNetstat

GoNetstat is a Go-based command-line tool that provides metrics for netstat listening ports on Linux systems. It is designed to run periodically along side the node_exporter.  

## Installation

You can install GoNetstat using `go get`:

```bash
go get github.com/ringsq/gonetstat
```

Alternatively, you can clone the repository and build the binary yourself:

```bash
git clone https://github.com/ringsq/gonetstat.git
cd gonetstat
go build
```

## Usage

To view listening port metrics, simply run the `gonetstat` command:

```bash
./gonetstat
```

This will create the metrics in `/var/lib/node_exporter/netstat_listen.prom` for use with the node_exporter.  A single metric, `netstat_listening_port` is created with labels for the different protocols/ports that are listening.

```
# HELP netstat_listening_port Information about listening ports
# TYPE netstat_listening_port gauge
netstat_listening_port{address="0.0.0.0",port="111",process="rpcbind",protocol="tcp"} 1
netstat_listening_port{address="0.0.0.0",port="111",process="rpcbind",protocol="udp"} 1
netstat_listening_port{address="0.0.0.0",port="22",process="sshd",protocol="tcp"} 1
netstat_listening_port{address="0.0.0.0",port="514",process="rsyslogd",protocol="tcp"} 1
netstat_listening_port{address="0.0.0.0",port="514",process="rsyslogd",protocol="udp"} 1
netstat_listening_port{address="0.0.0.0",port="5432",process="docker-proxy",protocol="tcp"} 1
netstat_listening_port{address="0.0.0.0",port="8080",process="docker-proxy",protocol="tcp"} 1
netstat_listening_port{address="0.0.0.0",port="938",process="rpcbind",protocol="udp"} 1
netstat_listening_port{address="0:1::",port="25",process="master",protocol="tcp6"} 1
netstat_listening_port{address="127.0.0.1",port="25",process="master",protocol="tcp"} 1
netstat_listening_port{address="127.0.0.1",port="3101",process="loki",protocol="tcp"} 1
netstat_listening_port{address="::",port="111",process="rpcbind",protocol="tcp6"} 1
netstat_listening_port{address="::",port="111",process="rpcbind",protocol="udp6"} 1
netstat_listening_port{address="::",port="1514",process="promtail-linux-amd64",protocol="tcp6"} 1
netstat_listening_port{address="::",port="22",process="sshd",protocol="tcp6"} 1
netstat_listening_port{address="::",port="3000",process="grafana",protocol="tcp6"} 1
netstat_listening_port{address="::",port="3100",process="promtail-linux-amd64",protocol="tcp6"} 1
netstat_listening_port{address="::",port="3200",process="promtail-linux-amd64",protocol="tcp6"} 1
netstat_listening_port{address="::",port="3500",process="promtail-linux-amd64",protocol="tcp6"} 1
netstat_listening_port{address="::",port="3600",process="promtail-linux-amd64",protocol="tcp6"} 1
netstat_listening_port{address="::",port="514",process="rsyslogd",protocol="tcp6"} 1
netstat_listening_port{address="::",port="514",process="rsyslogd",protocol="udp6"} 1
netstat_listening_port{address="::",port="5432",process="docker-proxy",protocol="tcp6"} 1
netstat_listening_port{address="::",port="8080",process="docker-proxy",protocol="tcp6"} 1
netstat_listening_port{address="::",port="9080",process="promtail-linux-amd64",protocol="tcp6"} 1
netstat_listening_port{address="::",port="9090",process="prometheus",protocol="tcp6"} 1
netstat_listening_port{address="::",port="9095",process="promtail-linux-amd64",protocol="tcp6"} 1
netstat_listening_port{address="::",port="9100",process="node_exporter",protocol="tcp6"} 1
netstat_listening_port{address="::",port="9115",process="blackbox_exporter",protocol="tcp6"} 1
netstat_listening_port{address="::",port="9180",process="promtail-linux-amd64",protocol="tcp6"} 1
netstat_listening_port{address="::",port="9195",process="promtail-linux-amd64",protocol="tcp6"} 1
netstat_listening_port{address="::",port="9295",process="loki",protocol="tcp6"} 1
netstat_listening_port{address="::",port="938",process="rpcbind",protocol="udp6"} 1
netstat_listening_port{address="::",port="9700",process="sonus_exporter",protocol="tcp6"} 1
```



## License

GoNetstat is licensed under the [MIT License](https://github.com/ringsq/gonetstat/blob/master/LICENSE).