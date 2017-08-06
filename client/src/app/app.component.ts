import { Component } from '@angular/core';

@Component({
    selector: 'app-root',
    template: `
        <div id="page-wrapper">
            <main>
                <nav class="navbar navbar-light bg-faded">
                    <div class="container">
                        <a class="navbar-brand" href="#">Namecard Scanner</a>
                    </div>
                </nav>
                <router-outlet></router-outlet>
            </main>
        </div>`,
})

export class AppComponent {
}