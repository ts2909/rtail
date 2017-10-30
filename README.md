# rtail
```
NAME:
   rtail - serve a file, or stdin to http

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -f, --follow   follow a file
   -V, --verbose  enable verbose output
   --addr value   listen addr ie localhost:8000 (default: "localhost:8000")
   --help, -h     show help
   --version, -v  print the version
```
### Examples   
```
rtail -f /var/log/messages

rtail -f -addr=0.0.0.0:9000 /var/log/messages

journalctl -f | rtail -f  
```