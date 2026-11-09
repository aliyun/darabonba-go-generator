'use strict';

const assert = require('assert');
const fs = require('fs');
const path = require('path');
const DSL = require('@darabonba/parser');
const Generator = require('../lib/generator');

describe('exception names', function () {
  const outputDir = path.join(__dirname, 'output/exceptionNames');

  before(function () {
    const source = `
exception AlibabaCloud extends $ResponseError {
  detail?: string,
  requestId?: string,
}
exception Client extends AlibabaCloud {}
exception Server extends AlibabaCloud {}
exception Throttling extends AlibabaCloud {}
model Named {
  name?: string,
}
`;
    const mainFilePath = path.join(__dirname, 'fixtures/exception/main.dara');
    const generator = new Generator({
      outputDir,
      pkgDir: path.join(__dirname, 'fixtures/exception'),
      baseClient: 'github.com/aliyun/test/go',
    });
    generator.visit(DSL.parse(source, mainFilePath));
  });

  [
    ['alibaba_cloud_error.go', 'AlibabaCloudError'],
    ['client_error.go', 'ClientError'],
    ['server_error.go', 'ServerError'],
    ['throttling_error.go', 'ThrottlingError'],
  ].forEach(([filename, typeName]) => {
    it(`${typeName} should default to its concrete name and preserve explicit names`, function () {
      const generated = fs.readFileSync(path.join(outputDir, 'client', filename), 'utf8');
      assert.ok(generated.includes(
        `func (s *${typeName}) GetName() *string  {\n` +
        '  if s.Name == nil {\n' +
        `    return dara.String("${typeName}")\n` +
        '  }\n' +
        '  return s.Name\n' +
        '}\n'
      ), `${typeName}.GetName must expose a name for exception retry conditions`);
    });
  });

  it('ordinary model name getters should not acquire exception defaults', function () {
    const generated = fs.readFileSync(path.join(outputDir, 'client/named_model.go'), 'utf8');
    assert.ok(generated.includes(
      'func (s *Named) GetName() *string  {\n' +
      '  return s.Name\n' +
      '}\n'
    ));
  });
});
