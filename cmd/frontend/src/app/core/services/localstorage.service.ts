import {Injectable} from '@angular/core';

export enum Keys {
  BearerToken = 'token',
}

@Injectable()
export class LocalStorageService {
  constructor() {}

  get(key: Keys): string {
    return localStorage.getItem(key);
  }

  set(key: Keys, value: string) {
    localStorage.setItem(key, value);
  }

  delete(key: Keys) {
    localStorage.removeItem(key);
  }
}
