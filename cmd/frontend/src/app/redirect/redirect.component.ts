import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {AuthService, Keys, LocalStorageService} from '../core/services';

@Component({
  selector: 'app-redirect-component',
  templateUrl: 'redirect.component.html',
})
export class RedirectComponent implements OnInit {
  constructor(
    private route: ActivatedRoute,
    private authService: AuthService,
    private localStorage: LocalStorageService,
    private router: Router) {}

  ngOnInit(): void {
    this.route.queryParams.subscribe((params) => {
      this.authService.handleoAuth(params.oauth_token, params.oauth_verifier).subscribe((res) => {
        this.localStorage.set(Keys.BearerToken, res.token);
        this.router.navigateByUrl('/');
      });
    });
  }
}
