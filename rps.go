package main

import (
    "fmt"
    "net/http"
    "math/rand"
    "golang.org/x/exp/slices"
)


const RPS_HELP = `Please choose rock, paper, or scissors:
Use URL query params to input your choice:
Example:
http://localhost:8080/rock_paper_scissors?choice=rock`

var OPTIONS = []string{"rock", "paper", "scissors"}

func rock_paper_scissors(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    if len(q["choice"]) == 0 {
        fmt.Fprintf(w, RPS_HELP)
        return
    }
    choice := q["choice"][0]
    if !slices.Contains(OPTIONS, choice) {
        fmt.Fprintf(w, "Your input of %s is not an option\n", choice)
        fmt.Fprintf(w, RPS_HELP)
    } else {
        play_choice_and_get_output(w, choice)
    }
}

func tie_message(w http.ResponseWriter, choice string) {
    fmt.Fprintf(w, "You and the computer both chose %s!\nIt's a TIE!", choice)
}

func pick_computer_choice() string {
    return OPTIONS[rand.Intn(len(OPTIONS))]
}

func play_choice_and_get_output(w http.ResponseWriter, player_choice string) {
    computer_choice := pick_computer_choice()

    if computer_choice == player_choice {
      tie_message(w, player_choice)
      return
    }

    fmt.Fprintf(w,"Computer chose %s\n", computer_choice)
    fmt.Fprintf(w,"You chose %s\n", player_choice)

    if computer_choice == "rock" && player_choice == "paper" {
      fmt.Fprintf(w,"You win!")
    }
    

    if computer_choice == "rock" && player_choice == "scissors" {
      fmt.Fprintf(w,"Computer wins!")
    }

    if computer_choice == "paper" && player_choice == "scissors"{
      fmt.Fprintf(w,"You win!")
    }
    

    if computer_choice == "paper" && player_choice == "rock"{
      fmt.Fprintf(w,"Computer wins!")
    }
    

    if computer_choice == "scissors" && player_choice == "rock"{
      fmt.Fprintf(w,"You win!")
    }
    

    if computer_choice == "scissors" && player_choice == "paper" {
      fmt.Fprintf(w,"Computer wins!")
    }
      
}