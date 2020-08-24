import {Component, OnInit, Pipe, PipeTransform} from '@angular/core';
import {DomSanitizer} from '@angular/platform-browser';
import {MatDialog} from '@angular/material/dialog';
import {AuthService, NotificationService} from '../core/services';
import {AddSubscriptionDialog} from './add-streamer/add-subscription-dialog.component';
import {NgxSpinnerService} from 'ngx-spinner';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  constructor(
    public dialog: MatDialog,
    private authService: AuthService,
    private notificationService: NotificationService,
    private spiner: NgxSpinnerService) { }
  ngOnInit(): void {
    this.spiner.show();
    this.authService.populate();
    this.spiner.hide();
  }

  openDialog() {
    const dialogRef = this.dialog.open(AddSubscriptionDialog, {
      width: '250px',
    });

    dialogRef.afterClosed().subscribe(result => {

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
