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
    if (this.localStorageService.get(Keys.BearerToken) === null) {
      this.login().subscribe(x => {
        window.location.assign(x.url);
      });
    }
  }

  login(): Observable<any> {
    return this.apiService.get('/auth/login');
  }

  handleoAuth(oauthToken: string, oauthVerifier: string): Observable<any> {
    return this.apiService.post('/auth/login', {
      oauth_token: oauthToken,
      oauth_verifier: oauthVerifier,
    });
  }
}
