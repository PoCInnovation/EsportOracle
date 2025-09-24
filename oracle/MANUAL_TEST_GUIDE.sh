# 🚀 Guide Manuel de Test - EsportOracle Betting System
# Copiez-collez ces commandes dans votre terminal

# Variables d'environnement
export TOKEN_ADDRESS="0x5FbDB2315678afecb367f032d93F642f64180aa3"
export ORACLE_ADDRESS="0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
export BET_CONTRACT_ADDRESS="0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
export RPC_URL="http://127.0.0.1:8545"

# Comptes Anvil pré-financés
export DEPLOYER="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
export USER1="0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
export USER2="0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"

# =====================================
# 🔍 1. VÉRIFICATIONS INITIALES
# =====================================

# Vérifier les soldes ETH
echo "=== SOLDES ETH ==="
cast balance $DEPLOYER --rpc-url $RPC_URL
cast balance $USER1 --rpc-url $RPC_URL
cast balance $USER2 --rpc-url $RPC_URL

# Vérifier les soldes tokens
echo "=== SOLDES TOKENS ==="
cast call $TOKEN_ADDRESS "balanceOf(address)" $DEPLOYER --rpc-url $RPC_URL
cast call $TOKEN_ADDRESS "balanceOf(address)" $USER1 --rpc-url $RPC_URL
cast call $TOKEN_ADDRESS "balanceOf(address)" $USER2 --rpc-url $RPC_URL

# Vérifier le nombre de paris
echo "=== NOMBRE DE PARIS ==="
cast call $BET_CONTRACT_ADDRESS "getBetCount()" --rpc-url $RPC_URL

# Vérifier le solde ETH du contrat
echo "=== SOLDE ETH DU CONTRAT ==="
cast call $BET_CONTRACT_ADDRESS "getContractETHBalance()" --rpc-url $RPC_URL

# =====================================
# 🎯 2. CRÉER UN PARI (en tant que DEPLOYER)
# =====================================

# Créer un pari
echo "=== CRÉATION D'UN PARI ==="
cast send $BET_CONTRACT_ADDRESS "createBet(string,uint256,uint256,uint256,uint256)" \
    "Team Liquid vs Team G2 Esports - Championship Final" \
    136396 \
    137372 \
    $(($(date +%s) + 3600)) \
    888 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --rpc-url $RPC_URL

# Vérifier que le pari a été créé
echo "=== VÉRIFICATION DU PARI CRÉÉ ==="
cast call $BET_CONTRACT_ADDRESS "getBetCount()" --rpc-url $RPC_URL
cast call $BET_CONTRACT_ADDRESS "getBet(uint256)" 0 --rpc-url $RPC_URL

# =====================================
# 🎮 3. TRANSFÉRER DES TOKENS AUX UTILISATEURS
# =====================================

echo "=== TRANSFERT DE TOKENS ==="
# Transférer 1000 tokens à USER1
cast send $TOKEN_ADDRESS "transfer(address,uint256)" \
    $USER1 \
    "1000000000000000000000" \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --rpc-url $RPC_URL

# Transférer 1000 tokens à USER2
cast send $TOKEN_ADDRESS "transfer(address,uint256)" \
    $USER2 \
    "1000000000000000000000" \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --rpc-url $RPC_URL

# Vérifier les nouveaux soldes
echo "=== NOUVEAUX SOLDES TOKENS ==="
cast call $TOKEN_ADDRESS "balanceOf(address)" $USER1 --rpc-url $RPC_URL
cast call $TOKEN_ADDRESS "balanceOf(address)" $USER2 --rpc-url $RPC_URL

# =====================================
# 🏆 4. USER1 PARIE SUR TEAM ALPHA (équipe 1)
# =====================================

echo "=== USER1 PARIE SUR TEAM ALPHA ==="
# USER1 approuve 500 tokens
cast send $TOKEN_ADDRESS "approve(address,uint256)" \
    $BET_CONTRACT_ADDRESS \
    "500000000000000000000" \
    --private-key 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d \
    --rpc-url $RPC_URL

# USER1 parie 500 tokens sur l'équipe 1 (Team Alpha)
cast send $BET_CONTRACT_ADDRESS "placeBet(uint256,uint8,uint256)" \
    0 \
    1 \
    "500000000000000000000" \
    --private-key 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d \
    --rpc-url $RPC_URL

# =====================================
# 🏆 5. USER2 PARIE SUR TEAM BETA (équipe 2)
# =====================================

echo "=== USER2 PARIE SUR TEAM BETA ==="
# USER2 approuve 300 tokens
cast send $TOKEN_ADDRESS "approve(address,uint256)" \
    $BET_CONTRACT_ADDRESS \
    "300000000000000000000" \
    --private-key 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a \
    --rpc-url $RPC_URL

# USER2 parie 300 tokens sur l'équipe 2 (Team Beta)
cast send $BET_CONTRACT_ADDRESS "placeBet(uint256,uint8,uint256)" \
    0 \
    2 \
    "300000000000000000000" \
    --private-key 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a \
    --rpc-url $RPC_URL

# =====================================
# 📊 6. VÉRIFIER L'ÉTAT DU PARI
# =====================================

echo "=== ÉTAT DU PARI ==="
# Voir les détails du pari
cast call $BET_CONTRACT_ADDRESS "getBet(uint256)" 0 --rpc-url $RPC_URL

# Voir le pari de USER1
cast call $BET_CONTRACT_ADDRESS "getUserBet(address,uint256)" $USER1 0 --rpc-url $RPC_URL

# Voir le pari de USER2
cast call $BET_CONTRACT_ADDRESS "getUserBet(address,uint256)" $USER2 0 --rpc-url $RPC_URL

# Voir les participants
cast call $BET_CONTRACT_ADDRESS "getBetParticipants(uint256)" 0 --rpc-url $RPC_URL

# Vérifier les soldes finaux
echo "=== SOLDES FINAUX ==="
cast call $TOKEN_ADDRESS "balanceOf(address)" $USER1 --rpc-url $RPC_URL
cast call $TOKEN_ADDRESS "balanceOf(address)" $USER2 --rpc-url $RPC_URL
cast call $TOKEN_ADDRESS "balanceOf(address)" $BET_CONTRACT_ADDRESS --rpc-url $RPC_URL

# =====================================
# 🎉 RÉSUMÉ DES RÉSULTATS
# =====================================

echo "=== RÉSUMÉ ==="
echo "✅ Pari créé : Team Liquid(27786) vs Team G2 Esports (46037)"
echo "✅ USER1 a parié 500 tokens sur Team Alpha"
echo "✅ USER2 a parié 300 tokens sur Team Beta"
echo "✅ Pool total : 800 tokens"
echo "✅ Si Team Liquid gagne : USER1 récupère 800 tokens (gain de 300)"
echo "✅ Si Team G2 Esports gagne : USER2 récupère 800 tokens (gain de 500)"
