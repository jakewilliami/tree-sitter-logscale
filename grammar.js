/**
 * @file Tree-sitter parser for LogScale/CQL
 * @author Jake W. Ireland <jakewilliami@icloud.com>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "logscale",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
