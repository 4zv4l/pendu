# pendu
pendu written in go.
### jeu.go
contains all the code for the game
### Words file
the game needs the `words.txt` file to work, you can put all the words you want inside.

> need to write the entire word to succeed the level

## BUILD
### for your plateform
```
go build jeu.go
```
### for other plateform
```
GOOS=(windows/linux/darwin/android) GOARCH=(amd64, arm64,...) go build jeu.go
```

## ERROR
- [x] work even when there is a blank line at the end of the `words.txt`.
- [ ] need to write the word even when all letters are found.
- [ ] write a letter again still remove a life
- [ ] write many letters at once isn't supposed to arrive
