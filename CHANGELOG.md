# Changelog


## diffs khuedoan/homelab branch prod vs dev :

- kube.localhost => 127-0-0-1.nip.io
- bootstrap/root/templates/stack.yaml : syncOptions:[] + SkipDryRunOnMissingResource=true
- metal/Makefile default boot cluster != k3d (+k3d definition)
- storageClassName: longhorn => local-path (vault et trow)
- suppression system/longhorn-system
- system-upgrade => commenter resources: k3s
- system-upgrade/[agent|server] => version: v1.24.9+k3s1 vs channel: https://update.k3s.io/v1-release/channels/v1.23

+ variabiliser le git externe (pas de git interne)
+ gestion des environnements (et des différents clusters ?)
+ options k3s vs kubesphere

# TODO

[X] Change bootstrap/seek-values to match with my repo to get changes
[X] Remove dedicated domain names (khuedoan.com, 127-0-0-1.nip.io) to a local usable domain kube.localhost and push it
[ ] Makefile: add environment selector (make env target:dev)