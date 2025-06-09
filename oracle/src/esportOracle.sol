// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Pausable} from "openzeppelin-contracts/contracts/utils/Pausable.sol";

contract EsportOracle is Pausable {
    address public _owner;

    struct Result {
        uint8 _score;
        uint256 _teamId;
    }

    struct Games {
        uint256 _id;
        bool _finished;
        uint256 _winnerId;
    }

    struct Opponents {
        string _acronym;
        uint256 _id;
        string _name;
    }

    struct Match {
        uint256 _id;
        Opponents[] _opponents;
        Games[] _game;
        Result[] _result;
        uint256 _winnerId;
        uint256 _beginAt;
    }

    // Structure pour suivre les violations des nœuds
    struct NodeViolation {
        uint256 incorrectMatches;
        bool isBanned;
    }

    mapping(uint256 => Match) public _matchMapping;
    address[] private listedNodes;
    mapping(address => uint256) public _fundsStaked;
    mapping(bytes32 => uint8) public _matchVotes;
    mapping(bytes32 => address[]) public _addressByHash;
    bytes32[] public _pendingMatchesHashes;
    uint8 nbMatchSent;

    mapping(uint256 => bytes32[]) public _matchIdToHashes;
    mapping(bytes32 => uint256) public _hashToMatchId;

    mapping(address => NodeViolation) public _nodeViolations;
    uint256 public constant MAX_VIOLATIONS = 3;
    uint256 public constant PUNISHMENT_AMOUNT = 0.0001 ether;

    constructor() {
        _owner = msg.sender;
        nbMatchSent = 0;
    }

    // EVENTS
    event newNodeAdded(address indexed addressAdded);
    event stakingSuccess(address indexed addressAdded, uint256 amount);
    event NodePunished(address indexed node, uint256 amount, uint256 violationsCount);
    event NodeBanned(address indexed node);

    // MODIFIERS
    modifier onlyOwner() {
        require(msg.sender == _owner, "Not the contract owner");
        _;
    }

    modifier notBanned(address node) {
        require(!_nodeViolations[node].isBanned, "Node is banned");
        _;
    }

    modifier onlyListedNodes() {
        bool isListed = false;

        for (uint i = 0; i < listedNodes.length; i++) {
            if (listedNodes[i] == msg.sender) {
                isListed = true;
                break;
            }
        }
        require(isListed == true, "Node is not listed, please call addNewNode function to register a new node");
        _;
    }

    modifier nodeAlreadyStake() {
        bool isListed = false;

        for (uint i = 0; i < listedNodes.length; i++) {
            if (listedNodes[i] == msg.sender) {
                isListed = true;
                break;
            }
        }
        require(isListed == false, "Already staked");
        _;
    }

    //////////////// GOVERNANCE FUNCTIONS ////////////////

    /**
     * @notice Set the new owner of the contract
     * @param newOwner The address of the new owner
     * @dev Only the current owner can call this function
     */
    function setOwner(address newOwner) public onlyOwner {
        require(newOwner != address(0), "New owner cannot be zero address");
        _owner = newOwner;
    }

    /**
     * @notice Permet au propriétaire de mettre le contrat en pause
     * @dev Utilise le modificateur onlyOwner et la fonction _pause de Pausable
     */
    function pause() external onlyOwner {
        _pause();
    }

    /**
     * @notice Permet au propriétaire de sortir le contrat de l'état de pause
     * @dev Utilise le modificateur onlyOwner et la fonction _unpause de Pausable
     */
    function unpause() external onlyOwner {
        _unpause();
    }

    //////////////// STAKING FUNCTIONS ////////////////

    /**
     * @notice Allows nodes to stake funds
     * @dev Uses low-level call to transfer funds with gas optimization
    */
    function addFundToStaking() external payable whenNotPaused notBanned(msg.sender) nodeAlreadyStake {
        require(
            msg.sender != address(0) &&
            msg.sender != address(this),
            "Invalid staking parameters"
        );
        require(msg.value == 0.001 ether, "amount must be exactly 0.001 ether");
        _fundsStaked[msg.sender] = msg.value;
        emit stakingSuccess(msg.sender, msg.value);
        addNewNode();
    }

    /**
     * @notice Fonction permettant à un nœud de retirer ses fonds s'il n'est pas banni
     */
    function withdrawStake() external whenNotPaused notBanned(msg.sender) {
        uint256 amount = _fundsStaked[msg.sender];
        require(amount > 0, "No funds to withdraw");

        for (uint i = 0; i < listedNodes.length; i++) {
            if (listedNodes[i] == msg.sender) {
                listedNodes[i] = listedNodes[listedNodes.length - 1];
                listedNodes.pop();
                break;
            }
        }
        _fundsStaked[msg.sender] = 0;
        (bool success, ) = msg.sender.call{value: amount}("");
        require(success, "Transfer failed");
    }

    //////////////// NODE MANAGEMENT FUNCTIONS ////////////////

    /**
     * @notice function to add a new node
     * @dev Gas optimized by combining requirements
     */
    function addNewNode() internal nodeAlreadyStake {
        require(msg.sender != address(0),"New node cannot be zero address");
        listedNodes.push(msg.sender);
        emit newNodeAdded(msg.sender);
    }

    /**
     * @notice Fonction pour supprimer un nœud manuellement par le propriétaire
     * @param node Adresse du nœud à supprimer
     */
    function deleteNode(address node) internal {
        for (uint i = 0; i < listedNodes.length; i++) {
            if (listedNodes[i] == node) {
                listedNodes[i] = listedNodes[listedNodes.length - 1];
                listedNodes.pop();
                break;
            }
        }
    }

    /**
     * @notice function to return the list of nodes addresses
     * @return The list of addresses of the nodes
     */
    function getListedNodes() external view returns (address[] memory) {
        return listedNodes;
    }

    //////////////// PUNISHMENT SYSTEM FUNCTIONS ////////////////

    /**
     * @notice Fonction pour punir un nœud qui a soumis des données incorrectes
     * @param node Adresse du nœud à punir
     */
    function punishNode(address node, address[] memory correctVoters) internal {
        require(_fundsStaked[node] > 0, "Node has no staked funds");

        _nodeViolations[node].incorrectMatches++;

        uint256 correctVotersLen = correctVoters.length;
        if (_nodeViolations[node].incorrectMatches >= MAX_VIOLATIONS) {
            _nodeViolations[node].isBanned = true;
            emit NodeBanned(node);

            uint256 amountToSlash = _fundsStaked[node];
            _fundsStaked[node] = 0;

            deleteNode(node);
            // Distribuer les fonds aux autres nœuds - optimisé pour réduire le gas
            if (correctVotersLen > 0) {
                uint256 amountPerNode = amountToSlash / correctVotersLen;
                for (uint i = 0; i < correctVotersLen;) {
                    _fundsStaked[correctVoters[i]] += amountPerNode;
                    unchecked { ++i; }
                }
            }
        } else {
            // Si le noeud n'est pas encore banni, appliquer la punition
            uint256 amountToSlash = PUNISHMENT_AMOUNT;
            require(_fundsStaked[node] >= amountToSlash, "Insufficient staked funds");

            _fundsStaked[node] -= amountToSlash;

            emit NodePunished(node, amountToSlash, _nodeViolations[node].incorrectMatches);

            // Distribuer les fonds aux autres nœuds
            if (correctVotersLen > 0) {
                uint256 amountPerNode = amountToSlash / correctVotersLen;
                // Utilisation d'une boucle non-vérifiée pour économiser du gas
                for (uint i = 0; i < correctVotersLen;) {
                    _fundsStaked[correctVoters[i]] += amountPerNode;
                    unchecked { ++i; }
                }
            }
        }
    }

    /**
     * @notice Fonction pour bannir un nœud manuellement par le propriétaire
     * @param node Adresse du nœud à bannir
     */
    function banNode(address node) external onlyOwner {
        require(!_nodeViolations[node].isBanned, "Node is already banned");

        _nodeViolations[node].isBanned = true;
        emit NodeBanned(node);

        /// Confisquer tous les fonds stakés
        uint256 amountToSlash = _fundsStaked[node];
        _fundsStaked[node] = 0;

        /// Redistribution des fonds optimisée pour réduire le gas
        uint256 remainingNodes = listedNodes.length - 1;
        if (remainingNodes > 0) {
            uint256 amountToDistribute = amountToSlash / remainingNodes;

            // Utilisation d'une boucle non-vérifiée pour économiser du gas
            for (uint i = 0; i < listedNodes.length;) {
                if (listedNodes[i] != node) {
                    _fundsStaked[listedNodes[i]] += amountToDistribute;
                }
                // Incrémentation non-vérifiée pour économiser du gas
                unchecked { ++i; }
            }
        }
        deleteNode(node);
    }

    /**
     * @notice Fonction pour réhabiliter un nœud banni, uniquement par le propriétaire
     * @param node Adresse du nœud à réhabiliter
     */
    function rehabilitateNode(address node) external onlyOwner {
        require(_nodeViolations[node].isBanned, "Node is not banned");

        _nodeViolations[node].isBanned = false;
        _nodeViolations[node].incorrectMatches = 0;
    }

    //////////////// MATCH HANDLING FUNCTIONS ////////////////

    /**
     * @notice function called by listed nodes only, to register new matches
     * @param newMatch : a list of matches to register
     */
    function handleNewMatches(Match[] memory newMatch) external whenNotPaused notBanned(msg.sender) onlyListedNodes {
        require(newMatch.length > 0, "No match data provided");
        nbMatchSent++;

        for (uint256 i = 0; i < newMatch.length; i++) {
            bytes32 matchHash = keccak256(abi.encode(newMatch[i]));
            bool alreadyVoted = false;
            uint256 matchId = newMatch[i]._id;

            // Associer le hash à l'ID du match
            _hashToMatchId[matchHash] = matchId;

            // Ajouter le hash à la liste des hashs pour cet ID de match
            bool hashExists = false;
            for (uint j = 0; j < _matchIdToHashes[matchId].length; j++) {
                if (_matchIdToHashes[matchId][j] == matchHash) {
                    hashExists = true;
                    break;
                }
            }
            if (!hashExists) {
                _matchIdToHashes[matchId].push(matchHash);
            }

            /// Vérifier si l'adresse a déjà voté pour ce match
            for (uint j = 0; j < _addressByHash[matchHash].length; j++) {
                if (_addressByHash[matchHash][j] == msg.sender) {
                    alreadyVoted = true;
                    break;
                }
            }

            // Incrémenter les votes seulement si le nœud n'a pas déjà voté
            if (!alreadyVoted) {
                _matchVotes[matchHash]++;
                _addressByHash[matchHash].push(msg.sender);
            }

            if (_matchVotes[matchHash] == 1)
                _pendingMatchesHashes.push(matchHash);

            if (quorumIsReached(_matchVotes[matchHash])) {
                uint256 validMatchId = newMatch[i]._id;
                addNewMatch(newMatch[i]);

                for (uint j = 0; j < _pendingMatchesHashes.length; j++) {
                    bytes32 currentHash = _pendingMatchesHashes[j];

                    if (currentHash != matchHash) {
                        uint256 currentMatchId = _hashToMatchId[currentHash];
                        bool isConflictingMatch = (currentMatchId == validMatchId && currentMatchId != 0);

                        address[] memory votersForCurrentHash = _addressByHash[currentHash];

                        if (isConflictingMatch) {
                            for (uint k = 0; k < votersForCurrentHash.length; k++) {
                                if (!_nodeViolations[votersForCurrentHash[k]].isBanned) {
                                    punishNode(votersForCurrentHash[k], _addressByHash[matchHash]);
                                }
                            }
                        }
                        delete _matchVotes[currentHash];
                        delete _addressByHash[currentHash];
                    }
                }
                delete _matchVotes[matchHash];
                delete _addressByHash[matchHash];

                bytes32[] memory newPendingHashes = new bytes32[](_pendingMatchesHashes.length - 1);
                uint256 newIndex = 0;

                for (uint p = 0; p < _pendingMatchesHashes.length; p++) {
                    if (_pendingMatchesHashes[p] != matchHash) {
                        newPendingHashes[newIndex] = _pendingMatchesHashes[p];
                        newIndex++;
                    }
                }

                // Mettre à jour le tableau des hashs en attente
                delete _pendingMatchesHashes;
                for (uint p = 0; p < newIndex; p++) {
                    _pendingMatchesHashes.push(newPendingHashes[p]);
                }

                // Nettoyer les hashs obsolètes pour ce match ID
                for (uint m = 0; m < _matchIdToHashes[validMatchId].length; m++) {
                    bytes32 hashToCheck = _matchIdToHashes[validMatchId][m];
                    if (hashToCheck != matchHash) {
                        delete _hashToMatchId[hashToCheck];
                    }
                }
                // Ne conserver que le hash validé pour ce match ID
                delete _matchIdToHashes[validMatchId];
                _matchIdToHashes[validMatchId].push(matchHash);
            }
        }

        /// Réinitialiser le compteur si tous les nœuds ont envoyé leurs données
        if (nbMatchSent == listedNodes.length)
            nbMatchSent = 0;
    }

    /**
     * @notice add match blockchain
     * @param newMatch a tab of a Match
     */
    function addNewMatch(Match memory newMatch) internal {
        uint256 matchId = newMatch._id;

        // SECTION 1: INITIAL MATCH CREATION
        // If the match doesn't exist yet (ID = 0), initialize it with all data
        if (_matchMapping[matchId]._id == 0) {
            delete _matchMapping[matchId];
            _matchMapping[matchId]._id = matchId;
            _matchMapping[matchId]._winnerId = newMatch._winnerId;
            _matchMapping[matchId]._beginAt = newMatch._beginAt;

            // Copy all opponents data
            for (uint256 j = 0; j < newMatch._opponents.length; j++) {
                Opponents memory opponent = newMatch._opponents[j];
                _matchMapping[matchId]._opponents.push(opponent);
            }

            // Copy all games data
            for (uint256 j = 0; j < newMatch._game.length; j++) {
                Games memory game = newMatch._game[j];
                _matchMapping[matchId]._game.push(game);
            }

            // Copy all results data
            for (uint256 j = 0; j < newMatch._result.length; j++) {
                Result memory result = newMatch._result[j];
                _matchMapping[matchId]._result.push(result);
            }
            return;
        }

        // SECTION 2: UPDATE WINNER ID
        if (newMatch._winnerId != _matchMapping[matchId]._winnerId) {
            _matchMapping[matchId]._winnerId = newMatch._winnerId;
        }

        // SECTION 3: UPDATE GAMES DATA
        bytes32 currentGameHash = keccak256(abi.encode(_matchMapping[matchId]._game));
        bytes32 newGameHash = keccak256(abi.encode(newMatch._game));

        if (currentGameHash != newGameHash) {
            delete (_matchMapping[matchId]._game);
            for (uint256 j = 0; j < newMatch._game.length; j++) {
                Games memory game = newMatch._game[j];
                _matchMapping[matchId]._game.push(game);
            }
        }

        // SECTION 4: UPDATE RESULTS DATA
        bytes32 currentResultHash = keccak256(abi.encode(_matchMapping[matchId]._result));
        bytes32 newResultHash = keccak256(abi.encode(newMatch._result));

        if (currentResultHash != newResultHash) {
            delete (_matchMapping[matchId]._result);
            for (uint256 j = 0; j < newMatch._result.length; j++) {
                Result memory result = newMatch._result[j];
                _matchMapping[matchId]._result.push(result);
            }
        }

        // SECTION 5: UPDATE OPPONENTS DATA
        bytes32 currentOpponentHash = keccak256(abi.encode(_matchMapping[matchId]._opponents));
        bytes32 newOppenentHash = keccak256(abi.encode(newMatch._opponents));

        if (currentOpponentHash != newOppenentHash) {
            delete (_matchMapping[matchId]._opponents);
            for (uint256 j = 0; j < newMatch._opponents.length; j++) {
                Opponents memory opponent = newMatch._opponents[j];
                _matchMapping[matchId]._opponents.push(opponent);
            }
        }
    }

    /**
     * @notice returns the match by id
     * @param matchId The id of the match
     * @return The match object
     * @dev This function retrieves a match by its ID from the mapping
     */
    function getMatchById(uint256 matchId) external view returns (Match memory) {
        return (_matchMapping[matchId]);
    }

    /**
     * @notice function to return the list pending match hash
     * @return The list of hashes
     */
    function getPendingMatches() external view returns (bytes32[] memory) {
        return (_pendingMatchesHashes);
    }

    /**
     * @notice Vérifie si le quorum est atteint pour un match
     * @param nbVote Nombre de votes reçus pour un match
     * @return true si le quorum est atteint, false sinon
     */
    function quorumIsReached(uint8 nbVote) private view returns (bool) {
        if (listedNodes.length <= 2) {
            return nbVote >= listedNodes.length; // Tous les nœuds doivent voter si 2 ou moins
        } else if (listedNodes.length == 3) {
            return nbVote >= 2; // Au moins 2 votes sur 3
        } else {
            return nbVote > (listedNodes.length / 2); // Majorité simple pour 4+ nœuds
        }
    }
}
