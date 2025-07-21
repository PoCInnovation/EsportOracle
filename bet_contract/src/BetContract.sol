// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "../lib/openzeppelin-contracts/contracts/access/Ownable.sol";
import "../lib/openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol";
import "../lib/openzeppelin-contracts/contracts/utils/Pausable.sol";

contract BetContract is Ownable, ReentrancyGuard, Pausable {
    struct Bet {
        string description;       // Description du pari
        string team1;            // Nom de la première équipe
        string team2;            // Nom de la deuxième équipe
        uint256 deadline;        // Date limite pour parier
        uint256 team1Pool;       // Montant total misé sur équipe 1
        uint256 team2Pool;       // Montant total misé sur équipe 2
        uint8 winningTeam;       // 1 = team1 gagne, 2 = team2 gagne, 0 = pas résolu
        bool resolved;           // Pari résolu ou non
        address creator;         // Créateur du pari
    }

    struct UserBet {
        uint256 amount;
        uint8 teamChosen;        // ✅ Changé de uint256 à uint8 pour cohérence
        bool claimed;
    }

    Bet[] public CurrentBets;
    address public oracle;
    mapping(address => mapping(uint256 => UserBet)) public userBets;
    mapping(uint256 => address[]) public betParticipants; // ✅ Ajout du mapping manquant
    IERC20 public bettingToken;

    event BetPlaced(address indexed user, uint256 indexed betId, uint8 team, uint256 amount);
    event WinningsClaimed(address indexed user, uint256 indexed betId, uint256 payout);
    event BetResolved(uint256 indexed betId, uint8 winningTeam);
    event BetCreated(uint256 indexed betId, string description, string team1, string team2, uint256 deadline); // ✅ Ajout d'événement

    modifier onlyOracle() {
        require(msg.sender == oracle, "Only oracle can call this");
        _;
    }

    constructor(address _oracle, address _tokenAddress) Ownable(msg.sender) {
        oracle = _oracle;
        bettingToken = IERC20(_tokenAddress);
    }

    function createBet(
        string memory description,
        string memory team1,
        string memory team2,
        uint256 deadline
    ) external {
        require(deadline > block.timestamp, "Deadline must be in the future");

        Bet memory newBet = Bet({
            description: description,
            team1: team1,
            team2: team2,
            deadline: deadline,
            team1Pool: 0,
            team2Pool: 0,
            winningTeam: 0,
            resolved: false,
            creator: msg.sender
        });

        CurrentBets.push(newBet);

        uint256 betId = CurrentBets.length - 1;
        emit BetCreated(betId, description, team1, team2, deadline);
    }

    function placeBet(uint256 betId, uint8 team, uint256 amount) external whenNotPaused {
        require(betId < CurrentBets.length, "Invalid bet ID"); // ✅ Vérification ajoutée
        require(CurrentBets[betId].deadline > block.timestamp, "Pari expire");
        require(team == 1 || team == 2, "Equipe invalide");
        require(amount > 0, "Montant invalide");

        // Transfert des tokens du joueur vers le contrat
        require(bettingToken.transferFrom(msg.sender, address(this), amount), "Transfer failed");

        // Vérifier si l'utilisateur a déjà parié sur ce bet
        UserBet storage existingBet = userBets[msg.sender][betId];
        bool isNewParticipant = existingBet.amount == 0;

        // Mise à jour des pools
        if (team == 1) {
            CurrentBets[betId].team1Pool += amount;
        } else {
            CurrentBets[betId].team2Pool += amount;
        }

        // Enregistrer/mettre à jour le pari de l'utilisateur
        if (isNewParticipant) {
            userBets[msg.sender][betId] = UserBet({
                amount: amount,
                teamChosen: team,
                claimed: false
            });
            // Ajouter à la liste des participants
            betParticipants[betId].push(msg.sender);
        } else {
            // Mise à jour du pari existant
            require(existingBet.teamChosen == team, "Cannot change team choice");
            existingBet.amount += amount;
        }

        emit BetPlaced(msg.sender, betId, team, amount);
    }

    function claimWinnings(uint256 betId) external nonReentrant {
        require(betId < CurrentBets.length, "Invalid bet ID");
        require(CurrentBets[betId].resolved, "Pari non resolu");
        require(CurrentBets[betId].winningTeam != 0, "Equipe gagnante non definie");

        UserBet storage userBet = userBets[msg.sender][betId];
        require(!userBet.claimed, "Gains deja reclames");
        require(userBet.amount > 0, "Aucun pari trouve");
        require(userBet.teamChosen == CurrentBets[betId].winningTeam, "Equipe perdante");
        
        uint256 totalPool = CurrentBets[betId].team1Pool + CurrentBets[betId].team2Pool;
        uint256 winningPool = (CurrentBets[betId].winningTeam == 1) ? CurrentBets[betId].team1Pool : CurrentBets[betId].team2Pool;
        
        require(winningPool > 0, "No winning pool");
        
        uint256 payout = (userBet.amount * totalPool) / winningPool;
        
        userBet.claimed = true;
        
        // Transfert des tokens ERC20 vers le gagnant
        require(bettingToken.transfer(msg.sender, payout), "Transfer failed");
        
        emit WinningsClaimed(msg.sender, betId, payout);
    }

    function resolveBet(uint256 betId, uint8 winningTeam) external onlyOracle {
        require(betId < CurrentBets.length, "Invalid bet ID");
        require(!CurrentBets[betId].resolved, "Bet already resolved");
        require(winningTeam == 1 || winningTeam == 2, "Invalid winning team");
        require(block.timestamp >= CurrentBets[betId].deadline, "Bet still active"); // ✅ Vérification deadline
        
        CurrentBets[betId].winningTeam = winningTeam;
        CurrentBets[betId].resolved = true;
        
        emit BetResolved(betId, winningTeam);
    }

    function getBetCount() external view returns (uint256) {
        return CurrentBets.length;
    }

    function getBetParticipants(uint256 betId) external view returns (address[] memory) {
        require(betId < CurrentBets.length, "Invalid bet ID");
        return betParticipants[betId];
    }

    function getUserBet(address user, uint256 betId) external view returns (UserBet memory) {
        return userBets[user][betId];
    }

    function setOracle(address _newOracle) external onlyOwner {
        oracle = _newOracle;
    }

    function emergencyWithdraw(address token, uint256 amount) external onlyOwner {
        IERC20(token).transfer(owner(), amount);
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }
}