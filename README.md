# gocommit v1.1.0 
### under construction...

Awesome auto git commit generator and pusher.

## Installation
```bash
go install github.com/po1yb1ank/gocommit
```
## How to use

Gocommit have different profiles to help with routine git work
- ***reminit***
  ```bash
  gocommit -p reminit -r https://github.com/foo/bar
  ```
  will set git repo in current directory and commit & push present content
- ***committer***
  ```bash
  gocommit -p committer
  ```
  will work same as in v1.0.0
  
  >___Notice:___ you can use flags to specify profiles options e.g:
  ```bash
  gocommit -p committer -b develop -f notmain.go
  ```
  ```bash
  gocommit -p reminit -b develop -f notmain.go -r https://github.com/foo/bar
  ```
  
  
### Credits
- [whatthecommit](http://whatthecommit.com/)
