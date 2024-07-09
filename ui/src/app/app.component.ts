import { HttpClient } from '@angular/common/http';
import { Component, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, FormsModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent {
  constructor(private httpClient: HttpClient) {}
  title = 'ui';

  text = signal('');

  getNewText() {
    this.httpClient
      .get<string>(`api/${this.text()}`, {
        responseType: 'json',
      })
      .subscribe((data) => {
        this.title = data;
      });
  }
}
