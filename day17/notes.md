## Notes

### First cycle
- Round **1**, **475th** direction
- Contains **1767** settled rocks
- **ReachedY** equals **2756**
- **565930956** full cycles can occur in the entire run
- **999999999252** rocks would be settled in full cycles
- **748** rockes would be left to calculate

### NumMemoized options:
- 100 - **1566272189352** - **correct**
- 200 - **1566272189346** - too low
- 50 - **1566272189349** - too low
- 20 - **1566272189357** - too high

For some reason, the result slightly differ (+/- 10) when using different memoization options.
This may be due to the key generation itself, or some other bug in the algorithm.
