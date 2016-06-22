envtool
=======

Example:

Use `envtool` to persist some current environment variables to `/etc/environment` and `/etc/sysconfig/docker`.

```
./envtool persist --edit /etc/environment --edit /etc/sysconfig/docker --select HTTP_PROXY --select http_proxy --select HTTPS_PROXY --select https_proxy --select NO_PROXY --select no_proxy
```
