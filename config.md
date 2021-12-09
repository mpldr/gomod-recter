---
title: "Configuration"
---

# Configuration

The config for recter is split into multiple chapters. The configuration format
is TOML and it is automatically reloaded and applied on the fly (with a few
exceptions).

# Chapters

## Root

The root chapter is the highest level. It does not have to be specified.

### Domain

The domain you wish to use for the import path. This must be the same as the
Reverse Proxy is available at.

### VersionRefreshInterval

Default: `5m`

How often to query for new versions.

## Directories

### DataDir

*not yet relevant*

This is where recter might save runtime and cache data in the future. Currently
this setting has no effect.

### TemplateDir

This indicates where recter is looking for templates used for generating the
html pages.

### AssetDir

This indicates where assets for the HTML page can be found. You can access the
directory structure by requesting it with the prefix `/assets/`.

#### Example

Your asset is
`/home/user/recter-test/themes/my-theme/assetfiles/css/my-style.css`. You can
set `AssetDir` to `/home/user/recter-test/themes/my-theme/assetfiles` and
request your file by requesting `/assets/css/my-style.css`.

## Network

### Type

Possible Values: `tcp`, `unix`

Type allows setting the network type to either TCP or UNIX-Socket.

### SocketPath

The path to where the UNIX-Socket will be generated. It will take the
permissions of the parent directory and there is no way to specify permissions
another way.

### ListenAddr

Format: `Address:Port`

Default: `127.0.0.1:25000`

Indicates interface and Port to listen on.

## Proxy

### Address

Default: `https://proxy.golang.org`

The GOPROXY to use for querying new versions.

### IgnoreCert

Do not validate SSL-Certificates. This is currently implemented until the
container is built with the CA Certs.

## Projects

### *

Every subkey is the path of a project. The Project Settings are defined in the
subkeys.

#### Name¹

The human readable name of the project, shown on the HTML page.

#### Redirect

Whether to directly redirect to the Projects repository. If `Redirect` is set
to `true`, Keys marked with ¹ are not used.

#### Description¹

A description of what the project is about.

#### VCS

The kind of VCS that is used. This is usually `git` but `svn`, `hg`, `fossil`,
and `bzr` are also officially supported.

#### Repo

The actual repo where your code lives.

#### License¹

This is the license of the project.

#### DefaultBranch

This setting is used to allow Documentation renderers to link to the
sourcecode. If this is not set, the instructions on where to find the code is
omitted.

#### GoSourceFmt

Some popular sourcehosting sites are automatically detected through the Repo-Host. If your favourite site is not automatically detected, feel free to send a patch.

Common formats for Gitea, Gitlab, and more can be found
[here](https://git.sr.ht/~poldi1405/gomod-recter/tree/master/item/internal/data/get-data.go#L107-127,135-155)
and can be added as "[dirstring] [filestring]".

#### Note¹

Note can be used to display a notice (for example regarding an announcement) on
the project page.

##### Show

Whether or not to display the notice.

##### Text

The text of the notice.

##### Style

The style of the notice.

# Footnotes

¹) This option does not apply if Redirect is set to true for the project.
