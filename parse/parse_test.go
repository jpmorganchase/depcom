package parse

import (
	"reflect"
	"testing"

	"github.com/ije/esbuild-internal/logger"
)

type testCaseECMA struct {
	name      string
	code      string
	importMap []string
	log       LogMap
	extension string
}

type testCaseCSS struct {
	name      string
	code      string
	importMap []string
	log       LogMap
}

func TestFromECMA(t *testing.T) {
	tests := []testCaseECMA{
		{
			name: "cjs require",
			code: `
			const foo = require("cjs-require/foo");
			let bar = require.resolve("cjs-require-resolve");
			require("./local-import");
			require("/absolute/import");
			console.log(foo, bar);
			console.log(require("cjs-require-in-expression"));
			`,
			extension: ".js",
			importMap: []string{"cjs-require/foo", "cjs-require-resolve", "cjs-require-in-expression"},
			log:       LogMap{},
		},
		{
			name: "esm import",
			code: `
			import foo from "esm-import/foo";
			import { bar } from "esm-import";
			
			console.log(foo, bar);
			console.log(import("esm-dynamic-import").then(() => null))
			`,
			extension: ".js",
			importMap: []string{"esm-import/foo", "esm-import", "esm-dynamic-import"},
			log:       LogMap{},
		},
		{
			name: "esm import in ts",
			code: `
			import foo from "foo/foo";
			import bar from "bar";
			import { baz } from "@scope/baz/qux";
			import corge from "/absolute/path/corge";

			const s: string = "8";
			console.log(foo, bar, baz, corge);
			`,
			extension: ".ts",
			importMap: []string{"foo/foo", "bar", "@scope/baz/qux"},
			log:       LogMap{},
		},
		{
			name: "jsx in js",
			code: `
			// This file has a js extension, but in reality it's jsx
			import React from "react";
			import foo from "jsx-in-js-static-import";

			export default function JsxComponent() {
				return (
					<div>
						{/* Static import */}
						{foo}
						{/* Dynamic import */}
						{import("jsx-in-js-dynamic-import").then(() => console.log("imported!"))}
					</div>
				);
			}
			`,
			extension: ".js",
			importMap: []string{"react", "jsx-in-js-static-import", "jsx-in-js-dynamic-import"},
			log:       LogMap{Err: []string{"./test: The JSX syntax extension is not currently enabled\n"}},
		},
		{
			name: "jsx in jsx",
			code: `
			import react from "react";
			console.log(react);
			console.log(<div>Hello there + {import("@my-scope/foo/bar")}</div>);
			`,
			extension: ".jsx",
			importMap: []string{"react", "@my-scope/foo/bar"},
			log:       LogMap{},
		},
		{
			name: "tsx in jsx, no react import",
			code: `
			// This file has a ts extension, but in reality it's tsx
			import React from "react";
			import foo from "tsx-in-ts-static-import";

			export default function JsxComponent() {
				return (
					<div>
						{/* Static import */}
						{foo}
						{/* Dynamic import */}
						{import("tsx-in-ts-dynamic-import/foo").then(() => console.log("imported foo!"))}
					</div>
				);
			}
			`,
			extension: ".ts",
			// Beware: this will blow the parser up and no import will be parsed as a result
			importMap: nil,
			log:       LogMap{Err: []string{"./test: Expected \")\" but found \"{\"\n"}},
		},
		{
			name: "tsx in tsx",
			code: `
			import React from "react";
			const msg: string = "Hello there";
			console.log(<div>{import("@scope/tsx-in-tsx/bar")}</div>);
			`,
			extension: ".tsx",
			// Beware: this will blow the parser up and no import will be parsed as a result
			importMap: []string{"react", "@scope/tsx-in-tsx/bar"},
			log:       LogMap{},
		},
		{
			name: "import type",
			code: `
			import type myType from "foo";
			import bar from "bar"
			const msg: myType = "Hello there";
			console.log(msg);
			`,
			extension: ".ts",
			importMap: []string{"bar"},
			log:       LogMap{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sourceFile := logger.Source{
				Index:          0,
				KeyPath:        logger.Path{Text: "test"},
				PrettyPath:     "./test",
				Contents:       tt.code,
				IdentifierName: "./test",
			}
			var logMap LogMap
			var importMap []string

			log := NewLogMap(logger.OutputOptions{LogLevel: logger.LevelDebug}, &logMap)

			importMap = FromECMA(&log, &sourceFile, tt.extension)
			if !reflect.DeepEqual(importMap, tt.importMap) {
				t.Errorf("Expected %+v do not match actual %+v", tt.importMap, importMap)
			}
			if !reflect.DeepEqual(logMap, tt.log) {
				t.Errorf("Expected %+v do not match actual %+v", tt.log, logMap)
			}
		})

	}

}

func TestFromCSS(t *testing.T) {
	tests := []testCaseCSS{
		{
			name: "simple css test",
			code: `
			@import url("antd/dist/antd.css");
			@import url("./index.css");
			
			body {
				margin-top: 8px;
			}
				`,
			importMap: []string{"antd/dist/antd.css"},
			log:       LogMap{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sourceFile := logger.Source{
				Index:          0,
				KeyPath:        logger.Path{Text: "test"},
				PrettyPath:     "./test",
				Contents:       tt.code,
				IdentifierName: "./test",
			}
			var logMap LogMap
			var importMap []string

			log := NewLogMap(logger.OutputOptions{LogLevel: logger.LevelDebug}, &logMap)

			importMap = FromCSS(&log, &sourceFile)
			if !reflect.DeepEqual(importMap, tt.importMap) {
				t.Errorf("Expected %+v do not match actual %+v", tt.importMap, importMap)
			}
			if !reflect.DeepEqual(logMap, tt.log) {
				t.Errorf("Expected %+v do not match actual %+v", tt.log, logMap)
			}
		})

	}

}
