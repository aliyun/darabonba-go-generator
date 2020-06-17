'use strict';

const path = require('path');
const fs = require('fs');
const assert = require('assert');

const DSL = require('@darabonba/parser');

let Generator = require('../lib/generator');

function check(mainFilePath, outputDir, expectedPath, pkgInfo = {}) {
  const generator = new Generator({
    outputDir,
    baseClient: 'github.com/aliyun/test/go',
    ...pkgInfo
  });

  const dsl = fs.readFileSync(mainFilePath, 'utf8');
  const ast = DSL.parse(dsl, mainFilePath);
  generator.visit(ast);
  const clientPath = path.join(outputDir, 'client/client.go');
  const expected = fs.readFileSync(expectedPath, 'utf8');
  assert.deepStrictEqual(fs.readFileSync(clientPath, 'utf8'), expected);
}

describe('new Generator', function() {
  it('must pass in outputDir', function () {
    assert.throws(function () {
      new Generator({});
    }, function(err) {
      assert.deepStrictEqual(err.message, '`option.outputDir` should not empty');
      return true;
    });
  });

  it('one model should ok', function () {
    const outputDir = path.join(__dirname, 'output/model');
    const mainFilePath = path.join(__dirname, 'fixtures/model/main.tea');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/model/Teafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/model/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/model'),
      ...pkg
    });
  });

  it('one api should ok', function () {
    const outputDir = path.join(__dirname, 'output/api');
    const mainFilePath = path.join(__dirname, 'fixtures/api/main.tea');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/api/Teafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/api/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/api'),
      ...pkg
    });
  });

  it('one function should ok', function () {
    const outputDir = path.join(__dirname, 'output/function');
    const mainFilePath = path.join(__dirname, 'fixtures/function/main.tea');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/function/Teafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/function/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/function'),
      ...pkg
    });
  });

  it('statements should ok', function () {
    const outputDir = path.join(__dirname, 'output/statements');
    const mainFilePath = path.join(__dirname, 'fixtures/statements/main.tea');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/statements/Teafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/statements/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/statements'),
      ...pkg
    });
  });

  it('import should ok', function () {
    const outputDir = path.join(__dirname, 'output/import');
    const mainFilePath = path.join(__dirname, 'fixtures/import/main.tea');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/import/Teafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/import/client.go'),{
      pkgDir: path.join(__dirname, 'fixtures/import'),
      ...pkg
    });
  });

  it('complex should ok', function () {
    const outputDir = path.join(__dirname, 'output/complex');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/complex/Teafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    const mainFilePath = path.join(__dirname, 'fixtures/complex/main.tea');
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/complex/client.go') , {
      pkgDir: path.join(__dirname, 'fixtures/complex'),
      ...pkg
    });
    const modPath = path.join(__dirname, 'fixtures/complex/go.mod');
    const expected = fs.readFileSync(path.join(outputDir, 'go.mod'), 'utf8');
    assert.deepStrictEqual(fs.readFileSync(modPath, 'utf8'), expected);
  });

  it('add annotation should ok', function () {
    const outputDir = path.join(__dirname, 'output/annotation');
    const mainFilePath = path.join(__dirname, 'fixtures/annotation/main.tea');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/annotation/Teafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/annotation/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/annotation'),
      ...pkg
    });
  });

  it('add comment should ok', function () {
    const outputDir = path.join(__dirname, 'output/comment');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/comment/Teafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    const mainFilePath = path.join(__dirname, 'fixtures/comment/main.tea');
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/comment/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/comment'),
      ...pkg
    });
  });
});
