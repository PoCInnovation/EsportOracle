// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "../lib/openzeppelin-contracts/contracts/access/Ownable.sol";
import "../lib/openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol";
import "../lib/openzeppelin-contracts/contracts/utils/Pausable.sol";
import "./esportOracleClientRequester.sol";

contract BetContract is EsportOracleClientRequester, Ownable, ReentrancyGuard, Pausable {
    struct Bet {
        string description;       // Description du pari
        uint256 team1Id;          // ID de la première équipe (Pandascore API)
        uint256 team2Id;          // ID de la deuxième équipe (Pandascore API)
        uint256 deadline;        // Date limite pour parier
        uint256 team1Pool;       // Montant total misé sur équipe 1
        uint256 team2Pool;       // Montant total misé sur équipe 2
        uint8 winningTeam;       // 1 = team1 gagne, 2 = team2 gagne, 0 = pas résolu
        bool resolved;           // Pari résolu ou non
        address creator;         // Créateur du pari
        uint256 matchId;         // ID du match associé au pari
    }

    struct UserBet {
        uint256 amount;
        uint8 teamChosen;        // ✅ Changé de uint256 à uint8 pour cohérence
        bool claimed;
    }

    Bet[] public CurrentBets;
    mapping(address => mapping(uint256 => UserBet)) public userBets;
    mapping(uint256 => address[]) public betParticipants; // ✅ Ajout du mapping manquant
    IERC20 public bettingToken;
    
    // Fee pour demander un match à l'oracle
    uint256 public matchRequestFee = 0.001 ether; // Fee par défaut

    event BetPlaced(address indexed user, uint256 indexed betId, uint8 team, uint256 amount);
    event WinningsClaimed(address indexed user, uint256 indexed betId, uint256 payout);
    event BetResolved(uint256 indexed betId, uint8 winningTeam);
    event BetCreated(uint256 indexed betId, string description, uint256 team1Id, uint256 team2Id, uint256 deadline);

    constructor(address _oracle, address _tokenAddress) EsportOracleClientRequester(_oracle) Ownable(msg.sender) {
        bettingToken = IERC20(_tokenAddress);
    }

    receive() external payable {
        // ETH reçu pour payer les fees oracle
    }

    function createBet(
        string memory description,
        uint256 team1Id,
        uint256 team2Id,
        uint256 deadline,
        uint256 matchId
    ) external {
        require(deadline > block.timestamp, "Deadline must be in the future");

        Bet memory newBet = Bet({
            description: description,
            team1Id: team1Id,
            team2Id: team2Id,
            deadline: deadline,
            team1Pool: 0,
            team2Pool: 0,
            winningTeam: 0,
            resolved: false,
            creator: msg.sender,
            matchId: matchId
        });

        CurrentBets.push(newBet);

        uint256 betId = CurrentBets.length - 1;
        emit BetCreated(betId, description, team1Id, team2Id, deadline);
    }

    function placeBet(uint256 betId, uint8 team, uint256 amount) external whenNotPaused {
        require(betId < CurrentBets.length, "Invalid bet ID"); // ✅ Vérification ajoutée
        require(CurrentBets[betId].deadline > block.timestamp, "Pari expire");
        require(team == 1 || team == 2, "Equipe invalide");
        require(amount > 0, "Montant invalide");

        require(bettingToken.transferFrom(msg.sender, address(this), amount), "Transfer failed");

        UserBet storage existingBet = userBets[msg.sender][betId];
        bool isNewParticipant = existingBet.amount == 0;

        if (team == 1) {
            CurrentBets[betId].team1Pool += amount;
        } else {
            CurrentBets[betId].team2Pool += amount;
        }

        if (isNewParticipant) {
            userBets[msg.sender][betId] = UserBet({
                amount: amount,
                teamChosen: team,
                claimed: false
            });
            betParticipants[betId].push(msg.sender);
        } else {
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

    function callMatchReceived(EsportOracleTypes.Match memory _match) external override onlyOracle
    {
        _callMatchReceivedInternal(_match);
    }

    function _callMatchReceivedInternal(EsportOracleTypes.Match memory _match) internal {
        if (!oracle.isMatchRequested(_match._id)) {
            require(address(this).balance >= matchRequestFee, "Insufficient contract balance for match request");
            this.receiveMatch{value: matchRequestFee}(_match._id);
        }
        
        EsportOracleTypes.MatchRequest memory matchRequest = oracle.getMatchRequest(_match._id);
        require(matchRequest.fulfilled, "Match request not fulfilled");
        require(_match._winnerId != 0, "Match not finished");
    }

    function resolveBet(uint256 betId, EsportOracleTypes.Match memory matchData) external onlyOracle {
        require(betId < CurrentBets.length, "Invalid bet ID");
        Bet storage bet = CurrentBets[betId];
        require(!bet.resolved, "Bet already resolved");
        require(block.timestamp >= bet.deadline, "Bet still active");
        require(bet.matchId == matchData._id, "Match ID mismatch");

        // Appeler la logique de callMatchReceived
        _callMatchReceivedInternal(matchData);

        // Déterminer quelle équipe a gagné en comparant avec les IDs
        uint8 betWinningTeam;
        if (matchData._winnerId == bet.team1Id) {
            betWinningTeam = 1; // Team1 gagne
        } else if (matchData._winnerId == bet.team2Id) {
            betWinningTeam = 2; // Team2 gagne
        } else {
            revert("Winner ID doesn't match any team in this bet");
        }

        bet.winningTeam = betWinningTeam;
        bet.resolved = true;

        emit BetResolved(betId, betWinningTeam);
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

    function getBet(uint256 betId) external view returns (Bet memory) {
        require(betId < CurrentBets.length, "Invalid bet ID");
        return CurrentBets[betId];
    }

    /**
     * @notice Permet de demander un match à l'oracle s'il n'a pas encore été récupéré
     * @param matchId L'ID du match à demander
     */
    function requestMatchIfNeeded(uint256 matchId) external payable {
        if (!oracle.isMatchRequested(matchId)) {
            // Vérifier qu'on a reçu assez d'ETH pour payer les fees
            require(msg.value > 0, "Fee required to request match");
            
            // Déléguer l'appel à la fonction receiveMatch héritée
            this.receiveMatch{value: msg.value}(matchId);
        } else {
            // Si le match est déjà demandé, rembourser l'ETH envoyé
            if (msg.value > 0) {
                payable(msg.sender).transfer(msg.value);
            }
        }
    }

    /**
     * @notice Permet au propriétaire de définir les fees pour demander un match
     * @param _fee Le nouveau montant de fee en wei
     */
    function setMatchRequestFee(uint256 _fee) external onlyOwner {
        matchRequestFee = _fee;
    }

    /**
     * @notice Permet d'ajouter des fonds au contrat pour payer les fees oracle
     */
    function depositForFees() external payable {
        require(msg.value > 0, "Must send ETH");
    }

    /**
     * @notice Permet au propriétaire de retirer l'ETH du contrat
     * @param amount Montant à retirer en wei
     */
    function withdrawETH(uint256 amount) external onlyOwner {
        require(address(this).balance >= amount, "Insufficient balance");
        payable(owner()).transfer(amount);
    }

    /**
     * @notice Obtenir le solde ETH du contrat
     */
    function getContractETHBalance() external view returns (uint256) {
        return address(this).balance;
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