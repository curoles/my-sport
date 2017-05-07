# See if Firebase works, Sun May 7 2017

Following instructions at
https://medium.com/codingthesmartway-com-blog/angular-2-firebase-introduction-b4f32e844db2

Got errors:
```
WARNING in ./src/app/app.component.ts
25:57-68 "export 'AngularFire' was not found in 'angularfire2'

WARNING in ./src/app/app.component.ts
25:88-99 "export 'AngularFire' was not found in 'angularfire2'

ERROR in /home/igor/prj/github/mysport/mysport/src/app/app.component.ts (2,10): Module '"/home/igor/prj/github/mysport/mysport/node_modules/angularfire2/index"' has no exported member 'AngularFire'.

ERROR in /home/igor/prj/github/mysport/mysport/src/app/app.component.ts (2,23): Module '"/home/igor/prj/github/mysport/mysport/node_modules/angularfire2/index"' has no exported member 'FirebaseListObservable'.

ERROR in /home/igor/prj/github/mysport/mysport/src/app/app.module.ts (18,37): Cannot find name 'firebaseConfig'.
```

Reading:
https://github.com/angular/angularfire2/blob/master/docs/version-4-upgrade.md

Compiles without errors, but I can't see items from the DB.

Trying latest Angular 4
```
npm install @angular/{common,compiler,compiler-cli,core,forms,http,platform-browser,platform-browser-dynamic,platform-server,router,animations}@latest typescript@latest --save
```

for future:
http://stackoverflow.com/questions/38309758/get-user-auth-profile-info-in-firebase-using-angular2

Here are my changes, still can't work with Firebase
```
 git diff tsconfig.json src/app/app.module.ts src/app/app.component.ts src/app/app.component.html | cat
diff --git a/src/app/app.component.html b/src/app/app.component.html
index dbdf5d1..835babe 100644
--- a/src/app/app.component.html
+++ b/src/app/app.component.html
@@ -1,6 +1,23 @@
 <h1>
   {{title}}
 </h1>
+
  <div class="w3-container w3-teal">
   <h1>{{title}}</h1>
-</div> 
+ </div>
+
+<div>test firebase
+<ul>
+<p>user={{user.$value}}</p>
+<p>fitness={{exercises.fitness.pull-up.$value}}</p>
+  <li class="text" *ngFor="let exercise of exercises | async">
+    {{exercise.$value}}
+  </li>
+</ul>
+</div>
+
+test2
+<!--div> {{ (exercises | async)? | json }} </div-->
+<!--div> {{ (user | async)? | json }} </div-->
+<button (click)="login()">Login</button>
+<button (click)="logout()">Logout</button>
diff --git a/src/app/app.component.ts b/src/app/app.component.ts
index 48b015f..1265293 100644
--- a/src/app/app.component.ts
+++ b/src/app/app.component.ts
@@ -1,5 +1,14 @@
 import { Component } from '@angular/core';
 
+import { Observable } from 'rxjs/Observable';
+
+import { AngularFireModule } from 'angularfire2';
+import { AngularFireDatabaseModule, AngularFireDatabase, FirebaseListObservable } from 'angularfire2/database';
+import { AngularFireAuthModule, AngularFireAuth } from 'angularfire2/auth';
+
+// Do not import from 'firebase' as you'd lose the tree shaking benefits
+import * as firebase from 'firebase/app';
+
 @Component({
   selector: 'app-root',
   templateUrl: './app.component.html',
@@ -7,4 +16,21 @@ import { Component } from '@angular/core';
 })
 export class AppComponent {
   title = 'MySport';
+  user: Observable<firebase.User>;
+  exercises: FirebaseListObservable<any[]>;
+
+  constructor(
+    private afAuth: AngularFireAuth,
+    private db: AngularFireDatabase
+  ) {
+    this.user = afAuth.authState;
+    this.exercises = db.list('exercises');
+  }
+
+  login() {
+    this.afAuth.auth.signInWithPopup(new firebase.auth.GoogleAuthProvider());
+  }
+  logout() {
+     this.afAuth.auth.signOut();
+  }
 }
diff --git a/src/app/app.module.ts b/src/app/app.module.ts
index 67ae491..6970616 100644
--- a/src/app/app.module.ts
+++ b/src/app/app.module.ts
@@ -3,6 +3,17 @@ import { NgModule } from '@angular/core';
 import { FormsModule } from '@angular/forms';
 import { HttpModule } from '@angular/http';
 
+import { AngularFireModule } from 'angularfire2';
+import { AngularFireDatabaseModule, AngularFireDatabase, FirebaseListObservable } from 'angularfire2/database';
+import { AngularFireAuthModule, AngularFireAuth } from 'angularfire2/auth';
+
+//import { environment } from '../environments/environment';
+import { firebaseConfig } from '../environments/firebase.config';
+
+// Do not import from 'firebase' as you'd lose the tree shaking benefits
+//import * as firebase from 'firebase/app';
+
+
 import { AppComponent } from './app.component';
 
 @NgModule({
@@ -12,7 +23,11 @@ import { AppComponent } from './app.component';
   imports: [
     BrowserModule,
     FormsModule,
-    HttpModule
+    HttpModule,
+    AngularFireModule.initializeApp(firebaseConfig),
+    //AngularFireModule.initializeApp(environment.firebase, 'my-app'),
+    AngularFireDatabaseModule,
+    AngularFireAuthModule
   ],
   providers: [],
   bootstrap: [AppComponent]
diff --git a/tsconfig.json b/tsconfig.json
index a35a8ee..30cc247 100644
--- a/tsconfig.json
+++ b/tsconfig.json
@@ -15,6 +15,9 @@
     "lib": [
       "es2016",
       "dom"
+    ],
+    "types": [
+      "firebase"
     ]
   }
 }

```

