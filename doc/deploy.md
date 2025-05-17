---
order: 2
section: dev
slug: deploy
title: Deploy
---

# Deploy

Apps are easy to deploy to Linux, either as a static site or a Go binary served
with Caddy.

## Caddy

Install Caddy and configure it to serve multiple sites

```bash
apt-get install caddy

cat <<EOF >/etc/caddy/Caddyfile
import sites/*
EOF

# view logs
journalctl -f -u caddy

# view config
find /etc/caddy

# view sites
ls -al /srv
```

## Go Binary

To deploy a Go binary, cross-compile for Linux, run under Systemd and serve with
Caddy.

```bash
DOMAIN=example.com
HOST=5.161.XX.XXX
PORT=2234
USER=root
SITE=app

go generate ./...
mkdir -p build/app
GOOS=linux GOARCH=amd64 go build -o build/app/app cmd/app/main.go

rsync -avz -e ssh build/app/* $USER@$HOST:/srv/$SITE/

cat <<EOF | ssh $USER@$HOST -T "cat > /etc/systemd/system/$SITE.service; sudo systemctl daemon-reload; sudo systemctl enable $SITE.service; sudo systemctl start $SITE.service"
[Unit]
Description=$SITE service
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/srv/$SITE
ExecStart=/srv/$SITE/app -port $PORT
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

cat <<EOF | ssh $USER@$HOST -T "cat > /etc/caddy/sites/$SITE; systemctl reload caddy"
https://$DOMAIN {
  reverse_proxy localhost:$PORT
}
EOF
```

Review logs with `journalctl -f -u app`.

## Static Site

To deploy a static site, build the directory and serve with Caddy.

```bash
DOMAIN=example.com
HOST=5.161.XX.XXX
USER=root
SITE=codon

rm -rf build/dist
bun run build
rsync -avz -e ssh build/dist/* $USER@$HOST:/srv/$SITE

cat <<EOF | ssh $USER@$HOST -T "cat > /etc/caddy/sites/codon; systemctl reload caddy"
https://$SITE.lab.mixable.net {
	file_server
	root * /srv/codon
}
EOF
```
