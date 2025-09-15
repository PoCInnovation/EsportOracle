#!/usr/bin/env bash
set -euo pipefail

ANVIL_PORT=${ANVIL_PORT:-8545}
RPC_URL=${RPC_URL:-http://127.0.0.1:$ANVIL_PORT}

echo "Starting anvil on port $ANVIL_PORT..."
anvil --silent --balance 1000 --port $ANVIL_PORT --block-time 1 --disable-code-size-limit &
ANVIL_PID=$!
trap 'kill $ANVIL_PID' EXIT

sleep 1

EOA_PK1=${EOA_PK1:-0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80}
EOA_PK2=${EOA_PK2:-0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d}

export EOA_PK1
export EOA_PK2

echo "Broadcasting Foundry script..."
export FOUNDRY_IGNORE_CONTRACT_SIZE=1
yes | forge script oracle/script/OracleBroadcast.s.sol:OracleBroadcast \
  --rpc-url $RPC_URL \
  --broadcast \
  --force \
  --via-ir \
  --optimize \
  --optimizer-runs 200 \
  --skip-simulation \
  --non-interactive | cat

echo
echo "Artifacts in broadcast directory:"
ls -la broadcast/OracleBroadcast.s.sol/31337/ || true

echo "Done."
