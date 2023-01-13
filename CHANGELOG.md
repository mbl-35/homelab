# Changelog

## TODO

[X] Change bootstrap/seek-values to match with my repo to get changes
[X] Remove dedicated domain names (khuedoan.com, 127-0-0-1.nip.io) to a local usable domain kube.localhost and push it
[X] Know nixos user name: `NIX_USER` env variable
[X] Makefile: add environment selector (`make env target:dev`)
[ ] Make this directory as safe for git
[X] Update dev dynamic ip address to inventory (e.g. wsl)
[ ] Ability to set a external git repo (vs internal gitea) - variabilize
[ ] Manage dotfiles for staging levels (maybe multi cluster management)
[ ] Manage metal cluster by configuration (k3s, kubesphere, ..)
[ ] Make clean : add remove k3d cluster


## Initial diffs khuedoan/homelab branch prod vs dev
- kube.localhost => 127-0-0-1.nip.io
- Makefile default changes 
- bootstrap/root/templates/stack.yaml : syncOptions:[] + SkipDryRunOnMissingResource=true
- metal/Makefile default boot cluster != k3d (+k3d definition)
- storageClassName: longhorn => local-path (vault et trow)
- suppression system/longhorn-system
- system-upgrade => commenter resources: k3s
- system-upgrade/[agent|server] => version: v1.24.9+k3s1 vs channel: https://update.k3s.io/v1-release/channels/v1.23
