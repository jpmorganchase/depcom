# state machine

## start state

- if you see '/' and the next is '/' go to inline comment state
- if you see '/' and the next is `*` go to block comment state
- if you see `'`, `"` or '`' go to string state
- if you see "import " (import plush whitespace), go to import state (advance counter)

## inline comment state

- if you see \n go to previous state

## block comment state

- if you see `*` and the next is '/' go to previous state

## string state

- if you see the matching starting character and the previous character is not backslash, return to previous state with string value

## import state

- if you see '/' and the next is '/' go to inline comment state
- if you see '/' and the next is `*` go to block comment state
- if you see `'`, `"` or '`' go to string state
- if you see 'from' go to from state

## from state

- if you see '/' and the next is '/' go to inline comment state
- if you see '/' and the next is `*` go to block comment state
- if you see `'`, `"` or '`' go to string state
- when exiting from string state, save the value and go back to start state