# Setting Up Firebase, Sun May 7 2017

To set up Firebase you first need to go to https://firebase.google.com/
create an account and login to the Firebase console.

- sign-in with curo..s@gmail.com
- go to Firebase Console, https://console.firebase.google.com
- created new App called "MySport"
- on left panel see "Database", https://console.firebase.google.com/project/mysport-f3dab/database/data
- created simple test items `exercises>fitness>pull-up`, https://console.firebase.google.com/project/mysport-f3dab/database/data/exercises/fitness/pull-up

## Firebase Configuration

The Firebase configuration consists for four key-value pairs.
- apiKey
- authDomain
- databaseURL
- storageBucket

The easiest way to get all four pieces of configuration information is to go to the Firebase console,
open the project view and use the Link _Add Firebase to your web app_.
From the pop up which is opened you can copy a code snipped containing
the configuration settings for your specific Firebase project.

```
<script src="https://www.gstatic.com/firebasejs/3.9.0/firebase.js"></script>
<script>
  // Initialize Firebase
  var config = {
    apiKey: "AIzaSyAlcubwURXTZcjIAjcPP3IS1JO8i587YiE",
    authDomain: "mysport-f3dab.firebaseapp.com",
    databaseURL: "https://mysport-f3dab.firebaseio.com",
    projectId: "mysport-f3dab",
    storageBucket: "mysport-f3dab.appspot.com",
    messagingSenderId: "450011041877"
  };
  firebase.initializeApp(config);
</script>
```
Take the settings and paste it into a new file src/environments/firebase.config.ts in the following form:

```
export const firebaseConfig = {
  apiKey: 'AIzaSyAlcubwURXTZcjIAjcPP3IS1JO8i587YiE',
  authDomain: 'mysport-f3dab.firebaseapp.com',
  databaseURL: 'https://mysport-f3dab.firebaseio.com',
  storageBucket: 'mysport-f3dab.appspot.com'
};
```

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
