import { Component } from '@angular/core';
import { InputTextModule } from 'primeng/inputtext';
import { FormGroup, FormControl, ReactiveFormsModule} from '@angular/forms';
import { FloatLabelModule } from 'primeng/floatlabel';
import { TextareaModule } from 'primeng/textarea';
import { ButtonModule } from 'primeng/button';

@Component({
  selector: 'app-slides',
  imports: [
    InputTextModule,
    ReactiveFormsModule,
    FloatLabelModule,
    TextareaModule,
    ButtonModule
  ],
  templateUrl: './slides.component.html',
  styleUrl: './slides.component.scss'
})
export class SlidesComponent {
  slideForm = new FormGroup({
    title: new FormControl(''),
    description: new FormControl('')
  });

  onSubmit() {
    
  }
}
