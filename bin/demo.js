#!/usr/bin/env node

'use strict';

const Generator = require('../lib/generator');

class NodeGenerator extends Generator {

  modelBefore(level) {
    this.emit(`
    before models
`);
  }

  modelAfter(level) {
    this.emit(`
    after models
`);
  }

  moduleBefore(level) {
    this.emit(`
    before module
`);
  }

  moduleAfter(level) {
    this.emit(`
    after module
`);
  }

//   # BEGIN. TO BE Implement.

// class Client

  apiBefore(__module, level) {
    this.emit(`
`);
  }

  apiAfter(__module, level = 0) {
    this.emit(`
`);
  }
}

const ast = require('./demo.ast.json');
const generator = new NodeGenerator();
generator.visit(ast);
console.log(generator.output);
