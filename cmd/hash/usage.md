```bash
$ echo -n "hello" > hello.txt
$ dxpcli hash sha1 --file hello.txt
aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
$ echo -n "hello" | dxpcli hash sha1
aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
$ dxpcli hash sha1 hello
aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
```
