// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

contract EOASpammer {
    function spamFund(uint256 iterations) external payable {
        bytes32 currentHash = keccak256(abi.encodePacked(block.timestamp));
        
        for (uint256 i = 0; i < iterations;) {
            address payable target = payable(address(uint160(uint256(currentHash))));
            
            assembly {
                let success := call(gas(), target, 1, 0, 0, 0, 0)
            }
            
            currentHash = keccak256(abi.encodePacked(currentHash));
            
            unchecked {
                ++i;
            }
        }
        
        if (msg.value > iterations) {
            payable(msg.sender).transfer(msg.value - iterations);
        }
    }
    
    receive() external payable {}
}