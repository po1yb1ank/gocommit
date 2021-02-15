# gocommit v1.1.0
### under construction...

Awesome auto git commit generator and pusher.

## Installation
```bash
go install github.com/po1yb1ank/gocommit
```
## How to use
### Command-line arguments

    -p i -r <link> (-b/-f)      Remote init profile

    -p c (-b/-f)                Commit profile

        -b                      Specify branch  (default if ommitted)
        -f                      Add files (. if ommitted)
Gocommit have different profiles to help with routine git work
- <b>i</b>
  ```bash
  gocommit -p i -r https://github.com/foo/bar
  ```
  will set git repo in the current directory and commit & push present content
- <b>c</b>
  ```bash
  gocommit -p c
  ```
  will work same as in v1.0.0

    ___Notice:___ you can use flags to specify profiles options e.g:
  ```bash
  gocommit -p c -b develop -f notmain.go
  ```
  ```bash
  gocommit -p i -b develop -f notmain.go -r https://github.com/foo/bar
  ```

  However, you can just go with:
  ```bash
    gocommit
  ```
  this will execute committer profile for all files and default branch

### Credits
- [whatthecommit](http://whatthecommit.com/)
