# Path Of Exile server status


![Build](https://github.com/Gonzih/poe-status.com/workflows/CI/badge.svg)

Path Of Exile server scanner, written in go, based on ping and nmap data for server status, centralized view/storage server, UI built in ClojureScript.
Communications between scanner workers and server are based on Twirp protocol.

https://poe-status.com.

Released under MIT License.


## Running a scanner

Download latest client from releases and run it by giving it correct token and pointing it to server:

```bash
env MAIN_TOKEN=currenttoken ./poe-status-client --url https://poe-status.com
```

It is preferable to run your client from a reliable network to eliminate local packet loss.
