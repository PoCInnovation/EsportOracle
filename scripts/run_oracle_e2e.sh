#!/usr/bin/env bash
set -euo pipefail

ANVIL_PORT=${ANVIL_PORT:-8545}
RPC_URL=${RPC_URL:-http://127.0.0.1:$ANVIL_PORT}

echo "Starting anvil on port $ANVIL_PORT..."
anvil --silent --balance 1000 --port $ANVIL_PORT --block-time 1 --code-size-limit 0 &
ANVIL_PID=$!
trap 'kill $ANVIL_PID' EXIT

sleep 1

EOA_PK1=${EOA_PK1:-0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80}
EOA_PK2=${EOA_PK2:-0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d}
EOA_PK3=${EOA_PK3:-0x5de411648b1721dc9f6f07dc7174e43b0e6b0bacedbba1b6a0aee9b5a27f7bcd}

export EOA_PK1
export EOA_PK2
export EOA_PK3

echo "Running Foundry script..."
export FOUNDRY_IGNORE_CONTRACT_SIZE=1
forge script oracle/script/OracleE2E.s.sol:OracleE2E \
  --rpc-url $RPC_URL \
  --force \
  --via-ir \
  --optimize \
  --optimizer-runs 200 \
  --non-interactive | cat

BROADCAST_JSON="broadcast/OracleE2E.s.sol/31337/dry-run/run-latest.json"
TARGET_JSON="cache/OracleE2E.s.sol/31337/dry-run/run-latest.json"
if [ -f "$BROADCAST_JSON" ]; then
  mkdir -p "$(dirname \"$TARGET_JSON\")"
  cp "$BROADCAST_JSON" "$TARGET_JSON"
  echo "\nSaved detailed dry-run transactions to $TARGET_JSON"
else
  echo "\nNo broadcast dry-run file found at $BROADCAST_JSON"
fi

echo "Done."


