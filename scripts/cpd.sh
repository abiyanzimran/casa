#!/usr/bin/env bash
# Duplication gate QG-02 (ES §2; ADR-025 item 4): PMD CPD over one language.
# Usage: scripts/cpd.sh <language> <dir>
set -euo pipefail

language=${1:?usage: cpd.sh <language> <dir>}
dir=${2:?usage: cpd.sh <language> <dir>}

# Token thresholds are set by code-reviewer, not by the implementer
# (ADR-025 item 4; 13 November 2026).
case "$language" in
  go) minimum_tokens=100 file_pattern='*.go' ;;
  *) echo "cpd.sh: no threshold set for language \"$language\"" >&2; exit 2 ;;
esac

pmd_version=7.29.0
# SHA-256 of pmd-dist-7.29.0-bin.zip from the release page:
# https://github.com/pmd/pmd/releases/tag/pmd_releases%2F7.29.0
pmd_sha256=d27177c185b837fa57f097e482d1253468aa035379c5eb2671fa0a279109b179

root=$(cd "$(dirname "$0")/.." && pwd)
pmd_cache="$root/.cache/pmd"
pmd="$pmd_cache/pmd-bin-$pmd_version/bin/pmd"
if [ ! -x "$pmd" ]; then
  zip="$pmd_cache/pmd-dist-$pmd_version-bin.zip"
  "$root/scripts/fetch-verified.sh" \
    "https://github.com/pmd/pmd/releases/download/pmd_releases%2F$pmd_version/pmd-dist-$pmd_version-bin.zip" \
    "$pmd_sha256" "$zip"
  unzip -q -o "$zip" -d "$pmd_cache"
fi

# Generated code, third-party code and build output are not scanned
# (code-reviewer, 13 November 2026). Test files are scanned.
file_list=$(mktemp)
trap 'rm -f "$file_list"' EXIT
find "$dir" \
  \( -name vendor -o -name node_modules -o -name .svelte-kit -o -name gen -o -name build -o -name dist \) -prune \
  -o -type f -name "$file_pattern" ! -name '*.pb.go' -print |
  while IFS= read -r file; do
    grep -qE '^// Code generated .* DO NOT EDIT\.$' "$file" || echo "$file"
  done >"$file_list"

if [ ! -s "$file_list" ]; then
  echo "cpd.sh: no $language files under $dir, so nothing would be checked" >&2
  exit 1
fi

"$pmd" cpd --language "$language" --minimum-tokens "$minimum_tokens" \
  --file-list "$file_list" --fail-on-violation
