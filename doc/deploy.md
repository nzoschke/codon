---
order: 2
section: dev
slug: deploy
title: Deploy
---

# Deploy (Linux)

Build a static site and deploy under Caddy

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
