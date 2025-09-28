# Gateway setup

FreePN gateways terminate VPN connections and forward traffic to the
Internet.  They run on Linux and rely on WireGuard for the data plane
and nftables for NAT and filtering.  This directory contains an
Ansible playbook and templates to provision a gateway from scratch.

## Requirements

* A Linux host (Ubuntu 22.04 or similar) with SSH access.
* Ansible 2.10+ installed on your control machine.

## Usage

1. Copy `ansible/inventory.ini.sample` to `inventory.ini` and edit it
   to point at your gateway host.  For example:

   ```ini
   [freepn_gateways]
   my-gateway ansible_host=203.0.113.42 ansible_user=root
   ```

2. Run the playbook:

   ```bash
   cd ansible
   ansible-playbook -i inventory.ini playbook.yml
   ```

   The playbook installs WireGuard, sets up the `wg0` interface using
   `wg0.conf.j2`, enables IP forwarding and applies nftables rules
   defined in `freepn.nft.j2`.

3. Adjust the templates for your environment.  For example, change
   the `Address` and `ListenPort` in `wg0.conf.j2` or add additional
   nftables rules.

This is only a starting point; see `docs/architecture.md` and
`docs/security-baseline.md` for guidance on egress filtering and
kill‑switch policies.
