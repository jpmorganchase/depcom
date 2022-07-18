// This file has a js extension, but in reality it's jsx
import React from "react";
import foo from "jsx-in-js-static-import";

export default function JsxComponent() {
  return (
    <div>
      {/* Static import */}
      {foo}
      {/* Dynamic import */}
      {import("jsx-in-js-static-import").then(() => console.log("imported!"))}
    </div>
  );
}
