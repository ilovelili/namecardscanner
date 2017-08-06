import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { rootRouterConfig } from './app.routes';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { HomeComponent } from './component/home.component';

import { FileDropDirective, FileSelectDirective } from 'ng2-file-upload';

@NgModule({
    declarations: [
        AppComponent,
        HomeComponent,
        FileDropDirective,
        FileSelectDirective,
    ],
    imports: [
        BrowserModule,
        FormsModule,
        HttpModule,
        RouterModule.forRoot(rootRouterConfig, { useHash: true })
    ],
    providers: [],
    bootstrap: [AppComponent]
})
export class AppModule { }
