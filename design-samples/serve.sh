#!/bin/sh
# Serves design-samples/ (and only it) to the home network for the 1.8 sample-screen gate
# (08 §19.8, §25.11). Run by the lead during the gate session only.
#
# Usage: design-samples/serve.sh <private IPv4 of this PC> <port>
#
# The address is always explicit: never 0.0.0.0, never a public address, so the samples are
# not offered on every interface by accident. http.server drops ".." segments and serves
# nothing outside --directory; the symlinks here point only at web/static and tokens.css.
set -eu

usage() {
  echo "usage: $0 <private IPv4 of this PC> <port 1024-65535>" >&2
  exit 2
}

[ $# -eq 2 ] || usage
ip=$1
port=$2

case $ip in
  127.* | 10.* | 192.168.* | 172.1[6-9].* | 172.2[0-9].* | 172.3[01].*) ;;
  *)
    echo "refusing $ip: give this PC's private LAN IPv4 address (not 0.0.0.0, not public)" >&2
    exit 2
    ;;
esac

case $port in
  '' | *[!0-9]*) usage ;;
esac
[ "$port" -ge 1024 ] && [ "$port" -le 65535 ] || usage

dir=$(cd "$(dirname "$0")" && pwd)
echo "Serving $dir at http://$ip:$port/ (Ctrl+C to stop)" >&2
exec python3 -m http.server "$port" --bind "$ip" --directory "$dir"
