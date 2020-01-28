# ipld-dagcbor-floats

Experiments in DagCBOR float encoding & decoding, how to harmonize JavaScript and Go & testing edge cases.

See https://github.com/ipld/specs/issues/227

```sh
# `make` will set up dependencies too if you have `node` and `go` installed
$ make

Running Node.js ...
in,cid,bytes,out
0.3333333333333333,bafyreige4pbaqdjzwnr42qjvsbydtljrrjw3wuddiaoccxuyap2p4rp2hu,fb3fd5555555555555,0.3333333333333333
1.5,bafyreifwro2f5svqgknlqfo26rhvuawsuenivod7xppuwcf4vyamvwqdeq,f93e00,1.5
-2.001953125,bafyreietuok7wedcqsauyaghvmcecx4z6kq7rweq26uc7mgfemm35pkkd4,f9c001,-2.001953125
-2.000000238418579,bafyreibxlwx2cdzubik3l2vr3f4fkdcfhuqhvpfwy4i3g4soqufcpmdxle,fac0000001,-2.000000238418579
-2.0000000000000004,bafyreiafjiuhmxfdwsqfnmcky2uvvvndocoiogp5ctbe3sjcdfgcx5yae4,fbc000000000000001,-2.0000000000000004
9.5367431640625e-7,bafyreifxfygrminq5xudttic33dhuoytj4r6zfhdugw4htduxwjp43wowi,f90010,9.5367431640625e-7
1.401298464324817e-45,bafyreigcum33q56zunz7czuf2raveyswxv2hkn7ymt5bbrmwmnzxjnboyi,fa00000001,1.401298464324817e-45

Running Go ...
in,cid,bytes,out
0.3333333333333333,bafyreige4pbaqdjzwnr42qjvsbydtljrrjw3wuddiaoccxuyap2p4rp2hu,fb3fd5555555555555,0.3333333333333333
1.5,bafyreib2ir5ittexhu5d3zopo6wzsshuwi6byb3cdtp67bfopa2fkbpfcy,fb3ff8000000000000,1.5
-2.001953125,bafyreihppogvojsegv3nfpiutahmyx27wee6gwneivfcglcemz46qffngi,fbc000040000000000,-2.001953125
-2.000000238418579,bafyreicuwh4aig3vhnfzcmod4o47rz2wfokvvrfypgb3vusz625g5dnml4,fbc000000020000000,-2.000000238418579
-2.0000000000000004,bafyreiafjiuhmxfdwsqfnmcky2uvvvndocoiogp5ctbe3sjcdfgcx5yae4,fbc000000000000001,-2.0000000000000004
9.5367431640625e-07,bafyreicn4nm47fqf622spn5uxrv5moujkel7nt36u2qmjtpjv4hiltzyte,fb3eb0000000000000,9.5367431640625e-07
1.401298464324817e-45,bafyreiajuqdj42qr3e65egrznxuyufaoanhthtastep2lsagskfzckepza,fb36a0000000000000,1.401298464324817e-45
```
