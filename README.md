# gocommit
Awesome auto git commit generator and pusher.
This little package allows you to make this:

```bash
git add .
git commit -m "Very long commit message I came up with. Who will ever read this???"
git push origin master
```
into this:
```bash
gocommit -f . -b master
```
or even this (will push to default branch all files in the repo):
```bash
gocommit
```

## Installation
```bash
go install github.com/po1yb1ank/gocommit
```

### Credits
- [whatthecommit](http://whatthecommit.com/)
