# go-oracle

> Install oracle driver in Macos <go-oci8>

1. Download instant-client basic and sdk here <https://www.oracle.com/database/technologies/instant-client/macos-intel-x86-downloads.html>
2. unzip into Downloads/instantclient_12_2 & Downloads/instantclient_12_2/sdk
3. brew install pkg-config
4. create file oci8.pc into Downloads/instantclient_12_2
5. edit file oci8.pc
```
prefixdir=/Users/<username>/Downloads/instantclient_12_2/
libdir=${prefixdir}
includedir=${prefixdir}/sdk/include

Name: OCI
Description: Oracle database driver
Version: 12.2
Libs: -L${libdir} -lclntsh
Cflags: -I${includedir}
```
6. Set ENV
export LD_LIBRARY_PATH=/Users/<username>/Downloads/instantclient_12_2
export PKG_CONFIG_PATH=/Users/<username>/Downloads/instantclient_12_2

7. go get -u github.com/mattn/go-oci8 