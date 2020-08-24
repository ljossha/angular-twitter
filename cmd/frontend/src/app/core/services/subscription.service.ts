import {Injectable} from '@angular/core';
import {ApiService} from './api.service';
import {Keys, LocalStorageService} from './localstorage.service';
import {Observable} from 'rxjs';
import {SubscriptionModel} from '../models';

@Injectable()
export class SubscriptionService {
  constructor(
    private apiService: ApiService,
    private localStorageService: LocalStorageService,
  ) {}

  addSubscription(name: string): Observable<any> {
    return this.apiService.post('/subscription', {
      name
    });
  }

  removeSubscription(id: number): Observable<any> {
    return this.apiService.delete(`/subscription/${id}`);
  }

  list(): Observable<SubscriptionModel[]> {
    return this.apiService.get('/subscription');
  }
}
