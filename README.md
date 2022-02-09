# Perkbox

A simple todo list server and cli that calls todo placeholder api based off a file of ids found in `./input/data.csv`.

### Setup
Cli:
```bash
make cli
```
Server:
```bash
make server
```

### Improvements
A very simple solution but can be improved quickly by using channels, waitgroups and goroutines to run the api calls in parallel speeding up both the cli and api.

### Notes
Wasn't too sure how you guys wanted the cli output - whether you wanted it pretty printed out in json or 
just log each todo seperately. I went for a pretty print to console as it felt right with the cli.
