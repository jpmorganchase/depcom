type PossibleStates = "start" | "inlineComment" | "blockComment";
export function manualParser(code: string) {
  let position = 0;
  let currentState: PossibleStates = "start";
  let lastState: PossibleStates = "start";
  const codeLength = code.length;

  function switchStateTo(newState: PossibleStates) {
    lastState = currentState;
    currentState = newState;
  }

  function returnToPreviousState() {
    currentState = lastState;
    lastState = currentState;
  }

  function advancePosition(n: number = 1) {
    position += n;
  }

  while (position < codeLength) {
    // if you see '/' and the next is '/' go to inline comment state
    process.stdout.write(code[position]);
    if (currentState === "start") {
      if (code[position] === "/" && code[position + 1] === "/") {
        console.log("inline comment starts at", position);
        switchStateTo("inlineComment");
        advancePosition(2);
        continue;
      }
      // if you see '/' and the next is `*` go to block comment state
      if (code[position] === "/" && code[position + 1] === "*") {
        console.log("block comment starts at", position);
        switchStateTo("blockComment");
        advancePosition(2);
        continue;
      }
      advancePosition();
      continue;
    }
    if (currentState === "inlineComment") {
      if (code[position] === "\n") {
        returnToPreviousState();
        advancePosition();
        console.log("inline comment ends at", position);
        continue;
      }
      advancePosition();
    }
    if (currentState === "blockComment") {
      if (code[position] === "*" && code[position + 1] === "/") {
        returnToPreviousState();
        advancePosition(2);
        console.log("inline comment ends at", position);
        continue;
      }
      advancePosition();
    }
  }
}

manualParser(`

// This is a comment
import a from "stocazzo"
import b from /* hey hey */ 'stafava'
/*// What? */
`);
