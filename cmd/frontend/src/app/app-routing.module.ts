import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {RedirectComponent} from './redirect/redirect.component';
import {HomeComponent} from './home/home.component';

const routes: Routes = [
    {
      path: 'redirect',
      component: RedirectComponent,
    },
    {
      path: '',
      component: HomeComponent,
    }
  ];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
