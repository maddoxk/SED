# SED
**SED** - Simple Encrypt and Decryption

**SED** is a **lightweight tool used for cryptography needs.**

    TODO:
    	- Finish decryption handler
		- Translate to GO API
		- Format output to JSON
		- Option to output to file
		
		(Maybe) - Set up webserver API

**Usage:**

    Usage: ./sed <command>
    Commands:
            Cryptography Utility:
                    --encrypt <data>
                            Encrypts data and outputs keys
                    -t <encryption algorithm>
                            Specify encryption algorithm
                                    aes / des
                    -l <len>
                            Sets the key length for generating the key
                    --encode <hex/base64>
                            Encodes the data with base64 (Use if data is too large)
                    --decrypt <data>
                            Decrypts data with RSA (Requires private key)
                    --key <key>
                            Sets the private key for crypting
                    --decode
                            Decodes the data with base64 (Use if data is encoded)
                    -o <file/dir>
                            Sets the output file or directory for the data and keys

**Example**

![enter image description here](https://cdn.discordapp.com/attachments/762193692827189268/905163711277326406/Screen_Shot_2021-11-02_at_11.37.21_AM.png)
