# Docktags
Docktags is a tool to list all available tags for a given docker hub repo or
repos.

## Usage
Invoke `docktags` with one or more repo names you found via the docker search.

### Example
```
$ docktags dock0/arch codekoala/arch centos

dock0/arch (1/1):
    latest

codekoala/arch (2/2):
    latest lfs

centos (21/21):
    5 5.11 6 6.6 6.7 6.8 7 7.0.1406 7.1.1503 7.2.1511 
    centos5 centos5.11 centos6 centos6.6 centos6.7 centos6.8
    centos7 centos7.0.1406 centos7.1.1503 centos7.2.1511 latest
```

## Purpose

Sometimes you need to know the tags a docker repository provides.
The docker command line app provides no way to list the available tags
and brwsing to the hub manually is tedious.

Searching for a pre-made solution brought me to examples like
[this](http://stackoverflow.com/questions/28320134/how-to-list-all-tags-for-a-docker-image-on-a-remote-registry)
or [this](http://stackoverflow.com/questions/28320134/how-to-list-all-tags-for-a-docker-image-on-a-remote-registry), which either use the the v1 API (which doesn't list all tags), can't handle the v2 pagination or tell me to just pull all tags from a repo, terabyte downloads be damned.

Obviously none of these "solutions" was acceptable. So, since I am trying to learn Go anyway, I decided to create my own tool. So far it can parse `hub.docker.com` and spew out all tags for a repository, that's it.

## Installation

No binaries yet.

`go get github.com/jinks/docktags`
