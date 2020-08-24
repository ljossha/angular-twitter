import {Injectable} from '@angular/core';
import {Keys, LocalStorageService} from './localstorage.service';
import {ApiService} from './api.service';
import {Observable} from 'rxjs';
import {Router} from '@angular/router';

@Injectable()
export class AuthService {
  constructor(
    private localStorageService: LocalStorageService,
    private apiService: ApiService,
    private router: Router,
  ) {}

  populate(): void {
    let token = this.localStorageService.get(Keys.BearerToken);
    // If already expired -> delete from storage and populate login.
    if (token !== null) {
      this.me().subscribe(() => {}, () => {
        this.localStorageService.delete(Keys.BearerToken);
        token = null;
      });
    }

    if (token === null) {
      this.login().subscribe(x => {
        window.location.assign(x.url);
      });
    }
  }

  login(): Observable<any> {
    return this.apiService.get('/auth/login');
  }

  me(): Observable<any> {
    return this.apiService.get('/auth/me');
  }

  handleoAuth(oauthToken: string, oauthVerifier: string): Observable<any> {
    return this.apiService.post('/auth/login', {
      oauth_token: oauthToken,
      oauth_verifier: oauthVerifier,
    });
  }
}
