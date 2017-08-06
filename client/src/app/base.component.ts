import { ActivatedRoute } from '@angular/router';

export abstract class BaseComponent {
    protected items: any[];
    private params: Object;

    constructor(protected activatedRoute: ActivatedRoute) {
        this.activatedRoute.params.subscribe((params) => {
            this.params = params;
        });
    }

    protected LogComplete = (msg: string) => {
        console.log(msg);
    }

    protected LogError = (err: string) => {
        console.error(err);
    }
}
