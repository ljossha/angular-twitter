import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';

import {ApiService} from './api.service';
import {Keys, LocalStorageService} from './localstorage.service';

@Injectable()
export class UserService {
  constructor(
    private apiService: ApiService,
    private localStorageService: LocalStorageService,
  ) {}

  isLogged(): boolean {
    const token = this.localStorageService.get(Keys.BearerToken);
    return token !== null;
  }
}
