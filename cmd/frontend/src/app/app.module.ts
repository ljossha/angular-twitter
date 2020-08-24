import { BrowserModule } from '@angular/platform-browser';
import {CUSTOM_ELEMENTS_SCHEMA, NgModule} from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatSidenavModule} from '@angular/material/sidenav';
import { RedirectComponent } from './redirect/redirect.component';
import {MatDialogModule} from '@angular/material/dialog';
import {HTTP_INTERCEPTORS, HttpClient, HttpClientModule} from '@angular/common/http';
import {HomeComponent, SafePipe} from './home/home.component';
import {MatListModule} from '@angular/material/list';
import {MatMenuModule} from '@angular/material/menu';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatToolbarModule} from '@angular/material/toolbar';
import {FlexLayoutModule} from '@angular/flex-layout';
import {HttpTokenInterceptor} from './core/services/http.token.interceptor';
import {AddSubscriptionDialogComponent} from './home/add-subscription/add-subscription-dialog.component';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {FormsModule} from '@angular/forms';
import {MatSnackBar} from '@angular/material/snack-bar';
import {LocalStorageService, UserService, ApiService, AuthService} from './core/services';
import {NgxSpinnerModule} from 'ngx-spinner';
import {SubscriptionService} from './core/services/subscription.service';
import {TweetService} from './core/services/tweet.service';
import {MatCardModule} from '@angular/material/card';

@NgModule({
  declarations: [
    AppComponent,
    RedirectComponent,
    HomeComponent,
    SafePipe,
    AddSubscriptionDialogComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatSidenavModule,
    MatDialogModule,
    HttpClientModule,
    MatListModule,
    MatMenuModule,
    MatIconModule,
    MatButtonModule,
    MatToolbarModule,
    FlexLayoutModule,
    MatFormFieldModule,
    MatInputModule,
    FormsModule,
    NgxSpinnerModule,
    MatCardModule,
  ],
  providers: [
    UserService,
    ApiService,
    HttpClient,
    LocalStorageService,
    AuthService,
    SubscriptionService,
    TweetService,
    { provide: HTTP_INTERCEPTORS, useClass: HttpTokenInterceptor, multi: true },
    MatSnackBar,
  ],
  bootstrap: [AppComponent],
})
export class AppModule { }
