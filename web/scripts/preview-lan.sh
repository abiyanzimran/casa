#!/usr/bin/env bash
# Preview of the web client on the pilot devices over the home network (plan
# 1.9b/1.10, owner decision Q5): fullstack-server.sh behind Caddy's TLS, because
# the service worker, crypto.subtle, Web Locks, the __Host- cookie and
# persist() need a trusted HTTPS origin.
#
#   web/scripts/preview-lan.sh <IP-LAN>        e.g. 192.168.1.37
#
# Serves https://<IP>:$CASA_PREVIEW_HTTPS_PORT (default 18443; Caddy, its own
# local CA, proxying to casa on 127.0.0.1:$CASA_PREVIEW_BACKEND_PORT, default
# 18081) and, only for installing the CA on the phones, the PUBLIC root
# certificate at http://<IP>:$CASA_PREVIEW_CERT_PORT/root.crt (default 18765).
# The ports are not 8080/8443/8765: other projects hold 8080, 8443 and Caddy's
# admin port 2019 on this PC, so the Caddyfile turns the admin API off.
# The CA's private key stays in ~/.local/share/casa-preview-caddy (mode 700) on this PC. Remove the
# certificate from every device when the preview is over (Q5: when WS-5 exists).
# Origin is passed through unchanged; nothing but the app is proxied.
#
# Caddy is a user-local binary in .cache/bin/caddy, pinned (ES §8):
#   v2.11.9, released 2026-11-05 (9 days old when pinned on 2026-11-14), the
#   current stable release; SHA-512 of the tarball verified against the release's caddy_2.11.9_checksums.txt.
#   Source: https://github.com/caddyserver/caddy/releases/tag/v2.11.9
#   Advisories: the five published 2026-11-04 (GHSA-jf74-f69g-hcj9,
#   GHSA-r766-rf9f-588h, GHSA-25cw-38mq-q7m6, GHSA-69hv-5jfc-mvmv,
#   GHSA-f278-mgc7-46j3) all say "<= 2.11.7, fixed in 2.11.8"; none applies.
set -euo pipefail
umask 077

caddy_version=v2.11.9
caddy_sha256=5b8cb7faf8726a2028ad96a66e54e6e81e8f361bc6175e5d22bbd997cc5177fb # SHA-256 of caddy_2.11.9_linux_amd64.tar.gz; the release lists SHA-512 (0e855443...edd75), checked equal when pinned

https_port=${CASA_PREVIEW_HTTPS_PORT:-18443}
backend_port=${CASA_PREVIEW_BACKEND_PORT:-18081}
cert_port=${CASA_PREVIEW_CERT_PORT:-18765}
ip=${1:?usage: preview-lan.sh <IP-LAN>}
[[ $ip =~ ^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$ ]] || { echo "preview-lan: <IP-LAN> must be an IPv4 address" >&2; exit 2; }
root=$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)
caddy=$root/.cache/bin/caddy
data=${CASA_PREVIEW_CADDY_DATA:-$HOME/.local/share/casa-preview-caddy} # off the repository: its mount may not keep file modes

if [ "$("$caddy" version 2>/dev/null | cut -d' ' -f1)" != "$caddy_version" ]; then
	[ "$(uname -sm)" = "Linux x86_64" ] || { echo "preview-lan: no pinned checksum for $(uname -sm)" >&2; exit 1; }
	dl=$(mktemp -d)
	"$root/scripts/fetch-verified.sh" \
		"https://github.com/caddyserver/caddy/releases/download/$caddy_version/caddy_${caddy_version#v}_linux_amd64.tar.gz" \
		"$caddy_sha256" "$dl/caddy.tar.gz"
	mkdir -p "$root/.cache/bin"
	tar -xzf "$dl/caddy.tar.gz" -C "$dl" caddy
	mv "$dl/caddy" "$caddy"
	rm -rf "$dl"
	[ "$("$caddy" version | cut -d' ' -f1)" = "$caddy_version" ] || { echo "preview-lan: unexpected caddy version" >&2; exit 1; }
fi
for p in "$backend_port" "$https_port" "$cert_port"; do
	if ss -ltn "sport = :$p" | grep -q LISTEN; then echo "preview-lan: port $p is already in use" >&2; exit 1; fi
done

run=$(mktemp -d)
mkdir -p "$run/pub"; mkdir -p -m 700 "$data"
pids=()
cleanup() {
	trap '' INT TERM
	for pid in ${pids[@]+"${pids[@]}"}; do kill "$pid" 2>/dev/null || true; done
	for pid in ${pids[@]+"${pids[@]}"}; do wait "$pid" 2>/dev/null || true; done
	rm -rf "$run"
}
trap cleanup EXIT
trap 'exit 130' INT
trap 'exit 143' TERM

CASA_RUN_DIR=$run CASA_ORIGIN="https://$ip:$https_port" CASA_LISTEN="127.0.0.1:$backend_port" "$root/web/scripts/fullstack-server.sh" &
pids+=($!)
for _ in $(seq 300); do
	[ -e "$run/ready" ] && break
	kill -0 "${pids[0]}" 2>/dev/null || { echo "preview-lan: fullstack-server.sh stopped" >&2; exit 1; }
	sleep 1
done
[ -e "$run/ready" ] || { echo "preview-lan: casa not ready after 300 s" >&2; exit 1; }

# Caddy keeps its CA (and so the devices' trust) between runs in $data.
# Same as `caddy reverse-proxy --from ... --to ... --internal-certs`, in a
# Caddyfile so the admin API can be off. Origin is passed through unchanged.
cat >"$run/Caddyfile" <<CADDYFILE
{
	admin off
	local_certs
	auto_https disable_redirects
}
https://$ip:$https_port {
	reverse_proxy 127.0.0.1:$backend_port
}
CADDYFILE
XDG_DATA_HOME=$data XDG_CONFIG_HOME=$data "$caddy" run --config "$run/Caddyfile" --adapter caddyfile >"$run/caddy.log" 2>&1 &
pids+=($!)
crt=$data/caddy/pki/authorities/local/root.crt
for _ in $(seq 60); do
	[ -s "$crt" ] && break
	kill -0 "${pids[1]}" 2>/dev/null || { echo "preview-lan: caddy stopped; see $run/caddy.log" >&2; cat "$run/caddy.log" >&2; exit 1; }
	sleep 0.5
done
[ -s "$crt" ] || { echo "preview-lan: Caddy made no root certificate" >&2; exit 1; }
cp "$crt" "$run/pub/root.crt" # the public part only
python3 -I -m http.server "$cert_port" --bind "$ip" --directory "$run/pub" >/dev/null 2>&1 &
pids+=($!)

cat <<MSG

Casa preview is up. Stop with Ctrl-C.
  App (after trusting the certificate): https://$ip:$https_port/
  Root certificate to install (public part): http://$ip:$cert_port/root.crt
  Mail files (open the link inside from the device browser): $run/mail
  Setup tokens (0600, one per person): $run/setup-token-1 and $run/setup-token-2
    first sign-up link: https://$ip:$https_port/signup#<contents of the file>
MSG
wait -n "${pids[@]}" || true
echo "preview-lan: a process stopped; shutting down" >&2
