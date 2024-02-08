# cpwg
Excercises from Cutajar's "Cuncurrent Programming with Go"

## How to initialize a sub-module with `cobra-cli`

EG, for the [directory chp2](./chp2/):
```sh
mkdir chp2
cd chp2
cobra-cli init
cobra-cli add catrand
cobra-cli add grepfiles
cobra-cli add grepdir
cobra-cli add greprec
```