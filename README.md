# s3cat
cat from s3 with kms decryption

## install

```shell
curl -L -s https://github.com/seqsense/s3cat/releases/download/v0.0.0/s3cat-linux-amd64.tgz | tar xzf - -C ~/.local/bin/
```
(replace OS, Arch, install directory, and package version)

## usage

Credentials and AWS_REGION setting should be provided in the environment.

```
usage: s3cat [options] bucket s3path
options:
  -decrypt
    	decrypt using kms
```
