import { Get, Controller, Render, Query } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import * as crypto from 'crypto';
import * as qs from 'querystring';
import { AuthService } from './auth.service';

@Controller('login')
export class AuthController {
  constructor(private authService: AuthService, private configService: ConfigService){

  }

  @Get()
  @Render('index')
  root() {
    const state:string = crypto.randomBytes(43).toString('hex');
    const nonce:string = crypto.randomBytes(43).toString('hex');
    const query:string = qs.stringify({
      response_type: 'code',
      client_id: this.configService.get<string>('LINE_CLIENT_ID'),
      redirect_uri: `${this.configService.get<string>('SERVER_URI')}/login/auth`,
      state,
      scope: 'profile openid',
      nonce
    })
    return { lineAuthLoginURI: `${this.configService.get<string>('LINE_ACCESS_URI')}/oauth2/v2.1/authorize?${query}` };
  }

  @Get('/auth')
  @Render('auth')
  async auth(@Query('code') code) {
    try {
      const token = await this.authService.postToken(code).toPromise()
      return { token: JSON.stringify(token)}
    } catch (err) {
      console.log(err)
    }
  }
}