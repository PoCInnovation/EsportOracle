#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
TMP_ENV="$ROOT_DIR/tmp/local-devchain.env"

# shellcheck disable=SC1090
if [[ -f "$TMP_ENV" ]]; then
  while IFS= read -r line; do
    [[ -z "$line" || "$line" == \#* ]] && continue
    key=${line%%=*}
    value=${line#*=}
    export "$key"="$value"
  done < "$TMP_ENV"
fi

# Allow overrides via the current shell environment
: "${ETHEREUM_RPC_URL:=http://127.0.0.1:8545}"
: "${BET_CONTRACT_ADDRESS:=}"
: "${TOKEN_ADDRESS:=}"

if [[ -z "$BET_CONTRACT_ADDRESS" ]]; then
  echo "[error] BET_CONTRACT_ADDRESS is not set. Run scripts/bootstrap_local_chain.sh first." >&2
  exit 1
fi

if [[ -z "$TOKEN_ADDRESS" ]]; then
  echo "[error] TOKEN_ADDRESS is not set. Run scripts/bootstrap_local_chain.sh first." >&2
  exit 1
fi

export PRIVATE_KEY=${PRIVATE_KEY:-0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80}
export BETTOR1_PRIVATE_KEY=${BETTOR1_PRIVATE_KEY:-0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d}
export BETTOR2_PRIVATE_KEY=${BETTOR2_PRIVATE_KEY:-0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a}
export BET_CONTRACT_ADDRESS
export TOKEN_ADDRESS

forge script \
  --root "$ROOT_DIR/oracle" \
  --rpc-url "$ETHEREUM_RPC_URL" \
  --broadcast \
  oracle/script/SeedLocalBets.s.sol:SeedLocalBets
