import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';

import {ApiService} from './api.service';
import {TweetModel} from '../models';

@Injectable()
export class TweetService {
  constructor(
    private apiService: ApiService
  ) {}

  getTimeline(id: number): Observable<TweetModel[]> {
    return this.apiService.get(`/tweet/user/${id}`);
  }
}
