# `dxpcli signer create`

> Auto-generated documentation.

## Table of Contents

- [Description](#description)
- [Usage](#usage)
- [Flags](#flags)
- [See Also](#see-also)

## Description

Create a new key

```bash
dxpcli signer create [flags]
```

## Usage

The create subcommand will create a new key pair. By default, a hex private key will be written to `stdout`.
```bash
dxpcli signer create > private-key.txt
```

If you need to work with a go-ethereum style keystore, a key can be added by setting a `--keystore` directory. When you run this command, you'll need to specify a password to encrypt the private key.

```bash
dxpcli signer create --keystore /tmp/keystore
```

dxpcli also has basic support for KMS with GCP. Creating a new key in the cloud can be accomplished with a command like this

```bash
# dxpcli assumes that there is default login that's been done already
gcloud auth application-default login
dxpcli signer create --kms GCP --gcp-project-id prj-polygonlabs-devtools-dev --key-id jhilliard-trash
```

## Flags

```bash
  -h, --help   help for create
```

The command also inherits flags from parent commands.

```bash
      --chain-id uint              The chain id for the transactions.
      --config string              config file (default is $HOME/.dxp-cli.yaml)
      --data-file string           File name holding data to be signed
      --gcp-import-job-id string   The GCP Import Job ID to use when importing a key
      --gcp-key-version int        The GCP crypto key version to use (default 1)
      --gcp-keyring-id string      The GCP Keyring ID to be used (default "dxpcli-keyring")
      --gcp-location string        The GCP Region to use (default "europe-west2")
      --gcp-project-id string      The GCP Project ID to use
      --key-id string              The id of the key to be used for signing
      --keystore string            Use the keystore in the given folder or file
      --kms string                 AWS or GCP if the key is stored in the cloud
      --pretty-logs                Should logs be in pretty format or JSON (default true)
      --private-key string         Use the provided hex encoded private key
      --type string                The type of signer to use: latest, cancun, london, eip2930, eip155 (default "london")
      --unsafe-password string     A non-interactively specified password for unlocking the keystore
  -v, --verbosity int              0 - Silent
                                   100 Panic
                                   200 Fatal
                                   300 Error
                                   400 Warning
                                   500 Info
                                   600 Debug
                                   700 Trace (default 500)
```

## See also

- [dxpcli signer](dxpcli_signer.md) - Utilities for security signing transactions
