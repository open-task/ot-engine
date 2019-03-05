package contracts

// OpenTaskABI is the input ABI used to generate the binding from.
const OpenTaskABI = `
[
	{
		"constant": false,
		"inputs": [
			{
				"name": "solutionId",
				"type": "string"
			}
		],
		"name": "accept",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "solutionId",
				"type": "string"
			}
		],
		"name": "reject",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "newDETAddress",
				"type": "address"
			}
		],
		"name": "setDETAddress",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "missionId",
				"type": "string"
			}
		],
		"name": "getMission",
		"outputs": [
			{
				"name": "rewardInWei",
				"type": "uint256"
			},
			{
				"name": "data",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [],
		"name": "renounceOwnership",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "getDETAddress",
		"outputs": [
			{
				"name": "",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"name": "",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "isOwner",
		"outputs": [
			{
				"name": "",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "solutionId",
				"type": "string"
			},
			{
				"name": "arbitrationId",
				"type": "string"
			}
		],
		"name": "confirm",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "missionId",
				"type": "string"
			},
			{
				"name": "rewardInWei",
				"type": "uint256"
			},
			{
				"name": "data",
				"type": "string"
			}
		],
		"name": "publish",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "solutionId",
				"type": "string"
			},
			{
				"name": "missionId",
				"type": "string"
			},
			{
				"name": "data",
				"type": "string"
			}
		],
		"name": "solve",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"name": "initialDETAddress",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "newDETAddress",
				"type": "address"
			}
		],
		"name": "SetDETAddress",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "missionId",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "rewardInWei",
				"type": "uint256"
			},
			{
				"indexed": false,
				"name": "data",
				"type": "string"
			}
		],
		"name": "Publish",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "solutionId",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "missionId",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "data",
				"type": "string"
			}
		],
		"name": "Solve",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "solutionId",
				"type": "string"
			}
		],
		"name": "Accept",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "solutionId",
				"type": "string"
			}
		],
		"name": "Reject",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "solutionId",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "arbitrationId",
				"type": "string"
			}
		],
		"name": "Confirm",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "OwnershipTransferred",
		"type": "event"
	}
]`