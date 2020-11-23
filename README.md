<p align="center">
    <img alt="Craft Logo" src="/static/logo.png?v=1.0.0" width="200" />
    <h3 align="center">Craft</h3>
    <p align="center">Code Snippets Manager</p>
    <p align="center">
        <a href="https://github.com/spacewalkio/Craft/actions/workflows/build.yml">
            <img src="https://github.com/spacewalkio/Craft/actions/workflows/build.yml/badge.svg">
        </a>
        <a href="https://github.com/spacewalkio/Craft/releases">
            <img src="https://img.shields.io/badge/Version-v1.0.0-red.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/spacewalkio/Craft">
            <img src="https://goreportcard.com/badge/github.com/spacewalkio/Craft?v=1.0.0">
        </a>
        <a href="https://godoc.org/github.com/spacewalkio/craft">
            <img src="https://godoc.org/github.com/spacewalkio/craft?status.svg">
        </a>
        <a href="https://github.com/spacewalkio/Craft/blob/main/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
        </a>
    </p>
</p>
<br/>


## Usage

Download [the latest `craft` binary](https://github.com/spacewalkio/Craft/releases). Make it executable from everywhere.

```zsh
$ export CRAFT_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/spacewalkio/Craft/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

# For Linux
$ curl -sL https://github.com/spacewalkio/Craft/releases/download/v{$CRAFT_LATEST_VERSION}/craft_{$CRAFT_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz

# For Mac
$ curl -sL https://github.com/spacewalkio/Craft/releases/download/v{$CRAFT_LATEST_VERSION}/craft_{$CRAFT_LATEST_VERSION}_Darwin_x86_64.tar.gz | tar xz
```

Or install with homebrew

```zsh
$ brew tap spacewalkio/tools
$ brew install spacewalkio/tools/craft
```

Configure the craft using the following command

```zsh
$ craft config
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Craft is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/spacewalkio/craft/releases) for changelogs for each release version of Craft. It contains summaries of the most noteworthy changes made in each release. Also see the [Milestones section](https://github.com/spacewalkio/craft/milestones) for the future roadmap.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/spacewalkio/craft/issues


## Security Issues

If you discover a security vulnerability within Craft, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2022, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Craft** is authored and maintained by [@clivern](http://github.com/clivern).
