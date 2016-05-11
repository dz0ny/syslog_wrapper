# syslog_wrapper

[![wercker status](https://app.wercker.com/status/1214b503d63d0bf2178dda7373d98983/m "wercker status")](https://app.wercker.com/project/bykey/1214b503d63d0bf2178dda7373d98983)

You can find precompiled binaries under *Releases*.

```
Usage of ./syslog_wrapper-linux-amd64:
  -cmd string
    	Command to wrap
  -syslog string
    	Address of internal syslog UDP server (default "127.0.0.1:514")
```

## Example
```
# ./syslog_wrapper-linux-amd64 -cmd "haproxy -D -f haproxy.cfg" -syslog "127.0.0.1:5555"
2016/05/11 21:50:49 Syslog server started on UDP: 127.0.0.1:5555
2016/05/11 21:50:49 Starting command: haproxy -D -f haproxy.cfg
[WARNING] 131/215049 (6564) : config : missing timeouts for frontend 'web'.
   | While not properly invalid, you will certainly encounter various problems
   | with such a configuration. To fix this, please ensure that all following
   | timeouts are set to a non-zero value: 'client', 'connect', 'server'.
[WARNING] 131/215049 (6564) : config : missing timeouts for backend 'apache_app'.
   | While not properly invalid, you will certainly encounter various problems
   | with such a configuration. To fix this, please ensure that all following
   | timeouts are set to a non-zero value: 'client', 'connect', 'server'.
2016/05/11 21:50:49 Proxy web started.
2016/05/11 21:50:49 Proxy apache_app started.
2016/05/11 21:51:02 127.0.0.1	GET	404	/	500	867	{localhost:8081||sl,en-US;q=0.8,en;q=0.6,pl;q=0.4||Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.75 Safari/537.36|}

```
