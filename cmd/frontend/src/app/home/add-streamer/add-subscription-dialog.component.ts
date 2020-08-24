import {Component, Inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';

@Component({
  selector: 'app-add-streamer-dialog',
  templateUrl: 'add-subscription-dialog.component.html',
})
export class AddSubscriptionDialog {
  name: string;

  constructor(public dialogRef: MatDialogRef<AddSubscriptionDialog>) {}

  submit() {
    // this.streamersService.addStreamer(this.name).subscribe(() => {
    //   this.dialogRef.close();
    // })
  }
}
