pragma solidity ^0.5.0;

contract IPrice {
    function post() external;

    function getWithError() external view returns (uint256, bool, bool);

}
