# Deploy (Linux)

```bash
USER=root
HOST=5.161.76.175
rsync -avz -e ssh build/dist/* $USER@$HOST:/srv/codon/
rsync -avz -e ssh src/scripts/Caddyfile $USER@$HOST:/etc/caddy/sites/codon
ssh $USER@$HOST -C "systemctl reload caddy"
curl -vik -H "Host: codon.com" http://$HOST/
```

## Caddy

```
http://codon.com {
	file_server
	root * /srv/codon
}
```

## Systemd

```
[Unit]
Description=Caddy
Documentation=https://caddyserver.com/docs/
After=network.target network-online.target
Requires=network-online.target

[Service]
Type=notify
User=caddy
Group=caddy
ExecStart=/usr/bin/caddy run --environ --config /etc/caddy/Caddyfile
ExecReload=/usr/bin/caddy reload --config /etc/caddy/Caddyfile --force
TimeoutStopSec=5s
LimitNOFILE=1048576
PrivateTmp=true
ProtectSystem=full
AmbientCapabilities=CAP_NET_ADMIN CAP_NET_BIND_SERVICE

[Install]
WantedBy=multi-user.target
```
