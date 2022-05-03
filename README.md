# tfvergen
.terraform-version generator for tfenv

tfvergen returns version by parsing `required_version` from tf files.

## Install

Please download binary from releases page.
Move `tfvergen` to your PATH.

## Usage

```bash
$ cd <your terraform directory (*)>

$ cat config.tf
terraform {
  required_version = "= 1.0.7"
}
  : (omitted)

$ tfvergen > .terraform-version
$ cat .terraform-version
1.0.7
```

(*): a directory where you run `terraform plan/apply`

## NOTE

### Create .terraform-version under the directory recursively

```bash
$ find . -type f -name 'config.tf' | xargs -L1 bash -c 'cd $(dirname $1) && tfvergen > .terraform-version' _
```

### Behavior when min-required value is given

If the following values are specified, tfvergen returns the first occurrence.

```hcl
required_version = ">= 1.0.7, < 1.1.0"
# -> tfvergen returns 1.0.7
```

This is identical to the [tfenv specification](https://github.com/tfutils/tfenv/blob/7e89520f4bdbacb5861aca209f0b8f89271287e1/README.md#min-required).
