#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."

addon_version="$(go run ./cmd/release validate)"
stage="$(mktemp -d)"
trap 'rm -rf "$stage"' EXIT
mkdir -p "$stage/BlessingBuddy" .release
for file in BlessingBuddy.lua Bindings.xml README.md LICENSE CHANGELOG.md RELEASE_NOTES.md; do
	cp "$file" "$stage/BlessingBuddy/"
done
cp BlessingBuddy.toc "$stage/BlessingBuddy/"
archive="$PWD/.release/BlessingBuddy-$addon_version.zip"
rm -f "$archive"
(cd "$stage" && touch -t 198001010000 BlessingBuddy/* && zip -X -q "$archive" BlessingBuddy/*)
package_hash="$(shasum -a 256 "$archive" | cut -d ' ' -f 1)"
notes_hash="$(shasum -a 256 RELEASE_NOTES.md | cut -d ' ' -f 1)"
commit="$(git rev-parse --verify HEAD 2>/dev/null || echo uncommitted)"
cat > "$archive.manifest.txt" <<EOF
version=$addon_version
commit=$commit
flavor=WoW Forever
interface=$(sed -n 's/^## Interface: //p' BlessingBuddy.toc)
channel=release
file=$(basename "$archive")
sha256=$package_hash
notes_sha256=$notes_hash
curseforge_project=1724693
EOF
cat "$archive.manifest.txt"
