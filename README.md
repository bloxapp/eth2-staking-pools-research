# ETH 2.0 Decentralized Staking Pools - Research
[![blox.io](https://s3.us-east-2.amazonaws.com/app-files.blox.io/static/media/powered_by.png)](https://blox.io)


This repo aims to have in one place all the research around decentralized staking pools for eth 2.0.

- Distribuited key generation and redistribuition
- Rewards/ penalties
- Consensus
	- create new pools
	- [coordinate rotation and assignment of participants between pools](https://github.com/bloxapp/eth2-staking-pools-research/blob/master/pool_rotation.md)
	- [coordinate the execution of the validator's epoch duties by the pool](https://github.com/bloxapp/eth2-staking-pools-research/blob/master/pool_duties.md)
	- coordinate participants exit from the protocol
- [bilinear pairings and BLS12-381 keys](https://github.com/bloxapp/eth2-staking-pools-research/blob/master/BLS_keys_and_pairings.pdf)
- Networking  

### Overview
The backbone of decentralized staking pools is in distribuiting the control of the keys that control the validator and its withdrawal key. You can think of it as a giant multisig setup with some M-of-N threshold for signing attestations, block proposals and withddrawal transactions.
A good starting point could be [this](https://www.youtube.com/watch?v=Jtz9b7yWbLo) presentation.

If we add a consensus protocol that rewards and punishes pool participants, controls withdrawal and onboarding then we have a full protocol for an open decentralzied staking pools network.

One issue that arises immediatley is the security around the oboarding process, how can we guarantee that a formed pool will includee (worst case) no more that 1/3 malicious participants?
Following the Binomial distribuition we can calculate how big does a pool needs to be, this is similar to ethereum committee selection as explained [here](https://notes.ethereum.org/@vbuterin/rkhCgQteN?type=view#Why-32-ETH-validator-sizes). 

Another issue is how big the actual set of participants to select from needs to be, if it's too small a malicious participant can hijack an entire pool, **example**: a pool consists of 60 participants but the available set of participants is only of 19. A malicious opponent can quickly add 39 of his own participants and kidnap the pool. 
This scenario is very possible as there is no guarantee that a large set of non allocated participants will exist. We can force the protocol to not create new pools as long as the non allocated participants set is below a certain number bu that will lead to bad user experience and the protocol's ability to move forward.

A solution to this porblem is to use all the participants (allocated and not) as the set from which pool participants are choosen from.  
A way this could work is a continous rotation of all participants between the pools such that new created pools can share the set of "rotating pariticipants" to create the necessary randomness in allocating new pools.

This repository aims to explore all those challenges (and others) twards a foraml protocol definition.

