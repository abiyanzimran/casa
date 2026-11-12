#!/usr/bin/env bash
# Downloads a pinned tool archive and keeps it only if its SHA-256 matches the
# value recorded in the repository (ES §8: tools are pinned exactly).
# Usage: scripts/fetch-verified.sh <url> <sha256> <dest>
set -euo pipefail

url=${1:?usage: fetch-verified.sh <url> <sha256> <dest>}
sha256=${2:?usage: fetch-verified.sh <url> <sha256> <dest>}
dest=${3:?usage: fetch-verified.sh <url> <sha256> <dest>}

mkdir -p "$(dirname "$dest")"
tmp=$(mktemp "$dest.XXXXXX")
trap 'rm -f "$tmp"' EXIT
curl -sSfL --retry 3 -o "$tmp" "$url"
echo "$sha256  $tmp" | sha256sum --check --quiet -
mv "$tmp" "$dest"
