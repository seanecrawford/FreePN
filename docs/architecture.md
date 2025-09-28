# Architecture

This document provides a high‑level overview of how FreePN is
structured.  It is intentionally concise; see
`docs/security-baseline.md` for protocol details and
`docs/threat-model.md` for adversarial assumptions.

## Components

### Client applications

Clients run on users’ devices and establish a secure tunnel to a
gateway.  The default transport is **WireGuard** running over UDP on
port 51820.  If UDP or WireGuard is blocked, the client may fall back
to **MASQUE** (HTTP/3 over QUIC) using `CONNECT-UDP`/`CONNECT-IP` on
port 443, or further to **OpenVPN** on TLS 1.3 and, as a last resort,
an **IKEv2/IPsec** profile.  Platform‑specific code lives under
`clients/` (see the README in that directory).

Clients also include a **kill switch** implementation that prevents
traffic from leaving the host outside the VPN tunnel.  DNS queries are
sent over the tunnel to an encrypted resolver (DoT/DoH) specified by
the control plane.

### Control plane

The control plane is a thin server that authenticates clients (via
tokens or privacy‑pass proofs), allocates tunnel IP addresses, and
programs gateway nodes with new peers.  It uses the
[`wgctrl`](https://pkg.go.dev/golang.zx2c4.com/wireguard/wgctrl) Go
library to add or remove peers from the WireGuard interface on the
gateway.  Future versions may support multiple gateways and dynamic
sharding of user sessions.

### Gateway nodes

Gateways are Linux machines that terminate WireGuard (and fallback
protocols) and forward packets to the wider Internet.  They enforce a
strict egress policy using nftables: only established flows and
explicitly permitted protocols are allowed; SMTP, NetBIOS and other
abuse‑prone ports are blocked.  Gateways also proxy DNS to an
encrypted resolver and perform NAT for client IPs.

Gateways should not persist per‑user logs.  Aggregate statistics may
be collected for capacity planning, but user activity is never
recorded.

### Resolvers

FreePN runs its own DNS resolvers or forwards queries to trusted
third‑party resolvers over **DNS over TLS (DoT)** or **DNS over HTTPS
(DoH)**.  QNAME minimisation is enabled to reduce leaked metadata.

### Fallback hierarchy

The choice of transport follows this order:

1. **WireGuard/UDP** – fast, modern, and with built‑in key rotation.
2. **MASQUE** – HTTP/3/QUIC with `CONNECT-UDP` or `CONNECT-IP` on
   port 443; indistinguishable from normal HTTPS to most middleboxes.
3. **OpenVPN** – TLS 1.3 over TCP/443 using only AEAD ciphers and
   `tls-crypt-v2` for control channel protection.
4. **IKEv2/IPsec** – optional profile using the algorithms described
   in the security baseline; supports UDP encapsulation and TCP
   encapsulation for NAT traversal.

The client should try these transports in order until one succeeds.
