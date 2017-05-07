

# Setting Up Firebase, Sun May 7 2017

To set up Firebase you first need to go to https://firebase.google.com/
create an account and login to the Firebase console.

- sign-in with curo..s@gmail.com
- go to Firebase Console, https://console.firebase.google.com
- created new App called "MySport"
- on left panel see "Database", https://console.firebase.google.com/project/mysport-f3dab/database/data
- created simple test items `exercises>fitness>pull-up`, https://console.firebase.google.com/project/mysport-f3dab/database/data/exercises/fitness/pull-up

## Installing Firebase and AngularFire 2

Now we need to add the Firebase and the AngularFire 2 library as dependencies to our project:
 
`$ npm install angularfire2 firebase --save`
 
By using the `--save` flag both dependencies are automatically added to the project's *package.json* file.

```
diff --git a/package.json b/package.json
index d24e9a0..7ba5bef 100644
--- a/package.json
+++ b/package.json
@@ -20,7 +20,9 @@
     "@angular/platform-browser": "^4.0.0",
     "@angular/platform-browser-dynamic": "^4.0.0",
     "@angular/router": "^4.0.0",
+    "angularfire2": "^4.0.0-rc.0",
     "core-js": "^2.4.1",
+    "firebase": "^3.9.0",
     "rxjs": "^5.1.0",
     "zone.js": "^0.8.4"
   },
```

# Sat May  6 23:49:14 PDT 2017

Reading about Firebase DB, looks like DB can **NOT** be hosted locally,
it always hosted on Google Cloud, https://firebase.google.com.

Tutorials:
- https://medium.com/codingthesmartway-com-blog/angular-2-firebase-introduction-b4f32e844db2
- https://progblog.io/Angular-2-Firebase-Tutorial-Part-1-Create-a-Firebase-3-CRUD-app-with-Angular-CLI/

If I use Firebase, then for web hosting using Google Cloud starts making sense :
- https://cloud.google.com/solutions/web-serving-overview
- https://cloud.google.com/storage/docs/hosting-static-website
- https://cloud.google.com/solutions/websites/

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
