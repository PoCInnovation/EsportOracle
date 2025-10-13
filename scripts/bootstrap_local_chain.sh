#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
TMP_DIR="$ROOT_DIR/tmp"
mkdir -p "$TMP_DIR"

if ! command -v anvil >/dev/null 2>&1; then
  echo "[error] anvil not found. Install Foundry (https://book.getfoundry.sh/getting-started/installation)." >&2
  exit 1
fi

if ! command -v forge >/dev/null 2>&1; then
  echo "[error] forge not found. Install Foundry." >&2
  exit 1
fi

if ! command -v cast >/dev/null 2>&1; then
  echo "[error] cast not found. Install Foundry." >&2
  exit 1
fi

if ! command -v python3 >/dev/null 2>&1; then
  echo "[error] python3 not found. Install Python 3 for address parsing." >&2
  exit 1
fi

ensure_foundry_deps() {
  local oracle_dir="$ROOT_DIR/oracle"

  if [[ -d "$oracle_dir/lib/forge-std" && ! -d "$oracle_dir/lib/forge-std/src" ]]; then
    rm -rf "$oracle_dir/lib/forge-std"
  fi
  if [[ -d "$oracle_dir/lib/openzeppelin-contracts" && ! -d "$oracle_dir/lib/openzeppelin-contracts/contracts" ]]; then
    rm -rf "$oracle_dir/lib/openzeppelin-contracts"
  fi

  if [[ ! -d "$oracle_dir/lib/forge-std/src" ]]; then
    echo "[info] installing forge-std dependency"
    forge install --no-git --root "$oracle_dir" foundry-rs/forge-std@v1.9.6
  fi

  if [[ ! -d "$oracle_dir/lib/openzeppelin-contracts/contracts" ]]; then
    echo "[info] installing openzeppelin-contracts dependency"
    forge install --no-git --root "$oracle_dir" openzeppelin/openzeppelin-contracts@v5.0.1
  fi
}

ensure_foundry_deps

RPC_BIND_HOST=${ANVIL_HOST:-0.0.0.0}
RPC_CONNECT_HOST=${ANVIL_CONNECT_HOST:-127.0.0.1}
RPC_PORT=${ANVIL_PORT:-8545}
CHAIN_ID=${ANVIL_CHAIN_ID:-31337}
RPC_URL="http://${RPC_CONNECT_HOST}:${RPC_PORT}"
CONTAINER_RPC_URL=${ANVIL_CONTAINER_RPC_URL:-http://host.docker.internal:${RPC_PORT}}

ANVIL_PID_FILE="$TMP_DIR/anvil.pid"
ANVIL_LOG="$TMP_DIR/anvil.log"

start_anvil() {
  if [[ -f "$ANVIL_PID_FILE" ]]; then
    if kill -0 "$(cat "$ANVIL_PID_FILE")" >/dev/null 2>&1; then
      echo "[info] anvil already running (pid $(cat "$ANVIL_PID_FILE"))"
      return
    fi
    rm -f "$ANVIL_PID_FILE"
  fi

  echo "[info] starting anvil on ${RPC_BIND_HOST}:${RPC_PORT} (chain id ${CHAIN_ID})"
  echo "[info] local RPC available via ${RPC_URL}"
  anvil \
    --host "$RPC_BIND_HOST" \
    --port "$RPC_PORT" \
    --chain-id "$CHAIN_ID" \
    --mnemonic "test test test test test test test test test test test junk" \
    --block-time 2 \
    --base-fee 0 \
    --gas-price 0 \
    >"$ANVIL_LOG" 2>&1 &
  ANVIL_PID=$!
  echo "$ANVIL_PID" > "$ANVIL_PID_FILE"
  echo "[info] anvil pid $ANVIL_PID (logs: $ANVIL_LOG)"

  echo -n "[info] waiting for anvil to be ready"
  for _ in {1..40}; do
    if curl -s "$RPC_URL" >/dev/null; then
      echo " - ready"
      return
    fi
    echo -n "."
    sleep 0.5
  done
  echo
  echo "[error] anvil did not start in time. Check $ANVIL_LOG" >&2
  exit 1
}

start_anvil

# Default private keys (anvil deterministic accounts); allow overrides through env
DEPLOYER_PK=${PRIVATE_KEY:-${DEPLOYER_PK:-0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80}}
BETTOR1_PK=${BETTOR1_PRIVATE_KEY:-0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d}
BETTOR2_PK=${BETTOR2_PRIVATE_KEY:-0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a}

DEPLOY_LOG="$TMP_DIR/deploy_bet.log"
SEED_LOG="$TMP_DIR/seed_bet.log"
BROADCAST_DIR="$ROOT_DIR/oracle/broadcast"
BROADCAST_FILE="$BROADCAST_DIR/DeployBetContract.s.sol/${CHAIN_ID}/run-latest.json"

mkdir -p "$BROADCAST_DIR"

export PRIVATE_KEY="$DEPLOYER_PK"

echo "[info] deploying oracle + bet contracts"
forge script \
  --root "$ROOT_DIR/oracle" \
  --rpc-url "$RPC_URL" \
  --broadcast \
  --slow \
  oracle/script/DeployBetContract.s.sol:DeployBetContract \
  | tee "$DEPLOY_LOG"

declare TOKEN_ADDRESS=""
declare BET_ADDRESS=""
declare ORACLE_ADDRESS=""

if [[ -f "$BROADCAST_FILE" ]]; then
  while IFS== read -r key value; do
    case "$key" in
      BET_CONTRACT_ADDRESS) BET_ADDRESS="$value" ;;
      TOKEN_ADDRESS) TOKEN_ADDRESS="$value" ;;
      ORACLE_ADDRESS) ORACLE_ADDRESS="$value" ;;
    esac
  done < <(python3 - "$BROADCAST_FILE" <<'PY'
import json, sys
path = sys.argv[1]
with open(path) as f:
    data = json.load(f)
addresses = {}
for tx in data.get("transactions", []):
    name = tx.get("contractName")
    addr = tx.get("contractAddress")
    if name and addr:
        addresses[name] = addr
for key, name in [("BET_CONTRACT_ADDRESS", "BetContract"), ("TOKEN_ADDRESS", "MockERC20"), ("ORACLE_ADDRESS", "EsportOracle")]:
    print(f"{key}={addresses.get(name, '')}")
PY
  )
