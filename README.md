# rtail
Remote tail, ie tail -f over http.

# Default addr is localhost:8000
rtail -f /var/log/messages

rtail -f -addr=0.0.0.0:9000 /var/log/messages

journalctl -f | rtail -f  
