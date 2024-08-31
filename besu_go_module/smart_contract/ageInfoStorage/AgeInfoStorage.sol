// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.1 <0.9.0;

contract AgeInfoStorage{
    mapping(string=>uint256) private _ageMap;
    string[] private _nameList;
    
    
    function setAge(string memory key,uint256 age) external {
        _ageMap[key] = age;
    }
    
    function getAge(string memory key) external view returns(uint256){
        return _ageMap[key];
    }
    
    function add(string memory name) external {
        _nameList.push(name);
    }
    
    function getNameList() external view returns(string[] memory ){
        return _nameList;
    }
    
}