import { HttpClient } from "@angular/common/http";
import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet],
  providers: [HttpClient],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {

  /**
   *
   */
  constructor(private http: HttpClient) { }
  value: string = "";

  ngOnInit(): void {
    this.http.get("./api", {
      headers: {
        "tenant": "5"
      }
    }).subscribe(
      {
        next: (response: any) => {
          this.value = response.data;
        },
        error: (error: any) => {
          console.log(error);
          this.value = "Some error happend";

        }
      }
    );

    this.http.get("", {}).pipe;
  }
  title = 'App';
}
