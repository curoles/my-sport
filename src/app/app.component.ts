import { Component } from '@angular/core';
import { Router } from "@angular/router";

import { Observable } from 'rxjs/Observable';

import { AngularFireModule } from 'angularfire2';
import { AngularFireDatabaseModule, AngularFireDatabase } from 'angularfire2/database';
import { FirebaseObjectObservable, FirebaseListObservable } from 'angularfire2/database';
import { AngularFireAuthModule, AngularFireAuth } from 'angularfire2/auth';

// Do not import from 'firebase' as you'd lose the tree shaking benefits
import * as firebase from 'firebase/app';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'MySport';
  user: Observable<firebase.User>;
  exercises: FirebaseListObservable<any[]>;
  items: FirebaseListObservable<any>;
  testitem: FirebaseObjectObservable<any>;

  constructor(
    private router: Router,
    private afAuth: AngularFireAuth,
    private db: AngularFireDatabase
  ) {
    console.log("constructing App");
    this.user = afAuth.authState;
    //this.exercises = db.list('/exercises');
    //this.items = db.list('/items');
    //this.testitem = db.object('/testitem');

    afAuth.auth.onAuthStateChanged( (authChange: any) => {
      this.onUserLogInOut(authChange);
    });
  }

  login() {
    this.afAuth.auth.signInWithPopup(new firebase.auth.GoogleAuthProvider());
    console.log("user login ",this.afAuth.auth.currentUser);
  }

  logout() {
     this.afAuth.auth.signOut();
  }

  // https://firebase.google.com/docs/auth/web/manage-users
  // https://firebase.google.com/docs/reference/js/firebase.auth.Auth
  onUserLogInOut(authChange: any) {
    console.log("auth state changed ", JSON.stringify(authChange));

    if (this.afAuth.auth.currentUser) {
      var currentUser = this.afAuth.auth.currentUser;
      console.log(`user ${currentUser.displayName} logged in`);
      this.router.navigate(['user-page']);
    }
    else {
      console.log("user logged out");
    }

  }
}
