import { Component, OnInit } from '@angular/core';
import { Text } from '../model/text';
import { BaseComponent } from '../base.component';
import { ActivatedRoute } from '@angular/router';
import { FileUploader } from 'ng2-file-upload';

@Component({
    // selector: 'home',
    templateUrl: './home.component.html',
    styleUrls: ['./home.component.css'],
})

export class HomeComponent extends BaseComponent implements OnInit {    
    private text: Text = {
        Text: '',
        Success: false,
    };

    private message: string;
    private fileUploader: FileUploader;

    private hasBaseDropZoneOver = false;
    private hasAnotherDropZoneOver = false;
    private endpoint = 'http://46.101.141.88:3000/text';
    constructor(protected activatedRoute: ActivatedRoute) {
        super(activatedRoute);
    }

    fileOverBase = (e) => {
        this.hasBaseDropZoneOver = e;
    }

    ngOnInit() {
        this.fileUploader = new FileUploader({
            url: this.endpoint,
            method: 'POST',
            autoUpload: true,
        });

        this.fileUploader.onCompleteItem = (item: any, response: any, status: any, headers: any) => {            
            this.text = JSON.parse(response);
        };

        this.fileUploader.onErrorItem = (item: any, response: any, status: any, headers: any) => {
            this.message = response;
            setTimeout(function () {
                this.message = '';
            }.bind(this), 10000);
        };
    };
}
