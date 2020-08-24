import {Component, Inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import {SubscriptionService} from '../../core/services/subscription.service';

@Component({
  selector: 'app-add-streamer-dialog',
  templateUrl: 'add-subscription-dialog.component.html',
})
export class AddSubscriptionDialogComponent {
  name: string;

  constructor(
    public dialogRef: MatDialogRef<AddSubscriptionDialogComponent>,
    private subscriptionService: SubscriptionService) {}

  submit() {
    this.subscriptionService.addSubscription(this.name).subscribe(() => {
      this.dialogRef.close();
    });
  }
}
