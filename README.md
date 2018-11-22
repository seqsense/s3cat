# s3cat
cat from s3 with kms decryption

## install

1. find latest version at [release page](https://github.com/seqsense/s3cat/releases)
2. install (replace OS, Arch, install directory, and package version)
    ```shell
    curl -L -s https://github.com/seqsense/s3cat/releases/download/v0.0.0/s3cat-linux-amd64.tgz | tar xzf - -C ~/.local/bin/
    ```

## usage

Credentials and AWS_REGION setting should be provided in the environment.

```
usage: s3cat [options] bucket s3path
options:
  -decrypt
    	decrypt using kms
```
