# General Notes

## Changes suggested in call 1:
1. See bank anchors and if possible, use their tethered assets for INVTokens instead of doing one ourselves (docs/anchors.md)
2. Split the PBToken idea into its own model and develop a model without the payback token in question
3. Add new entities - developers, contractors,  originators and provisions for them to do their assigned roles and duties.
4. Assume two models overall -
 - All banks are anchors on stellar - we can trade in tethered assets and will not face a problem technically
 - Banks involved with entities are not anchors and we need to come up with solutions to accommodate them in our model
5. Separate the role of an issuer and a platform. An issuer would be someone like neighbourly or swytch whereas the platform is something that we'd be developing upon

Observations:
 - Might be difficult to get banks as anchors since there seem to be only 4 popular ones around

## Discussion in Call 2:
1. Figure out how to incorporate investor USD into a stablecoin on Stellar
2. Discuss with Neighborly regarding their platform's integration into the whole data flow
  - Neighbourly could read the pi data either from the commitment made by Swytch's ERC 721 token on chain or they could read the data from Atonomi's MQTT data stream di rectly
3. Discuss integrating swytch into our process workflow
  - Test out swytch's staging platform, checkout their API endpoints
  - we can either listen to the mqtt service directly or read the ERC 721 details on chain

## Discussion in Call 3:
1. Modify some parts of the contract so that we use a bigger structure named User which can be imported where needed.
2. Change naming scheme of some parts to be more user friendly.

## Discussion in call 4:

Need to have stages for originator proposed contracts and contracts
pre-certified projects, documents need to be attached as well, people need
to be able to see stages at which the project is at - stage 1,2,3 and then it is Live

Kind of need provisions for initial seed that comes from a recipient / originator or an investor
you use this documentation and then put it out for tender (a solar engineer)
seed funding should come from the recipient.

Orders -> Projects -> Stages (can not invest until it arrives a particular stage)

munibond model that has a contract, tweak the model of that model to suit our needs
smart contract needs to point to a specific legal contract, stored on ipfs for eg
munibond vs equity crowdfunding vs opportunity zones

platform of platforms or platform of different contracts themselves

originators need consent before proposing a project, no consent, don't continue, propose RfPs
possibility to raise funds before contract becomes final, differences between munibonds
and equity crowdfunding, undershot vs overshot problems wrt crowdfunding - Matt Maroney, Jase

2 things to consider right now: need to have some form of registry, as payments are made,
ownership of solar systems change, we need to automate that via smart contracts, must be
able to query ownership percent of a solar system or, parties involved must
be able to check this. You could also consider escrow model, track ownership

start working on the base contract: per kwh that is read from the iot device, need to
find where we import the info from the iot device (atonomi), charge the recipient to
pay as per the tariff, similar to the electriicty bill and that changes ownership
as well, see breach scenarios (eth smart contract) and shift between funds, trigger hardware to export
to grid directly

have two contracts - investors and platforms, contractors' access for funds

## TODOs copied from the Ethereum Smart Contract

### TODO TENDER PROCESS & REQUEST FOR PROPOSALS ###
Once an Originator initiates a proposed solar System and Deployment, it needs to be put out for a RFP (Request for Proposal) from contractors/solar developers.
The call for RFP should receive an engineering proposal (i.e. not an engineering blueprint level but a general system architecture level), a quote for materials and labor, a deployment plan.

Note: The tender, review and selection process needs to be flexible to cater for different project modalities (eg. a Public tender vs. a private project)

### TODO PROCESS OF PRE-VERIFICATION BEFORE IT GETS CONFIRMED
Consider payment before setting the system live.

CONFIRMATION
agreement between contractor, participant, and investor
Adds the actual numbers and variables, details to the struct proposed Deployments

### TODO SOLAR ENGINEERING DOCUMENTS
The developers must present the blueprints of the proposed work before working on it. These must be stored in IPFS and hashed to the contract.
These documents will have to get updated at the end of the install and are saved as the blueprint of the 'installed system'

### TODO VERIFICATION OF INSTALMENT & SENSORS
The signature of a 3rd party verifier should be considered as part of onboarding a system and its data.
The verifier confirms system was built according to plan, is compliant with regulation, and has the appropriate working sensors and associated public keys.
This digital signature will allow the sensor data and oracle to commit payment transactions and REC minting.

### TODO CONTRACTOR PAYMENTS
Allow contractor to collect payout on system, once its confirmed

### TODO Consider payouts to be based on instalments throughout the buildout process and with energy generation data
Eg. the contractors receives an upfront payment before the process begins but once the deployment is confirmed, and receives another payment once the witness sensors shows the first generation data. Final payment to contractor should occur once the system shows steady data and behavior after a selected period of time. It also requires a 3rd party verification of system engineering and metering sensors.

### TODO Renewable Energy certification

### TODO Define the payment cycle
Consider making payments every two weeks or every month

### TODO LEGAL OWNERSHIP
Come up with a step so that when ownership is fully transferred, there is an automatic report that can change a registry that has legal validity
functions that require no gas, but will check the state of the system

Returns true if consumer has completely paid off any outstanding balance on the panel at ssAddress within their consumerBuffer period

### TODO BREACH SCENARIOS
Add all the situations and scenarios that need to be considered if the payments are not made or the wallet accounts have insufficient funds
Consider sending email notifications, bringing in the guarantors, activating hardware etc.
