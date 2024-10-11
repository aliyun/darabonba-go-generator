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
  const clientPath = path.join(outputDir, pkgInfo.exec ? 'main/main.go' : 'client/client.go');
  const expected = fs.readFileSync(expectedPath, 'utf8');
  assert.deepStrictEqual(fs.readFileSync(clientPath, 'utf8'), expected);
}

describe('new Generator', function () {
  it('must pass in outputDir', function () {
    assert.throws(function () {
      new Generator({});
    }, function (err) {
      assert.deepStrictEqual(err.message, '`option.outputDir` should not empty');
      return true;
    });
  });

  it('one model should ok', function () {
    const outputDir = path.join(__dirname, 'output/model');
    const mainFilePath = path.join(__dirname, 'fixtures/model/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/model/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/model/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/model'),
      ...pkg
    });
    pkg.go = {
      mapAndSliceWithoutOmitempty: true,
    };
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/model/client_no_omit.go'), {
      pkgDir: path.join(__dirname, 'fixtures/model'),
      ...pkg
    });
  });

  it('one api should ok', function () {
    const outputDir = path.join(__dirname, 'output/api');
    const mainFilePath = path.join(__dirname, 'fixtures/api/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/api/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/api/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/api'),
      ...pkg
    });
  });

  it('one function should ok', function () {
    const outputDir = path.join(__dirname, 'output/function');
    const mainFilePath = path.join(__dirname, 'fixtures/function/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/function/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/function/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/function'),
      ...pkg
    });
  });

  it('statements should ok', function () {
    const outputDir = path.join(__dirname, 'output/statements');
    const mainFilePath = path.join(__dirname, 'fixtures/statements/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/statements/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/statements/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/statements'),
      ...pkg
    });
  });

  it('import should ok', function () {
    const outputDir = path.join(__dirname, 'output/import');
    const mainFilePath = path.join(__dirname, 'fixtures/import/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/import/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/import/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/import'),
      ...pkg
    });
  });

  it('extends should ok', function () {
    const outputDir = path.join(__dirname, 'output/extends');
    const mainFilePath = path.join(__dirname, 'fixtures/extends/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/extends/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/extends/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/extends'),
      ...pkg
    });
  });

  it('extendsWithoutInit should ok', function () {
    const outputDir = path.join(__dirname, 'output/extendsWithoutInit');
    const mainFilePath = path.join(__dirname, 'fixtures/extendsWithoutInit/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/extendsWithoutInit/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/extendsWithoutInit/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/extendsWithoutInit'),
      ...pkg
    });
  });

  it('complex should ok', function () {
    const outputDir = path.join(__dirname, 'output/complex');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/complex/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    const mainFilePath = path.join(__dirname, 'fixtures/complex/main.dara');
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/complex/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/complex'),
      package: ['github.com/alibabacloud-go/tea'],
      ...pkg
    });
    const modPath = path.join(__dirname, 'fixtures/complex/go.mod');
    const expected = fs.readFileSync(path.join(outputDir, 'go.mod'), 'utf8');
    assert.deepStrictEqual(fs.readFileSync(modPath, 'utf8'), expected);
  });

  it('add annotation should ok', function () {
    const outputDir = path.join(__dirname, 'output/annotation');
    const mainFilePath = path.join(__dirname, 'fixtures/annotation/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/annotation/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/annotation/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/annotation'),
      ...pkg,
      editable: true
    });
  });

  it('add comment should ok', function () {
    const outputDir = path.join(__dirname, 'output/comment');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/comment/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    const mainFilePath = path.join(__dirname, 'fixtures/comment/main.dara');
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/comment/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/comment'),
      ...pkg
    });
  });

  it('tea should ok', function () {
    const outputDir = path.join(__dirname, 'output/tea');
    const mainFilePath = path.join(__dirname, 'fixtures/tea/main.tea');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/tea/Teafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/tea/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/tea'),
      ...pkg
    });
  });

  it('main should ok', function () {
    const outputDir = path.join(__dirname, 'output/main');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/main/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    const mainFilePath = path.join(__dirname, 'fixtures/main/main.dara');
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/main/main.go'), {
      pkgDir: path.join(__dirname, 'fixtures/main'),
      exec: true,
      ...pkg
    });
  });

  it('interface access should ok', function () {
    const outputDir = path.join(__dirname, 'output/interface');
    const mainFilePath = path.join(__dirname, 'fixtures/interface/main.tea');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/interface/Teafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/interface/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/interface'),
      ...pkg
    });
  });

  it('interface should ok', function () {
    const outputDir = path.join(__dirname, 'output/interface');
    const mainFilePath = path.join(__dirname, 'fixtures/interface/spi.tea');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/interface/Teafile1'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/interface/spi.go'), {
      pkgDir: path.join(__dirname, 'fixtures/interface'),
      ...pkg
    });
  });

  it('typedef should ok', function () {
    const outputDir = path.join(__dirname, 'output/typedef');
    const mainFilePath = path.join(__dirname, 'fixtures/typedef/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/typedef/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/typedef/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/typedef'),
      package: ['net/http'],
      ...pkg
    });
  });

  it('builtin should ok', function () {
    const outputDir = path.join(__dirname, 'output/builtin');
    const mainFilePath = path.join(__dirname, 'fixtures/builtin/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/builtin/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/builtin/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/builtin'),
      ...pkg
    });
  });

  it('try should ok', function () {
    const outputDir = path.join(__dirname, 'output/try');
    const mainFilePath = path.join(__dirname, 'fixtures/try/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/try/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/try/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/try'),
      ...pkg
    });
  });

  it('multi should ok', function () {
    const outputDir = path.join(__dirname, 'output/multi');
    const mainFilePath = path.join(__dirname, 'fixtures/multi/tea/sdk.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/multi/tea/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    const generator = new Generator({
      outputDir,
      pkgDir: path.join(__dirname, 'fixtures/multi/tea'),
      ...pkg,
    });
  
    const dsl = fs.readFileSync(mainFilePath, 'utf8');
    const ast = DSL.parse(dsl, mainFilePath);
    generator.visit(ast);
    const mainPath = path.join(outputDir, 'client/client.go');
    const apiPath = path.join(outputDir, 'api/client.go');
    const modelPath = path.join(outputDir, 'model/user/client.go');
    const utilPath = path.join(outputDir, 'lib/util/client.go');
    const expectedMainPath = path.join(__dirname, 'fixtures/multi/sdk/client.go');
    const expectedModelPath = path.join(__dirname, 'fixtures/multi/sdk/user.go');
    const expectedUtilPath = path.join(__dirname, 'fixtures/multi/sdk/util.go');
    const expectedApiPath = path.join(__dirname, 'fixtures/multi/sdk/api.go');
    const expectedMain = fs.readFileSync(expectedMainPath, 'utf8');
    assert.deepStrictEqual(fs.readFileSync(mainPath, 'utf8'), expectedMain);
    const expectedModel = fs.readFileSync(expectedModelPath, 'utf8');
    assert.deepStrictEqual(fs.readFileSync(modelPath, 'utf8'), expectedModel);
    const expectedUtil = fs.readFileSync(expectedUtilPath, 'utf8');
    assert.deepStrictEqual(fs.readFileSync(utilPath, 'utf8'), expectedUtil);
    const expectedApi = fs.readFileSync(expectedApiPath, 'utf8');
    assert.deepStrictEqual(fs.readFileSync(apiPath, 'utf8'), expectedApi);
  });


  it('yield should ok', function () {
    const outputDir = path.join(__dirname, 'output/yield');
    const mainFilePath = path.join(__dirname, 'fixtures/yield/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/yield/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/yield/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/yield'),
      ...pkg
    });
  });

  it('exception should ok', function () {
    const outputDir = path.join(__dirname, 'output/exception');
    const mainFilePath = path.join(__dirname, 'fixtures/exception/main.dara');
    const pkgContent = fs.readFileSync(path.join(__dirname, 'fixtures/exception/Darafile'), 'utf8');
    const pkg = JSON.parse(pkgContent);
    check(mainFilePath, outputDir, path.join(__dirname, 'fixtures/exception/client.go'), {
      pkgDir: path.join(__dirname, 'fixtures/exception'),
      ...pkg
    });
  
  });
});