fi

# Fallback parsing from log if broadcast file missing contract names
if [[ -z "$BET_ADDRESS" || -z "$TOKEN_ADDRESS" ]]; then
  BET_ADDRESS=$(grep -Eo 'BetContract deployed at:\s*0x[0-9a-fA-F]+' "$DEPLOY_LOG" | awk '{print $4}' | tail -n1 || true)
  TOKEN_ADDRESS=$(grep -Eo 'MockERC20 deployed at:\s*0x[0-9a-fA-F]+' "$DEPLOY_LOG" | awk '{print $4}' | tail -n1 || true)
  ORACLE_ADDRESS=$(grep -Eo 'EsportOracle deployed at:\s*0x[0-9a-fA-F]+' "$DEPLOY_LOG" | awk '{print $4}' | tail -n1 || true)
fi

if [[ -z "$BET_ADDRESS" || -z "$TOKEN_ADDRESS" ]]; then
  echo "[error] failed to determine deployed contract addresses" >&2
  exit 1
fi

echo "[info] BetContract: $BET_ADDRESS"
echo "[info] Token:       $TOKEN_ADDRESS"
echo "[info] Oracle:      ${ORACLE_ADDRESS:-unknown}"

export BET_CONTRACT_ADDRESS="$BET_ADDRESS"
export TOKEN_ADDRESS="$TOKEN_ADDRESS"
export BETTOR1_PRIVATE_KEY="$BETTOR1_PK"
export BETTOR2_PRIVATE_KEY="$BETTOR2_PK"

# Seed fake bets and liquidity
echo "[info] seeding sample bets"
forge script \
  --root "$ROOT_DIR/oracle" \
  --rpc-url "$RPC_URL" \
  --broadcast \
  oracle/script/SeedLocalBets.s.sol:SeedLocalBets \
  | tee "$SEED_LOG"

echo "[info] seeding finished"

ENV_SUGGESTION="$ROOT_DIR/tmp/local-devchain.env"
cat > "$ENV_SUGGESTION" <<EOV
ETHEREUM_RPC_URL=${RPC_URL}
ETHEREUM_RPC_URL_CONTAINER=${CONTAINER_RPC_URL}
BET_CONTRACT_ADDRESS=${BET_ADDRESS}
TOKEN_ADDRESS=${TOKEN_ADDRESS}
BETTOR1_ADDRESS=$(cast wallet address --private-key ${BETTOR1_PK})
BETTOR2_ADDRESS=$(cast wallet address --private-key ${BETTOR2_PK})
EOV

echo "[info] wrote suggested env values to $ENV_SUGGESTION"
echo
cat <<EOM
Local chain ready.
Add these entries to your .env before running docker compose:
  ETHEREUM_RPC_URL=${RPC_URL}
  ETHEREUM_RPC_URL_CONTAINER=${CONTAINER_RPC_URL}
  BET_CONTRACT_ADDRESS=${BET_ADDRESS}

Optional helpers saved in $ENV_SUGGESTION (token + bettor addresses).
Keep anvil running while the stack is up. Stop it with:
  kill \\$(cat $ANVIL_PID_FILE)
EOM
