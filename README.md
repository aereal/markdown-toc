# markdown-toc

<!-- toc -->
<!-- /toc -->

----

## Usage

```
cat README.markdown
=> # README
=> 
=> <!-- toc -->
=> <!-- /toc -->
=> 
=> ## Usage
=> ...
=> ## Install
=> ...
=> ## License

markdown-toc README.markdown
=> # README
=> 
=> <!-- toc -->
=> * [Usage](#usage)
=> * [Install](#install)
=> * [License](#license)
=> <!-- /toc -->
=> 
=> ## Usage
=> ...
=> ## Install
=> ...
=> ## License
```

## Install

```
go install github.com/aereal/markdown-toc/cmd/markdown-toc
```

## See also

Inspired by [hail2u/node-gfmtoc][]

[hail2u/node-gfmtoc]: https://github.com/hail2u/node-gfmtoc
