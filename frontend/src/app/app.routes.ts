import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { SlidesComponent } from './slides/slides.component';
import { SettingsComponent } from './settings/settings.component';
import { SlideshowComponent } from './slideshow/slideshow.component';

export const routes: Routes = [
    
  { path: '', component: HomeComponent },
  { path: 'slides', component: SlidesComponent },
  { path: 'settings', component: SettingsComponent },
  { path: 'slideshow', component: SlideshowComponent },
  { path: '**', redirectTo: '' }
];
