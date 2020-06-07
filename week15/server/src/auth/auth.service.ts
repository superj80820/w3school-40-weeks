import { Injectable, HttpService } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { map } from 'rxjs/operators';
import * as qs from 'querystring';

@Injectable()
export class AuthService {
  constructor(private http: HttpService, private configService: ConfigService) {

  }

  postToken(code){
    return this.http.post(
      `${this.configService.get<string>('LINE_API_URI')}/oauth2/v2.1/token`,
      qs.stringify({
        grant_type: 'authorization_code',
        code,
        redirect_uri: `${this.configService.get<string>('SERVER_URI')}/login/auth`,
        client_id: this.configService.get<string>('LINE_CLIENT_ID'),
        client_secret: this.configService.get<string>('LINE_CLICLIENT_SECRET')
      }),
      {
        headers: {'Content-Type': 'application/x-www-form-urlencoded'}
      })
      .pipe(
        map(response => response.data)
      );        
  }
}
