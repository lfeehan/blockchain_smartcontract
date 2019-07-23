pragma solidity ^0.4.3;

// Simple test smart contract
contract Inbox {

    string public message;

    function Inbox(string initialMessage) public {
        message = initialMessage;
    }

    function setMessage(string newMessage) public {
        message = newMessage;
    }
    
}