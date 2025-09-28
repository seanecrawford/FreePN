# FreePN Security & Privacy Baseline

This document describes the default security posture of FreePN.  It
summarises the recommended algorithms, transport protocols and
operational practices.  The content here is intended to be kept in
sync with upstream Internet standards and hardening guidelines.  If
you modify the baseline, please update this file accordingly.

## Transport and cryptography

* **TLS 1.3 only** for all TLS‑based channels.  Supported suites
  should be restricted to `TLS_AES_128_GCM_SHA256`,
  `TLS_AES_256_GCM_SHA384` and `TLS_CHACHA20_POLY1305_SHA256`.  Avoid
  legacy suites and disable TLS 1.2.
* **WireGuard** uses the Noise_IK pattern with X25519 for key
  exchange, ChaCha20‑Poly1305 for encryption, BLAKE2s for hashing, and
  HKDF for key derivation.  Clients and servers should rekey
  frequently.
* **MASQUE** (HTTP/3/QUIC) is used as the primary fallback.  Use
  `CONNECT-UDP` or `CONNECT-IP` over port 443.
* **OpenVPN** fallback must enforce TLS 1.3, AEAD‑only ciphers on the
  data channel (`AES-256-GCM`, `AES-128-GCM`, `CHACHA20-POLY1305`) and
  disable compression.
* **IKEv2/IPsec** (optional) proposals should include AES‑GCM and
  ChaCha20‑Poly1305 with PFS groups X25519 and P-256.  NAT traversal
  and TCP encapsulation should be enabled for networks that block
  UDP.

## DNS and resolver policy

All client DNS queries go through the tunnel.  The resolvers speak
**DoT** or **DoH** to upstream servers, and QNAME minimisation is
enabled to reduce leakage of user queries.  The client should never
fall back to plaintext DNS.

## Kill switch and leak prevention

Each client platform must implement a kill switch that blocks all
traffic outside the VPN interface.  On Windows use the Windows
Filtering Platform (WFP); on macOS and iOS use
NetworkExtension's packet tunnel; on Android use VpnService with
"always-on" and "lockdown" modes; on Linux use nftables rules.  The
kill switch must be active before establishing the tunnel and remain
enabled until the user disconnects.

## Logging and telemetry

FreePN does not retain activity logs or correlate user identities with
source IPs.  Gateways may keep aggregate counters (for example
bandwidth usage per point-of-presence) for a short period (less than
24 hours) to detect abuse and plan capacity.  If operational logs
exist, they must only store timestamps and source ports.  All log
retention policies should be documented publicly.

## Abuse controls

Running a free exit service requires protecting the wider Internet
from abuse.  Gateways enforce ingress filtering (BCP 38/84) to
prevent IP spoofing and block known abuse ports (SMTP, NetBIOS,
etc.).  The control plane should implement rate limits or token
buckets to throttle excessive usage.  Community exit nodes, if
offered, must be sandboxed and opt-in.

## Supply chain and software integrity

FreePN publishes a Software Bill of Materials (SBOM) with each
release.  Builds are reproducible; artifacts are signed using
Sigstore's cosign; update metadata is secured via The Update
Framework (TUF) and in-toto attestations.  We aim for SLSA Level 2
compliance or higher.

## Post‑quantum considerations

NIST has released the first batch of post-quantum cryptography (PQC)
standards.  FreePN will monitor progress and adopt hybrid key
exchanges (for example, X25519 combined with ML-KEM) in TLS and IPsec
once mainstream libraries support them.  For now, keep rekey intervals
short and use tls-crypt-v2 on OpenVPN control channels to protect
against passive recording.
