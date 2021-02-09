# gocommit
Auto git commit generator.
This little package allows you to make this:
```bash
git add .
git commit -m "very long commit message that I need to write. Who will ever read this???"
git push origin master
```
into this:
```bash
gocommit -f . -b master
```
## Installation
```bash
go install github.com/po1yb1ank/gocommit
```
### Credits
- [whatthecommit](http://whatthecommit.com/)
