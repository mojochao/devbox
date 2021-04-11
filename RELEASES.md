# Devbox Releases

## 0.13.1

- Fixes for bugs in use of kubeconfig path and namespace in kubectl commands 

## 0.13.0

- Changed the `done` command used in processing devbox manifest commands to
  `break`, as the semantics are the same as breaking out of a for loop over
  the manifest items

## 0.12.1

- Fixes bad help descriptions of `add` command flags caused by an overzealous
  reach inside string literals

## 0.12.0

- Improved the UX of the `add` command
- Added `--reset` flag to the `context` command

## 0.11.0

- Fixed a bug when executing setup commands in Docker devboxes.

## 0.10.0

- Enhanced box state to enable declarative provisioning per devbox

## 0.9.0

- Added setup command to provision devbox container

## 0.8.0

- Enhanced box state with the devbox username and reflected that in the add,
  ctx, and list commands.

## 0.7.0

- Enhanced shell command

## 0.6.0

- Bumped ulimits on file descriptors for load test usage.

## 0.5.1

- Bumped version for rebuild.

## 0.5.0

- Removed copy command. It's less capable than 'docker cp' or 'kubectl cp' so
  doesn't add any real value.
  
## 0.4.0

- Applied minor refactorings for clarity

## 0.3.0

- Fixed panic when extracting cluster name from kubeconfig

## 0.2.0

- Improved docs
- Improved flexibility of start, stop, and remove commands
- Renamed go module to just devbox instead of devbox-cli

## 0.1.0

- Initial implementation
