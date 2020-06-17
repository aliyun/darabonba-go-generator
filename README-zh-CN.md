[English](/README.md) | 简体中文

# Darabonba Go 生成器

## 安装

> Darabonba 生成器只能在 Node.js 环境下运行。
> 建议使用 [NPM](https://www.npmjs.com/) 包管理工具安装
> 在终端输入以下命令进行安装:

```shell
npm install @darabonba/go-generator
```

## 使用示例

> 生成 Go 代码

```javascript
'use strict';

const path = require('path');
const fs = require('fs');

const DarabonbaParser = require('@darabonba/parser');
const DarabonbaGoGenerator = require('@darabonba/go-generator');
const sourceDir = "<Darabonda package directory>";
const outputDir = "<Generate output directory>";

// generate AST data by DarabonbaParser
let darabonbaPackageMetaFilePath = path.join(sourceDir, 'Teafile');
let darabonbaMainFile = path.join(sourceDir, darabonbaPackageMeta.main);

let darabonbaPackageMeta = JSON.parse(fs.readFileSync(darabonbaPackageMetaFilePath, 'utf8'));
let darabonbaAST = DarabonbaParser.parse(fs.readFileSync(darabonbaMainFile, 'utf8'), darabonbaMainFile);

// initialize generator
let generatorConfig = {
      ...darabonbaPackageMeta,
      pkgDir: sourceDir,
      outputDir
    };

let generator = new DarabonbaGoGenerator(generatorConfig);

// generate go code by generator
generator.visit(darabonbaAST);

// The execution result will be output in the 'outputDir'
```

## 问题反馈

[提出问题](https://github.com/aliyun/darabonba-go-generator/issues/new/choose), 不符合指南的问题可能会立即关闭。

## 发布日志

发布详情会更新在 [release notes](/CHANGELOG.md) 文件中

## 许可证

[Apache-2.0](/LICENSE)
Copyright (c) 2009-present, Alibaba Cloud All rights reserved.