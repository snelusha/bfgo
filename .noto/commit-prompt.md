# Commit Message Guidelines

## Format
Use the conventional commits format without scopes: `type: description`.

## Style Rules
- **Tense**: Imperative/Present ("add", "implement", "use")
- **Capitalization**: Lowercase the first letter of the description
- **Length**: Concise and direct (typically 30-60 characters)
- **Tone**: Technical and matter-of-fact

## Commit Types
- `feat`: Implementation of new core logic or components (e.g., lexer, parser, interpreter)
- `fix`: Resolution of bugs or logic errors
- `refactor`: Improvements to existing code structure, error handling, or internal logic
- `test`: Addition or modification of test cases
- `chore`: General maintenance, updates to examples, or minor file adjustments

## Scope Usage
No scopes used. Changes are identified by the type and the description content rather than parenthetical scopes.

## Description Patterns
Descriptions usually start with an action verb like "add", "implement", "use", "set", or "update". They are specific about the technical change (e.g., naming a specific function like `countRepeats` or an error handling method like `errors.Is`).

## Examples from History
- feat: implement interpreter
- refactor: use errors.Is for EOF check
- test: add tests for input EOF and byte wrapping
- fix: set correct operand for jump instructions
- chore: add hello world example