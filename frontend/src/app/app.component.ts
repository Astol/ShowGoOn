import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MenubarModule } from 'primeng/menubar';
import { MenuItem } from 'primeng/api';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, MenubarModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  items: MenuItem[] | undefined;

  ngOnInit() {
    this.items = [
        {
            label: 'Home',
            icon: 'pi pi-home',
            routerLink: ['/']
        },
        {
            label: 'Slides',
            icon: 'pi pi-images',
            routerLink: ['/slides']
        },
        {
            label: 'Settings',
            icon: 'pi pi-cog',
            routerLink: ['/settings']
        },
        {
            label: 'Slideshow',
            icon: 'pi pi-play',
            routerLink: ['/slideshow']
        }
    ]
  }
}
