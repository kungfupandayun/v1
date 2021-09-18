# Protoc setup

In order to perform code generation, you will need to install Protoc on your computer.

## ============ MacOSX =============

It is actually very easy, open a command line interface and type `brew install protobuf`

## ============ Ubuntu (Linux) ============
Find the correct protocol buffers version based on your Linux Distro: [https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases)

Example with x64:

```sh
# Download the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip

# Unzip the archive
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3

# Move the protoc binary to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/

# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/

# Optional: change owner (replace [user] with your username)
sudo chown [user] /usr/local/bin/protoc
sudo chown -R [user] /usr/local/include/google
```

## ============ Windows ============

- Download the latest archive corresponding to your Windows: [https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases)

-  Extract all to `C:\proto3`. Your directory structure should now be:

```
C:\proto3\bin
C:\proto3\include
```

-  Finally, add `C:\proto3\bin` to your PATH:
    1. Open the Start Search, select **Edit the system environment variables**
    2. Click the Environment Variables button.
    3. Finally, in the Environment Variables window, highlight the Path variable in the Systems Variable section and click the Edit button. Add or modify the path lines with the paths you wish the computer to access.
    4. For more details see [Editing the Path variable on Windows](https://www.computerhope.com/issues/ch000549.htm)
