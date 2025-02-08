The create subcommand will create a new key pair. By default, a hex private key will be written to `stdout`.
```bash
dxpcli signer create > private-key.txt
```

If you need to work with a go-ethereum style keystore, a key can be added by setting a `--keystore` directory. When you run this command, you'll need to specify a password to encrypt the private key.

```bash
dxpcli signer create --keystore /tmp/keystore
```

Dxpcli also has basic support for KMS with GCP. Creating a new key in the cloud can be accomplished with a command like this

```bash
# dxpcli assumes that there is default login that's been done already
gcloud auth application-default login
dxpcli signer create --kms GCP --gcp-project-id prj-polygonlabs-devtools-dev --key-id jhilliard-trash
```
