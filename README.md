English | [简体中文](/README-zh-CN.md)

# Darabonba Code Generator for Go

[![NPM version][npm-image]][npm-url]
[![build status][travis-image]][travis-url]
[![codecov][cov-image]][cov-url]
[![dependency status][deps-image]][deps-url]
[![npm download][download-image]][download-url]

[npm-image]: https://img.shields.io/npm/v/@darabonba/go-generator.svg?style=flat-square
[npm-url]: https://npmjs.org/package/@darabonba/go-generator
[travis-image]: https://img.shields.io/travis/aliyun/darabonba-go-generator.svg?style=flat-square
[travis-url]: https://travis-ci.org/aliyun/darabonba-go-generator
[cov-image]: https://codecov.io/gh/aliyun/darabonba-go-generator/branch/master/graph/badge.svg
[cov-url]: https://codecov.io/gh/aliyun/darabonba-go-generator
[deps-image]: https://img.shields.io/librariesio/release/npm/@darabonba/go-generator
[deps-url]: https://libraries.io/npm/@darabonba%2Fgo-generator
[download-image]: https://img.shields.io/npm/dm/@darabonba/go-generator.svg?style=flat-square
[download-url]: https://npmjs.org/package/@darabonba/go-generator

## Installation

Darabonba Code Generator was designed to work in Node.js. The preferred way to install the Generator is to use the [NPM](https://www.npmjs.com/) package manager. Simply type the following into a terminal window:

```shell
npm install @darabonba/go-generator
```

## Usage

```js
'use strict';

const path = require('path');
const fs = require('fs');

const parser = require('@darabonba/parser');
const Generator = require('@darabonba/go-generator');

const sourceDir = "<Darabonda package directory>";
const outputDir = "<Generate output directory>";

// generate AST data by Parser
let packageMetaFilePath = path.join(sourceDir, 'Darafile');
let packageMeta = JSON.parse(fs.readFileSync(packageMetaFilePath, 'utf8'));
let mainFile = path.join(sourceDir, packageMeta.main);
let ast = parser.parse(fs.readFileSync(mainFile, 'utf8'), mainFile);

// initialize generator
let generatorConfig = {
  ...packageMeta,
  pkgDir: sourceDir,
  outputDir
};

let generator = new Generator(generatorConfig);

// generate go code by generator
generator.visit(ast);

// The execution result will be output in the 'outputDir'
```

## Issues

[Opening an Issue](https://github.com/aliyun/darabonba-go-generator/issues/new/choose), Issues not conforming to the guidelines may be closed immediately.

## Changelog

Detailed changes for each release are documented in the [release notes](/CHANGELOG.md).

## License

[Apache-2.0](/LICENSE)
Copyright (c) 2009-present, Alibaba Cloud All rights reserved.
