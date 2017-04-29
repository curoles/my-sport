# Fri Apr 28 17:14:01 PDT 2017

To get sources and run:
- `mkdir -p <path>/mysport/mysport`
- `cd <path>/mysport/mysport`
- `git init`
- `git remote add origin https://github.com/curoles/my-sport.git`
- if needed, run `get-tools.bash` and source generated `setup-env.bash`
- run `ng -v` to make sure it does not show "error"
- to populate|create directory `node_modules` run `npm install`
- now you can do `ng serve --open`

# Thu Apr 27 22:43:55 PDT 2017

To insert current date at the top of this journa;:
```
sed -i "1i`date`" doc/journal.md
```

To show git-ignored files:
```
git status --ignored
git check-ignore *
git check-ignore -v *
git check-ignore -v **/*
```

So far only one folder `node_modules/` is ignored.

Add it to my github.
[Add exisitng repo to github](https://help.github.com/articles/adding-an-existing-project-to-github-using-the-command-line/)

```
git remote add origin https://github.com/curoles/my-sport.git
git push -u origin master
```

# Apr 26

- install Node.js (see get-tools.bash)
    ```
    mkdir ${NODE_JS_DIR}
    wget https://nodejs.org/dist/v${NODE_JS_VER}/${NODE}.tar.xz -P ${NODE_JS_DIR}
    tar xvfJ ${NODE_JS_DIR}/${NODE}.tar.xz -C ${NODE_JS_DIR}
    ```
- install Angular
    ```
    npm install -g @angular/cli
    ```
- create App project
  + get-tools.bash generates setup-env.bash
  + source setup-env.bash, now node,npm and ng are available
  + Angular framework CLI command creates files/folders with certain structure ready for git
    (Angular seems git|github friendly, ng has command to create gh-pages)
    ```
    ng new my-sport
    ```
- to see this journal
  + install [Markdown Viewer](https://addons.mozilla.org/en-us/firefox/addon/markdown-viewer/)
  + launch Firefox
    ```
    $ firefox doc/journal.md &
    ```

# Apr 25

Old idea to have an online log for sport exercises.
- web design as Single Page App
- use Node and Angular
- some Data Base
- sounds IO bound instead of CPU/Mem, Node should be OK
- simple as single START/STOP button
  + exercise type can be specified later on Calendar
  + or before pressing START
- (public|private) exercise type; scope? global|team|personal 
