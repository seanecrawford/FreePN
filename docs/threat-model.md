# Threat model

Understanding the capabilities of potential adversaries is critical
when designing a VPN.  This document outlines the threat model for
FreePN and the assumptions that inform our security baseline.

## In‑scope adversaries

* **Local network observers** – attackers on the same Wi‑Fi or
  physical network segment may sniff unencrypted traffic or attempt to
  inject packets.  FreePN encrypts all user traffic and prevents
  leakage via kill switches.
* **Internet service providers (ISPs)** – ISPs can monitor and log
  customer traffic.  Encrypted tunnels hide contents and destinations,
  and MASQUE/OpenVPN fallbacks blend in with common HTTPS traffic.
* **Censors and firewalls** – middleboxes may block or throttle
  VPN‑specific traffic.  FreePN includes multiple fallbacks (WireGuard,
  MASQUE, OpenVPN, IKEv2) to traverse restrictive networks.  MASQUE
  uses port 443 and the HTTP/3 protocol to resemble normal web traffic.
* **Malicious exit/gateway operators** – we assume that gateway
  operators may be curious or malicious.  Gateways see decrypted
  traffic, so clients must not rely on the VPN for end‑to‑end
  confidentiality; users should use TLS/HTTPS when possible.  We
  mitigate this risk by not logging user activity and by allowing
  users to choose different exits or run their own.

## Out‑of‑scope adversaries

* **Global passive adversaries** – entities capable of monitoring
  traffic at multiple Internet exchange points simultaneously (e.g.
  state‑level actors) may perform traffic correlation attacks.  FreePN
  does not claim to defeat such attackers.  Multi‑hop or mix networks
  are outside the scope of this project.
* **Compromised client devices** – if a user's device is already
  compromised (e.g. malware, rootkit), a VPN cannot protect
  confidentiality.  Users should maintain general device hygiene and
  security.
* **Application‑layer vulnerabilities** – FreePN transports packets
  securely, but it does not patch vulnerabilities in higher‑layer
  protocols (e.g. HTTP).  Users should rely on application‑level
  security like TLS and follow secure coding practices.

## Trust assumptions

FreePN makes the following trust assumptions:

1. **Control plane operators** are honest and will not misuse
   cryptographic keys or allocate identical IP addresses to multiple
   clients.  The control plane should run in a secure environment with
   hardware security modules for key storage.
2. **Gateway operators** will run vetted gateway software and adhere
   to the documented egress policies.  They will not log user
   activity.  Users may choose to operate their own gateways or trust
   community‑operated ones.
3. **Clients** generate their own keypairs and protect private keys on
   disk or in secure enclaves.  Clients are responsible for verifying
   the authenticity of the control plane and gateways (e.g. via
   pinned certificates or pre‑shared keys).
