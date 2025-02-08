# `dxpcli`

> Auto-generated documentation.

## Table of Contents

- [Description](#description)
- [Usage](#usage)
- [Flags](#flags)
- [See Also](#see-also)

## Description

A Swiss Army knife of blockchain tools.

## Usage

dxpcli is a collection of tools that are meant to be useful while building, testing, and running block chain applications.
## Flags

```bash
      --config string   config file (default is $HOME/.dxp-cli.yaml)
  -h, --help            help for dxpcli
      --pretty-logs     Should logs be in pretty format or JSON (default true)
  -t, --toggle          Help message for toggle
  -v, --verbosity int   0 - Silent
                        100 Panic
                        200 Fatal
                        300 Error
                        400 Warning
                        500 Info
                        600 Debug
                        700 Trace (default 500)
```

## See also

- [dxpcli abi](dxpcli_abi.md) - Provides encoding and decoding functionalities with contract signatures and ABI.

- [dxpcli dbbench](dxpcli_dbbench.md) - Perform a level/pebble db benchmark

- [dxpcli dumpblocks](dxpcli_dumpblocks.md) - Export a range of blocks from a JSON-RPC endpoint.

- [dxpcli ecrecover](dxpcli_ecrecover.md) - Recovers and returns the public key of the signature

- [dxpcli enr](dxpcli_enr.md) - Convert between ENR and Enode format

- [dxpcli fork](dxpcli_fork.md) - Take a forked block and walk up the chain to do analysis.

- [dxpcli fund](dxpcli_fund.md) - Bulk fund crypto wallets automatically.

- [dxpcli hash](dxpcli_hash.md) - Provide common crypto hashing functions.

- [dxpcli loadtest](dxpcli_loadtest.md) - Run a generic load test against an Eth/EVM style JSON-RPC endpoint.

- [dxpcli metrics-to-dash](dxpcli_metrics-to-dash.md) - Create a dashboard from an Openmetrics / Prometheus response.

- [dxpcli mnemonic](dxpcli_mnemonic.md) - Generate a BIP39 mnemonic seed.

- [dxpcli monitor](dxpcli_monitor.md) - Monitor blocks using a JSON-RPC endpoint.

- [dxpcli nodekey](dxpcli_nodekey.md) - Generate node keys for different blockchain clients and protocols.

- [dxpcli p2p](dxpcli_p2p.md) - Set of commands related to devp2p.

- [dxpcli parseethwallet](dxpcli_parseethwallet.md) - Extract the private key from an eth wallet.

- [dxpcli retest](dxpcli_retest.md) - Convert the standard ETH test fillers into something to be replayed against an RPC

- [dxpcli rpcfuzz](dxpcli_rpcfuzz.md) - Continually run a variety of RPC calls and fuzzers.

- [dxpcli signer](dxpcli_signer.md) - Utilities for security signing transactions

- [dxpcli ulxly](dxpcli_ulxly.md) - Utilities for interacting with the uLxLy bridge

- [dxpcli version](dxpcli_version.md) - Get the current version of this application

- [dxpcli wallet](dxpcli_wallet.md) - Create or inspect BIP39(ish) wallets.

- [dxpcli wrap-contract](dxpcli_wrap-contract.md) - Wrap deployed bytecode into create bytecode.

