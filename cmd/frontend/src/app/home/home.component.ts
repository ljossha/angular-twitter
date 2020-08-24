import {Component, OnInit, Pipe, PipeTransform} from '@angular/core';
import {DomSanitizer} from '@angular/platform-browser';
import {MatDialog} from '@angular/material/dialog';
import {AuthService, NotificationService} from '../core/services';
import {AddSubscriptionDialogComponent} from './add-subscription/add-subscription-dialog.component';
import {NgxSpinnerService} from 'ngx-spinner';
import {SubscriptionService} from '../core/services/subscription.service';
import {SubscriptionModel, TweetModel} from '../core/models';
import {TweetService} from '../core/services/tweet.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  subscriptions: SubscriptionModel[];
  tweets: TweetModel[];

  constructor(
    public dialog: MatDialog,
    private authService: AuthService,
    private notificationService: NotificationService,
    private spiner: NgxSpinnerService,
    private tweetService: TweetService,
    private subscriptionService: SubscriptionService) { }
  ngOnInit(): void {
    this.authService.populate();
    this.populateListOfSubscriptions();
  }

  populateListOfSubscriptions() {
    this.spiner.show();
    this.subscriptionService.list().subscribe(res => {
      this.subscriptions = res;
      this.spiner.hide();
    });
  }

  openDialog() {
    const dialogRef = this.dialog.open(AddSubscriptionDialogComponent, {
      width: '250px',
    });

    dialogRef.afterClosed().subscribe(result => {
      this.populateListOfSubscriptions();
    });
  }

  changeSubscription(id: number) {
    this.spiner.show();
    this.tweetService.getTimeline(id).subscribe(res => {
      this.tweets = res;
      this.spiner.hide();
    });
  }
}

@Pipe({ name: 'safe' })
export class SafePipe implements PipeTransform {
  constructor(private sanitizer: DomSanitizer) {}
  transform(url) {
    return this.sanitizer.bypassSecurityTrustResourceUrl(url);
  }
}
