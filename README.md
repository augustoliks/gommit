# gommit

*golang* + *git commit* = _**gommit**_

CLI to standardized git commits messages.

:warning: **In development**

## How Install

[![asciicast](https://asciinema.org/a/413917.svg)](https://asciinema.org/a/413917)

## Pattern Suggested

- ACTION(TOPIC): COMMIT_MSG #TICKET_NUM [!TAG ...]

### ACTION:

- `feat`: is for new features, big improvement;
- `fix`: is for bug fixes;
- `refactor`: is for code refactoring, without any visual user change;
- `chore`: is for any else but the app;
- `test`: is for any kind of tests (functionnal, unitary...);
- `style`: is for cosmetic changes (spaces vs tab...);
- `docs`: is for documentation.

### TOPIC 

Is optional and to specify the part on the app you're working on is WHO is concerned by the change.

- `dev`:  is for developpers (API changes, refactors...)
- `usr`:  is for final users (UI changes)
- `pkg`:  is for packagers   (packaging changes)
- `test`: is for testers     (test only related changes)

### COMMIT_MSG 

is ... well ... the commit message itself.

### TICKET_NUM 

is the ticket number associated. If you don't want to specify a ticket number, use:

- `#noticket`; or
- `#nt`

### TAGs 

Are additionnal adjective as 'refactor' 'minor' 'cosmetic'. They are preceded with a '!' or a '@' (prefer the former, as the latter is wrongly interpreted in github.) Commonly used tags are:

- `refactor`: is obviously for refactoring code only
- `minor`: is for a very meaningless change (a typo, adding a comment)
- `cosmetic`: is for cosmetic driven change (re-indentation, 80-col...)
- `wip`: is for partial functionality but complete subfunctionality.

