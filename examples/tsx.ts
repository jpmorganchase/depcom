// This file has a ts extension, but in reality it's tsx
import React from "react";
import foo from "tsx-in-ts-static-import";

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
