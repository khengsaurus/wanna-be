### Dev

Important: execute scripts from root dir, i.e.

```bash
# With pwd = root
bash l7/init.sh
```

### Notes

```
# ssh into a container
docker exec -it <container-name> bash
```

### SSL keys

SSL keys can be generated using openssl (for a local environment, use hostname as per /etc/hosts config, i.e.)

```txt
# /etc/hosts
127.0.0.1	<hostname>
```

```bash
# Install openssl via brew
brew install openssl

# Generate keys
openssl req -x509 -sha256 -nodes -newkey rsa:2048 -days 365 -keyout <key_name>.key -out <cert_name>.crt
```

Ensure the key and certificate names are corretly reflected in nginx.conf
