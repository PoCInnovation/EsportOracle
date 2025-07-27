// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/EO.sol";

contract EOTest is Test {
    EO public token;
    address public owner;
    address public user1;
    address public user2;
    address public authorizedContract;
    
    uint256 public constant TOTAL_CAP = 100000000;
    uint256 public constant OWNER_TOKENS = 20000000;
    uint256 public constant TOKEN_SIZE = 10 ** 18;

    /**
     * @notice Permet de créer des contracts fictives
     * @dev Seulement le owner du contract
     */
    function setUp() public {
        owner = address(this); // le contract
        user1 = makeAddr("user1"); // fictive 1
        user2 = makeAddr("user2"); // fictive 2
        authorizedContract = makeAddr("authorizedContract");

        token = new EO(TOTAL_CAP);
    }
    
    /**
     * @notice Vérifications des valeurs pour le token
     * @dev Seulement le owner du contract
     */
    function test_Deployment() view public {
        //Vérification des valeurs pour le token
        assertEq(token.name(), "EsportOracleToken");
        assertEq(token.symbol(), "EO");
        assertEq(token.decimals(), 18);
        assertEq(token.cap(), TOTAL_CAP * TOKEN_SIZE);
        assertEq(token.owner(), owner);
        assertEq(token.totalSupply(), OWNER_TOKENS * TOKEN_SIZE);
        assertEq(token.balanceOf(owner), OWNER_TOKENS * TOKEN_SIZE);
    }
    
    /**
     * @notice Vérifie que les transactions marchent bien
     * @dev Seulement le owner du contract
     */
    function test_Transfer() public {
        //Vérifier que les transactions marchent bien
        uint256 amount = 1000 * TOKEN_SIZE;

        token.transfer(user1, amount);
        
        assertEq(token.balanceOf(user1), amount);
        assertEq(token.balanceOf(owner), (OWNER_TOKENS * TOKEN_SIZE) - amount);
    }

    /**
     * @notice Vérifie que les transactions marchent bien
     * @dev Seulement le owner du contract
     */
    function test_TransferFrom() public {
        uint256 amount = 1000 * TOKEN_SIZE;

        token.approve(user1, amount);
        

        vm.prank(user1);
        token.transferFrom(owner, user2, amount);
        
        assertEq(token.balanceOf(user2), amount);
        assertEq(token.allowance(owner, user1), 0);
    }

    function test_Approve() public {
        uint256 amount = 1000 * TOKEN_SIZE;
        
        token.approve(user1, amount);
        assertEq(token.allowance(owner, user1), amount);
    }
    
    /**
     * @notice Vérifie que le owner peut bien paused et s'attend à une erreur.
     * @dev Seulement le owner du contract
     */
    function test_OnlyOwnerCanPause() public {
        token.pause();
        assertTrue(token.paused()); //Return true si le token est paused

        vm.prank(user1);
        vm.expectRevert();
        token.pause();
    }

    /**
     * @notice Vérifie que le owner ne peut pas paused et s'attend que il doit échouer.
     * @dev Seulement le owner du contract
     */
    function test_OnlyOwnerCanUnpause() public {
        token.pause();
        assertTrue(token.paused());
        

        token.unpause();
        assertFalse(token.paused());

        token.pause();
        vm.prank(user1);
        vm.expectRevert();
        token.unpause();
    }
    
    /**
     * @notice Vérifie que la pause bloque bien les transfers.
     * @dev Seulement le owner du contract
     */
    function test_TransferWhenPaused() public {
        uint256 amount = 1000 * TOKEN_SIZE;

        token.pause();

        vm.expectRevert();
        token.transfer(user1, amount);
    }

    /**
     * @notice Vérifie que on ne peut pas mint pendant une pause.
     * @dev Seulement le owner du contract
     */
    function test_MintWhenPaused() public {
        uint256 amount = 1000 * TOKEN_SIZE;
    
        token.pause();

        vm.expectRevert();
        token.mintTokens(user1, amount);
    }
    
    function test_MintTokens() public {
        uint256 amount = 1000 * TOKEN_SIZE;
        uint256 initialSupply = token.totalSupply();
        
        token.mintTokens(user1, amount);
        
        assertEq(token.balanceOf(user1), amount);
        assertEq(token.totalSupply(), initialSupply + amount);
    }

    function test_OnlyOwnerCanMint() public {
        uint256 amount = 1000 * TOKEN_SIZE;
        
        vm.prank(user1);
        vm.expectRevert();
        token.mintTokens(user2, amount);
    }

    function test_CannotMintBeyondCap() public {
        uint256 currentSupply = token.totalSupply();
        uint256 cap = token.cap();
        uint256 remainingMintable = cap - currentSupply;
        
        token.mintTokens(user1, remainingMintable);
        assertEq(token.totalSupply(), cap);
        
        vm.expectRevert();
        token.mintTokens(user1, 1);
    }
    
    function test_SetAuthorizedContract() public {
        assertFalse(token.AuthorizedContract(authorizedContract));
        
        vm.expectEmit(true, false, false, true);
        emit EO.ContractAuthorized(authorizedContract, true);
        token.setAuthorizedContract(authorizedContract, true);
        
        assertTrue(token.AuthorizedContract(authorizedContract));
        
        vm.expectEmit(true, false, false, true);
        emit EO.ContractAuthorized(authorizedContract, false);
        token.setAuthorizedContract(authorizedContract, false);
        
        assertFalse(token.AuthorizedContract(authorizedContract));
    }

    function test_OnlyOwnerCanSetAuthorizedContract() public {
        vm.prank(user1);
        vm.expectRevert();
        token.setAuthorizedContract(authorizedContract, true);
    }
    
    /**
     * @notice Vérifie que les bons evenements sont émis.
     * @dev Seulement le owner du contract
     */
    function test_TokensMintedEvent() public {
        uint256 amount = 1000 * TOKEN_SIZE;
        
        vm.expectEmit(true, false, false, true);
        emit EO.TokensMinted(user1, amount);
        token.mintTokens(user1, amount);
    }
}