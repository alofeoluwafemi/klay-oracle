## Actions

- Filter Invalid values from external API response
- Set Adapter to have MAX_DEVIATION for Price Feed
- Add package **Initiator** that has HEART_BEAT variable to trigger Oracle request for one or more Adapter, to allow regular price feed update
- Allow option to use random-aggregator, median-aggregator or volume-weighted-aggregator to get aggregated value
- Add package **api** to return oracle data for wallet, web application usage
- Research Tokenomics
  - Native
    - Node Operators earn token, difficulty increases until allocation is exhausted
    - Consumer pays for Oracle request
  - DAO
    - Vote data feed to be added
    - Vote platform proposal changes
  - Import gitbook documentation into platform
- Sign messages sent by Node to Oracle with private key and confirm signature
- Write test for components and packages
- Setup CI/CD and Dockerize the Node