# particeps

Command-line utility to upload files to [AnonFiles](https://anonfiles.com/), [BayFiles](https://bayfiles.com/) or [Filebin](https://filebin.net).

```
Usage: ./particeps [-h, --help] [-a, --anonfiles] [-F, --filebin] [-b, --bayfiles] -f, --filename path-to-file
```

## Example:

```shell
>>>  ./particeps -F -f particeps
config folder: /home/vinicius/.config
particeps: file "particeps" has size 3.82 MB
particeps: uploading to https://filebin.com
particeps: successfully uploaded "particeps" to https://filebin.com
particeps: full-length link: https://filebin.net/21xhy0toq0ix3nrf
particeps: bear in mind that Filebin only stores the files for a week.
```

## Build

You can get a stripped, statically linked binary in the releases page.

If Go is installed, you can get Particeps by using:

```
go get github.com/carmesim/particeps
```
