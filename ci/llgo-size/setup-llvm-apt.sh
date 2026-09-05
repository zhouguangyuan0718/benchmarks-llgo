#!/usr/bin/env bash
# Required input: LLVM_VERSION (numeric LLVM major version).
# This Ubuntu-only helper requires curl, gpg, and sudo. It installs the verified
# apt.llvm.org keyring and the matching versioned apt source system-wide.
set -euo pipefail

: "${LLVM_VERSION:?LLVM_VERSION must select the LLVM apt repository}"
if [[ ! "$LLVM_VERSION" =~ ^[0-9]+$ ]]; then
  echo "LLVM_VERSION must be a numeric major version: $LLVM_VERSION" >&2
  exit 2
fi

if [[ ! -r /etc/os-release ]]; then
  echo "/etc/os-release is unavailable; cannot select an LLVM apt repository" >&2
  exit 1
fi
# shellcheck disable=SC1091
source /etc/os-release
if [[ "${ID:-}" != "ubuntu" || -z "${VERSION_CODENAME:-}" ]]; then
  echo "LLVM apt setup requires Ubuntu with VERSION_CODENAME" >&2
  exit 1
fi

for command in curl gpg; do
  if ! command -v "$command" >/dev/null 2>&1; then
    echo "$command is required to configure the LLVM apt repository" >&2
    exit 1
  fi
done

key_file="$(mktemp)"
keyring_file="$(mktemp)"
gpg_home="$(mktemp -d)"
trap 'rm -f "$key_file" "$keyring_file"; rm -rf "$gpg_home"' EXIT
curl --fail --location --retry 3 --retry-all-errors \
  --connect-timeout 30 --max-time 120 \
  --output "$key_file" https://apt.llvm.org/llvm-snapshot.gpg.key

expected_fingerprint=6084F3CF814B57C1CF12EFD515CF4D18AF4F7421
gpg --batch --homedir "$gpg_home" --import "$key_file"
if ! gpg --batch --homedir "$gpg_home" --with-colons \
    --list-keys "$expected_fingerprint" |
    awk -F: -v expected="$expected_fingerprint" \
      '$1 == "fpr" && $10 == expected { found = 1 } END { exit !found }'; then
  echo "the LLVM apt signing key does not contain $expected_fingerprint" >&2
  exit 1
fi
gpg --batch --yes --homedir "$gpg_home" --output "$keyring_file" \
  --export "$expected_fingerprint"
if [[ ! -s "$keyring_file" ]]; then
  echo "failed to export the verified LLVM apt signing key" >&2
  exit 1
fi

keyring=/etc/apt/keyrings/apt.llvm.org.gpg
source_list=/etc/apt/sources.list.d/apt.llvm.org.list
sudo install -d -m 0755 /etc/apt/keyrings
sudo install -m 0644 "$keyring_file" "$keyring"
printf 'deb [signed-by=%s] https://apt.llvm.org/%s/ llvm-toolchain-%s-%s main\n' \
  "$keyring" "$VERSION_CODENAME" "$VERSION_CODENAME" "$LLVM_VERSION" |
  sudo tee "$source_list" >/dev/null
