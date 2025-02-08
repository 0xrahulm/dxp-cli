# `dxpcli ulxly`

> Auto-generated documentation.

## Table of Contents

- [Description](#description)
- [Usage](#usage)
- [Flags](#flags)
- [See Also](#see-also)

## Description

Utilities for interacting with the uLxLy bridge

## Usage

Basic utility commands for interacting with the bridge contracts, bridge services, and generating proofs
## Flags

```bash
  -h, --help   help for ulxly
```

The command also inherits flags from parent commands.

```bash
      --config string   config file (default is $HOME/.dxp-cli.yaml)
      --pretty-logs     Should logs be in pretty format or JSON (default true)
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

- [dxpcli](dxpcli.md) - A Swiss Army knife of blockchain tools.
- [dxpcli ulxly bridge](dxpcli_ulxly_bridge.md) - Commands for moving funds and sending messages from one chain to another

- [dxpcli ulxly claim](dxpcli_ulxly_claim.md) - Commands for claiming deposits on a particular chain

- [dxpcli ulxly claim-everything](dxpcli_ulxly_claim-everything.md) - Attempt to claim as many deposits and messages as possible

- [dxpcli ulxly empty-proof](dxpcli_ulxly_empty-proof.md) - create an empty proof

- [dxpcli ulxly get-deposits](dxpcli_ulxly_get-deposits.md) - Generate ndjson for each bridge deposit over a particular range of blocks

- [dxpcli ulxly proof](dxpcli_ulxly_proof.md) - Generate a proof for a given range of deposits

- [dxpcli ulxly zero-proof](dxpcli_ulxly_zero-proof.md) - create a proof that's filled with zeros

