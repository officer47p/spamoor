// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.19;

contract LoadTest {
    // State variables to test SLOAD operations
    uint256 public testValue;
    mapping(uint256 => uint256) public testMapping;
    uint256[100] public testArray;

    // Counter for number of operations performed
    uint256 public operationCount;

    event DataReceived(bytes data);
    event SloadPerformed(uint256 operations, uint256 gasUsed);

    constructor() {
        // Initialize some values for testing
        testValue = 42;
        for (uint256 i = 0; i < 100; i++) {
            testMapping[i] = i * 2;
            testArray[i] = i * 3;
        }
    }

    function processCallData(bytes memory data) public {
        emit DataReceived(data);
    }

    // Function to perform multiple SLOAD operations with simple value
    function testSimpleSload(uint256 iterations) external returns (uint256) {
        uint256 gasStart = gasleft();
        uint256 temp;

        for (uint256 i = 0; i < iterations; i++) {
            temp = testValue; // SLOAD operation
        }

        uint256 gasUsed = gasStart - gasleft();
        operationCount += iterations;
        emit SloadPerformed(iterations, gasUsed);

        return temp;
    }

    // Function to test SLOAD operations with mapping
    function testMappingSload(uint256 iterations) external returns (uint256) {
        uint256 gasStart = gasleft();
        uint256 temp;

        for (uint256 i = 0; i < iterations; i++) {
            temp = testMapping[i % 100]; // SLOAD operation from mapping
        }

        uint256 gasUsed = gasStart - gasleft();
        operationCount += iterations;
        emit SloadPerformed(iterations, gasUsed);

        return temp;
    }

    // Function to test SLOAD operations with array
    function testArraySload(uint256 iterations) external returns (uint256) {
        uint256 gasStart = gasleft();
        uint256 temp;

        for (uint256 i = 0; i < iterations; i++) {
            temp = testArray[i % 100]; // SLOAD operation from array
        }

        uint256 gasUsed = gasStart - gasleft();
        operationCount += iterations;
        emit SloadPerformed(iterations, gasUsed);

        return temp;
    }

    // Combined test of all storage types
    function testCombinedSload(uint256 iterations) external returns (uint256) {
        uint256 gasStart = gasleft();
        uint256 temp;

        for (uint256 i = 0; i < iterations; i++) {
            temp = testValue; // Simple variable SLOAD
            temp += testMapping[i % 100]; // Mapping SLOAD
            temp += testArray[i % 100]; // Array SLOAD
        }

        uint256 gasUsed = gasStart - gasleft();
        operationCount += iterations * 3; // Three SLOADs per iteration
        emit SloadPerformed(iterations * 3, gasUsed);

        return temp;
    }

    // Function to reset operation counter
    function resetCounter() external {
        operationCount = 0;
    }
}
