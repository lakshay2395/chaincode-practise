# SCHOOLS NETWORK

## Introduction

Proposed hyperledger fabric network aims to create a decentralized network between schools all over india and Central Board of Secondary Education (CBSE) to maintain a registry of all the students admitted and easing the process of student migration to one school from another.

## Basic Network Layout

Proposed network consists of following features :
-   Two schools, SchoolOneOrg and SchoolTwoOrg as the initial members of the consortium.
-   CBSEOrg, representing CBSE.
-   Also, National Informatics Centre (NIC) is proposed to be the orderer node of the network.
-   Each school will have its private channel with CBSE to ensure privacy of its students' records.


## Basic Permission Heirarchy

- Only schools can add/delete/update student details.
- CBSE can centrally receive applications from students for migration/leave requests.
- CBSE can then issue a new school leaving/migration request to the specific school on its specific channel.
- Schools can either reject/accept on the basis of various factors like mid term transfer requests, military personnel cases ,etc.
- The student entry, on acceptance, will be created on the separate channel for separate school where student wants to migrate.
- There the same acceptance/rejection thing continues.
- By having shared ledgers, we can keep a track of all the phases of rejection and acceptance.


